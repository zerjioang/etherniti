// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package hashset

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"sync"

	"github.com/zerjioang/etherniti/core/logger"
)

var (
	none *struct{}
)

type HashSet struct {
	data map[string]*struct{}
	lock *sync.Mutex
}

func NewHashSet() *HashSet {
	hs := new(HashSet)
	hs.lock = new(sync.Mutex)
	hs.Clear()
	return hs
}

func (set *HashSet) Add(item string) {
	set.lock.Lock()
	set.data[item] = none
	set.lock.Unlock()
}

func (set *HashSet) Clear() {
	set.lock.Lock()
	set.data = make(map[string]*struct{})
	set.lock.Unlock()
}

func (set HashSet) Contains(item string) bool {
	set.lock.Lock()
	_, contains := set.data[item]
	set.lock.Unlock()
	return contains
}

func (set HashSet) MatchAny(item string) bool {
	set.lock.Lock()
	found := false
	for k := range set.data {
		found = strings.Contains(item, k)
		if found {
			break
		}
	}
	set.lock.Unlock()
	return found
}

func (set *HashSet) Remove(item string) {
	set.lock.Lock()
	delete(set.data, item)
	set.lock.Unlock()
}

func (set HashSet) Count() int {
	set.lock.Lock()
	count := len(set.data)
	set.lock.Unlock()
	return count
}

func (set *HashSet) LoadFromJsonArray(path string) {
	if path != "" {
		logger.Debug("loading hashset with data")
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

func (set *HashSet) LoadFromArray(data []string) {
	if data != nil {
		for _, v := range data {
			set.Add(v)
		}
	}
}
