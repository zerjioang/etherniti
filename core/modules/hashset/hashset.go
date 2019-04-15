// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package hashset

import (
	"encoding/json"
	"github.com/zerjioang/etherniti/core/logger"
	"io/ioutil"
	"strings"
)

var (
	none *struct{}
)

type HashSet struct {
	data map[string]*struct{}
}

func NewHashSet() *HashSet {
	hs := new(HashSet)
	hs.Clear()
	return hs
}

func (set *HashSet) Add(item string) {
	set.data[item] = none
}

func (set *HashSet) Clear() {
	set.data = make(map[string]*struct{})
}

func (set HashSet) Contains(item string) bool {
	_, contains := set.data[item]
	return contains
}

func (set HashSet) MatchAny(item string) bool {
	for k := range set.data {
		status := strings.Contains(item, k)
		if status {
			return true
		}
	}
	return false
}

func (set *HashSet) Remove(item string) {
	delete(set.data, item)
}

func (set HashSet) Count() int {
	return len(set.data)
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
