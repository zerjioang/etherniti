// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package random_test

import (
	"github.com/zerjioang/etherniti/thirdparty/gommon/random"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	assert.Len(t, random.String(32), 32)
	r := random.New()
	assert.Regexp(t, regexp.MustCompile("[0-9]+$"), r.String(8, random.Numeric))
	t.Run("unsafe-implementation", func(t *testing.T) {
		r := random.RandomUUID32()
		t.Log(r)
	})
	t.Run("shared-bytes-instance", func(t *testing.T) {
		r := random.RandomUUID32Shared()
		t.Log(r)
	})
}
