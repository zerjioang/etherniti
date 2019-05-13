// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package db

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"

	"github.com/zerjioang/etherniti/core/util/fs"

	"github.com/zerjioang/etherniti/core/config"
	"github.com/zerjioang/etherniti/core/data"

	"github.com/dgraph-io/badger"
	"github.com/zerjioang/etherniti/core/logger"
	"golang.org/x/crypto/bcrypt"
)

type Db struct {
	instance *badger.DB
}

var (
	defaultConfig = badger.DefaultOptions
	instance      *Db
	once          sync.Once
	uid           = os.Getuid()
	gid           = os.Getgid()
)

func init() {
	logger.Debug("loading db module data")
	err := createData(config.DatabaseRootPath + "db")
	if err != nil {
		logger.Error("failed to create shared database dir:", err)
	}
}

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

func NewCollection(name string) (*Db, error) {
	logger.Debug("creating new db collection")
	collection := new(Db)
	err := createData(config.DatabaseRootPath + name)
	if err != nil {
		return nil, err
	}
	var openErr error
	collection.instance, openErr = badger.Open(defaultConfig)
	if err != nil {
		return nil, openErr
	}
	return collection, nil
}

func (db *Db) Init() error {
	logger.Debug("initializing db file")
	// Open the Badger database located in the /data/badger directory.
	// It will be created if it doesn't exist.
	instance, err := badger.Open(defaultConfig)
	if err != nil {
		return err
	}
	db.instance = instance
	return nil
}

func (db *Db) Query() error {
	logger.Debug("querying db")
	err := db.instance.View(func(txn *badger.Txn) error {
		// Your code here…
		return nil
	})
	return err
}

func (db *Db) PutKeyValue(key []byte, value []byte) error {
	logger.Debug("inserting key-value in db")
	err := db.instance.Update(func(txn *badger.Txn) error {
		err := txn.Set(key, value)
		return err
	})
	return err
}

func (db *Db) Close() error {
	logger.Debug("closing database")
	return db.instance.Close()
}

func (db *Db) PutUniqueKeyValue(key []byte, value []byte) error {
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

func (db *Db) GetKeyValue(key []byte) ([]byte, error) {
	logger.Debug("reading key from db")
	var readedVal []byte
	err := db.instance.Update(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err == nil && item != nil {
			readedVal, err = item.ValueCopy(readedVal)
		}
		return err
	})
	return readedVal, err
}
func (db *Db) Add(key, data []byte) error {
	logger.Debug("adding new key-value to db")
	return db.PutKeyValue(key, data)
}

func (db *Db) List(prefixStr string) ([][]byte, error) {
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

func (db *Db) Delete(key []byte) error {
	logger.Debug("deleting key from db: ", string(key))
	execErr := db.instance.View(func(txn *badger.Txn) error {
		err := txn.Delete(key)
		if err != nil {
			logger.Error("failed to delete key from db: ", string(key), "error: ", err)
		}
		return err
	})
	return execErr
}

func GetInstance() *Db {
	logger.Debug("getting database instance")
	once.Do(func() {
		instance = &Db{}
		instance.Init()
	})
	return instance
}

func Hash(data string) string {
	// Use GenerateFromPassword to hash & salt pwd.
	// MinCost is just an integer constant provided by the bcrypt
	// package along with DefaultCost & MaxCost.
	// The cost can be any value you want provided it isn't lower
	// than the MinCost (4)
	hash, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("failed to hash user password")
		return ""
	}
	// GenerateFromPassword returns a byte slice so we need to
	// convert the bytes to a string and return it
	return string(hash)
}

func CompareHash(plainPwd, hash string) bool {
	// Since we'll be getting the hashed password from the DB it
	// will be a string so we'll need to convert it to a byte slice
	byteHash := []byte(hash)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	return err == nil
}

func Serialize(item interface{}) []byte {
	if item != nil {
		raw, err := json.Marshal(item)
		if err == nil {
			return raw
		}
	}
	return []byte{}
}

func Unserialize(data []byte, item interface{}) error {
	return json.Unmarshal(data, item)
}

func chownR(path string, uid, gid int, mode os.FileMode) error {
	return filepath.Walk(path, func(name string, info os.FileInfo, err error) error {
		if err == nil {
			err = os.Chown(name, uid, gid)
			err = os.Chmod(name, mode)
		}
		return err
	})
}
