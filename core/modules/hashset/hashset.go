// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package hashset

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"sync"

	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/zerjioang/etherniti/core/logger"
)

var (
	none struct{}
)

type HashSet struct {
	data map[string]struct{}
	lock *sync.RWMutex
}

func NewHashSet() HashSet {
	hs := HashSet{}
	hs.lock = new(sync.RWMutex)
	hs.data = map[string]struct{}{}
	return hs
}

func NewHashSetPtr() *HashSet {
	hs := NewHashSet()
	return &hs
}

func (set *HashSet) Add(item string) {
	set.lock.Lock()
	set.data[item] = none
	set.lock.Unlock()
}

func (set *HashSet) Size() int {
	set.lock.RLock()
	l := len(set.data)
	set.lock.RUnlock()
	return l
}

func (set *HashSet) Clear() {
	set.lock.Lock()
	set.data = map[string]struct{}{}
	set.lock.Unlock()
}

func (set HashSet) Contains(item string) bool {
	set.lock.RLock()
	_, contains := set.data[item]
	set.lock.RUnlock()
	return contains
}

func (set HashSet) MatchAny(item string) bool {
	set.lock.RLock()
	found := false
	for k := range set.data {
		found = strings.Contains(item, k)
		if found {
			break
		}
	}
	set.lock.RUnlock()
	return found
}

func (set *HashSet) Remove(item string) {
	set.lock.Lock()
	delete(set.data, item)
	set.lock.Unlock()
}

func (set HashSet) Count() int {
	set.lock.RLock()
	count := len(set.data)
	set.lock.RUnlock()
	return count
}

func (set *HashSet) LoadFromJsonArray(path string) {
	if path != "" {
		logger.Debug("loading hashset with json data")
		data, err := ioutil.ReadFile(path)
		if err != nil {
			logger.Error("could not read source data")
			return
		}
		var itemList []string
		unErr := json.Unmarshal(data, &itemList)
		if unErr != nil {
			logger.Error("could not unmarshal source data")
			return
		} else {
			set.LoadFromArray(itemList)
		}
	}
}

func (set *HashSet) LoadFromRaw(path string, splitChar string) {
	if path != "" {
		logger.Debug("loading hashset with raw data")
		data, err := ioutil.ReadFile(path)
		if err != nil {
			logger.Error("could not read source data")
			return
		}
		var itemList []string
		itemList = strings.Split(str.UnsafeString(data), splitChar)
		set.LoadFromArray(itemList)
	}
}

func (set *HashSet) LoadFromArray(data []string) {
	if data != nil {
		set.lock.Lock()
		for _, v := range data {
			set.data[v] = none
		}
		set.lock.Unlock()
	}
}
