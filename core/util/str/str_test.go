// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package str

import (
	"testing"
)

func TestGetJsonBytes(t *testing.T) {
	t.Run("get-bytes-nil", func(t *testing.T) {
		GetJsonBytes(nil)
	})
}

func TestToLowerAscii(t *testing.T) {
	t.Run("ToLowerAscii", func(t *testing.T) {
		val := "Hello World, This is AWESOME"
		converted := ToLowerAscii(val)
		t.Log(val)
		t.Log(converted)
		if converted != "hello world, this is awesome" {
			t.Error("failed to lowercase")
		}
	})
	t.Run("ToLowerAscii-ua", func(t *testing.T) {
		val := "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:61.0) Gecko/20100101 Firefox/61.0"
		converted := ToLowerAscii(val)
		t.Log(val)
		t.Log(converted)
		if converted != "mozilla/5.0 (x11; ubuntu; linux x86_64; rv:61.0) gecko/20100101 firefox/61.0" {
			t.Error("failed to lowercase")
		}
	})
}
