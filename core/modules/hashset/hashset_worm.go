// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package hashset

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/zerjioang/etherniti/core/util/str"

	"github.com/zerjioang/etherniti/core/logger"
)

// HashSet WORM (Write Once Read Many)
// is an unsafe data structure in where data is loaded once
// for example, at init() functions an it is readed, after loading many times
// avoiding further modifications after being loaded.
type HashSetWORM struct {
	set HashSet
}

func NewHashSetWORM() HashSetWORM {
	hs := HashSetWORM{}
	hs.set = NewHashSet()
	return hs
}

func NewHashSetWORMPtr() *HashSetWORM {
	hs := NewHashSetWORM()
	return &hs
}

func (s *HashSetWORM) Add(item string) {
	s.set[item] = none
}

func (s *HashSetWORM) Clear() {
	s.set = NewHashSet()
}

func (s HashSetWORM) Contains(item string) bool {
	_, found := s.set[item]
	return found
}

func (s HashSetWORM) MatchAny(item string) bool {
	found := false
	for k := range s.set {
		found = strings.Contains(item, k)
		if found {
			break
		}
	}
	return found
}

func (s *HashSetWORM) Remove(item string) {
	delete(s.set, item)
}

func (s *HashSetWORM) Size() int {
	l := len(s.set)
	return l
}

func (s *HashSetWORM) LoadFromJsonArray(path string) {
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

func (s *HashSetWORM) LoadFromRaw(path string, splitChar string) {
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
func (s *HashSetWORM) LoadFromArray(data []string) {
	if data != nil {
		for _, v := range data {
			s.set[v] = none
		}
	}
}
