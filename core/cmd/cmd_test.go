// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package cmd

import (
	"sync"
	"testing"
	"time"
)

func TestCmd(t *testing.T) {
	t.Run("test-server-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 50
		g.Add(total)
		notifier := make(chan error, total)
		for i := 0; i < total; i++ {
			go func() {
				RunServer(notifier)
				time.Sleep(time.Millisecond * 200)
				g.Done()
			}()
		}
		g.Wait()
	})
}
