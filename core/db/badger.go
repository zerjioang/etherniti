// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package db

import (
	"os"
	"time"

	"github.com/zerjioang/etherniti/core/bus"
	gobus "github.com/zerjioang/go-bus"

	"github.com/zerjioang/etherniti/shared/constants"

	"github.com/zerjioang/etherniti/core/util/fs"

	"github.com/zerjioang/etherniti/core/data"

	"github.com/dgraph-io/badger"
	"github.com/zerjioang/etherniti/core/logger"
)

type BadgerStorage struct {
	// our custom database configuration options
	options *Options
	// instance is the underlying handle to the db.
	instance            *badger.DB
	vlogTicker          *time.Ticker // runs every 1m, check size of vlog and run GC conditionally.
	mandatoryVlogTicker *time.Ticker // runs every 10m, we always run vlog GC.
}

var (
	defaultConfig = badger.DefaultOptions
	uid           = os.Getuid()
	gid           = os.Getgid()
)

func createData(path string) error {
	logger.Debug("creating db path")
	defaultConfig.Dir = path
	defaultConfig.ValueDir = path
	// create dir if does not exists
	if !fs.Exists(path) {
		logger.Debug("creating dir: ", path)
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			logger.Error("failed to create database dir:", err)
			return err
		}
	}
	// overwrite directory permissions
	logger.Debug("setting dir permissions and ownership: ", path)
	permErr := chownR(path, uid, gid, os.ModePerm)
	if permErr != nil {
		logger.Error("failed to set permissions: ", permErr)
	}
	return permErr
}

func NewCollection(name string) (*BadgerStorage, error) {
	logger.Debug("creating new db collection")
	err := createData(constants.DatabaseRootPath + name)
	if err != nil {
		return nil, err
	}
	var openErr error
	collection := new(BadgerStorage)
	collection.instance, openErr = badger.Open(defaultConfig)
	if err != nil {
		return nil, openErr
	}
	// register for listening poweroff events
	bus.SharedBus().Subscribe(bus.PowerOffEvent, func(message gobus.EventMessage) {
		logger.Debug("executing database poweroff routine in database: ", name)
		err := collection.Close()
		if err != nil {
			logger.Error("failed to close database due to: ", err)
		}
	})
	return collection, nil
}

func (db *BadgerStorage) Set(key []byte, val []byte) error {
	logger.Debug("inserting key-value in db")
	return db.instance.Update(func(txn *badger.Txn) error {
		return txn.Set(key, val)
	})
}

func (db *BadgerStorage) PutUniqueKeyValue(key []byte, value []byte) error {
	logger.Debug("inserting unique key-value in db")
	err := db.instance.Update(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err == nil && item != nil {
			return data.DuplicateKeyErr
		} else {
			return txn.Set(key, value)
		}
	})
	return err
}

// Get is used to retrieve a value from the k/v store by key
func (db *BadgerStorage) Get(key []byte) ([]byte, error) {
	logger.Debug("reading key from db")
	var value []byte
	err := db.instance.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			switch err {
			case badger.ErrKeyNotFound:
				return badger.ErrKeyNotFound
			default:
				return err
			}
		}
		value, err = item.ValueCopy(value)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return value, nil
}

func (db *BadgerStorage) Add(key, data []byte) error {
	logger.Debug("adding new key-value to db")
	return db.Set(key, data)
}

func (db *BadgerStorage) List(prefixStr string) ([][]byte, error) {
	logger.Debug("listing values from db")
	var results [][]byte
	execErr := db.instance.View(func(txn *badger.Txn) error {
		it := txn.NewIterator(badger.DefaultIteratorOptions)
		prefix := []byte(prefixStr)
		for it.Seek(prefix); it.ValidForPrefix(prefix); it.Next() {
			item := it.Item()
			k := item.Key()
			var readedVal []byte
			readedVal, err := item.ValueCopy(readedVal)
			if err != nil {
				it.Close()
				return err
			}
			logger.Debug("key: ", k, " value: ", string(readedVal))
		}
		it.Close()
		return nil
	})
	return results, execErr
}

func (db *BadgerStorage) Delete(key []byte) error {
	logger.Debug("deleting key from db: ", string(key))
	execErr := db.instance.Update(func(txn *badger.Txn) error {
		err := txn.Delete(key)
		if err != nil {
			logger.Error("failed to delete key from db: ", string(key), "error: ", err)
		}
		return err
	})
	return execErr
}

// DeleteRange deletes logs within a given range inclusively.
func (db *BadgerStorage) DeleteRange(min, max uint64) error {
	logger.Debug("deleting by range from db")
	// we manage the transaction manually in order to avoid ErrTxnTooBig errors
	txn := db.instance.NewTransaction(true)
	it := txn.NewIterator(badger.IteratorOptions{
		PrefetchValues: false,
		Reverse:        false,
	})

	for it.Seek(uint64ToBytes(min)); it.Valid(); it.Next() {
		key := make([]byte, 8)
		it.Item().KeyCopy(key)
		// Handle out-of-range log index
		if bytesToUint64(key) > max {
			break
		}
		// Delete in-range log index
		if err := txn.Delete(key); err != nil {
			if err == badger.ErrTxnTooBig {
				it.Close()
				err = txn.Commit(nil)
				if err != nil {
					return err
				}
				return db.DeleteRange(bytesToUint64(key), max)
			}
			return err
		}
	}
	it.Close()
	err := txn.Commit(nil)
	if err != nil {
		return err
	}
	return nil
}

// SetUint64 is like Set, but handles uint64 values
func (db *BadgerStorage) SetUint64(key []byte, val uint64) error {
	return db.Set(key, uint64ToBytes(val))
}

// GetUint64 is like Get, but handles uint64 values
func (db *BadgerStorage) GetUint64(key []byte) (uint64, error) {
	val, err := db.Get(key)
	if err != nil {
		return 0, err
	}
	return bytesToUint64(val), nil
}

func (db *BadgerStorage) runVlogGC(instance *badger.DB, threshold int64) {
	// Get initial size on start.
	_, lastVlogSize := instance.Size()

	runGC := func() {
		var err error
		for err == nil {
			// If a GC is successful, immediately run it again.
			err = instance.RunValueLogGC(0.7)
		}
		_, lastVlogSize = instance.Size()
	}

	for {
		select {
		case <-db.vlogTicker.C:
			_, currentVlogSize := instance.Size()
			if currentVlogSize < lastVlogSize+threshold {
				continue
			}
			runGC()
		case <-db.mandatoryVlogTicker.C:
			runGC()
		}
	}
}

// Close is used to gracefully close the DB connection.
func (db *BadgerStorage) Close() error {
	logger.Debug("closing database")
	if db.vlogTicker != nil {
		db.vlogTicker.Stop()
	}
	if db.mandatoryVlogTicker != nil {
		db.mandatoryVlogTicker.Stop()
	}
	return db.instance.Close()
}
