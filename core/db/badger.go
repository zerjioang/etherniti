// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package db

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"

	"github.com/dgraph-io/badger"
	"github.com/zerjioang/etherniti/core/logger"
	"golang.org/x/crypto/bcrypt"
)

var (
	defaultConfig   = badger.DefaultOptions
	home            = os.Getenv("HOME")
	baseData        = home + "/.etherniti/"
	duplicateKeyErr = errors.New("duplicate key found on database. cannot store")
)

type Db struct {
	instance *badger.DB
}

var instance *Db
var once sync.Once
var uid = os.Getuid()
var gid = os.Getgid()

func init() {
	logger.Debug("loading db module data")
	err := createData(baseData + "data")
	if err != nil {
		logger.Error("failed to create shared database dir:", err)
	}
}

func createData(path string) error {
	logger.Debug("creating db path")
	defaultConfig.Dir = path
	defaultConfig.ValueDir = path
	logger.Debug("creating dir: ", path)
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		logger.Error("failed to create database dir:", err)
	} else {
		logger.Debug("setting dir permissions and ownership: ", path)
		permErr := chownR(path, uid, gid, os.ModePerm)
		if permErr != nil {
			logger.Error("failed to set permissions: ", permErr)
		}
		return permErr
	}
	return err
}

func NewCollection(name string) (*Db, error) {
	logger.Debug("creating new db collection")
	collection := new(Db)
	err := createData(baseData + name)
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
	err := db.instance.View(func(txn *badger.Txn) error {
		// Your code here…
		return nil
	})
	return err
}

func (db *Db) PutKeyValue(key []byte, value []byte) error {
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
	err := db.instance.Update(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err == nil && item != nil {
			return duplicateKeyErr
		} else {
			return txn.Set(key, value)
		}
	})
	return err
}

func (db *Db) GetKeyValue(key []byte) ([]byte, error) {
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
	return db.PutKeyValue(key, data)
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
