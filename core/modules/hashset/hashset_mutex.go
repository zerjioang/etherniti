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

type HashSetMutex struct {
	set  HashSet
	lock *sync.RWMutex
}

func NewHashSetMutex() HashSetMutex {
	hs := HashSetMutex{}
	hs.lock = new(sync.RWMutex)
	hs.set = HashSet{}
	return hs
}

func NewHashSetMutexPtr() *HashSetMutex {
	hs := NewHashSetMutex()
	return &hs
}

func (s *HashSetMutex) Add(item string) {
	s.lock.Lock()
	s.set[item] = none
	s.lock.Unlock()
}

func (s *HashSetMutex) Clear() {
	s.lock.Lock()
	s.set = HashSet{}
	s.lock.Unlock()
}

func (s HashSetMutex) Contains(item string) bool {
	s.lock.RLock()
	_, contains := s.set[item]
	s.lock.RUnlock()
	return contains
}

func (s HashSetMutex) MatchAny(item string) bool {
	s.lock.RLock()
	found := false
	for k := range s.set {
		found = strings.Contains(item, k)
		if found {
			break
		}
	}
	s.lock.RUnlock()
	return found
}

func (s *HashSetMutex) Remove(item string) {
	s.lock.Lock()
	delete(s.set, item)
	s.lock.Unlock()
}

func (s *HashSetMutex) Size() int {
	s.lock.RLock()
	l := len(s.set)
	s.lock.RUnlock()
	return l
}

func (s *HashSetMutex) LoadFromJsonArray(path string) {
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
			s.LoadFromArray(itemList)
		}
	}
}

func (s *HashSetMutex) LoadFromRaw(path string, splitChar string) {
	if path != "" {
		logger.Debug("loading hashset with raw data")
		data, err := ioutil.ReadFile(path)
		if err != nil {
			logger.Error("could not read source data")
			return
		}
		var itemList []string
		itemList = strings.Split(str.UnsafeString(data), splitChar)
		s.LoadFromArray(itemList)
	}
}

// content loaded via this method will make to allocate data on the heap
func (s *HashSetMutex) LoadFromArray(data []string) {
	if data != nil {
		s.lock.Lock()
		for _, v := range data {
			s.set[v] = none
		}
		s.lock.Unlock()
	}
}
