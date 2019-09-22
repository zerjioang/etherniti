// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package cmd_test

import (
	"github.com/zerjioang/etherniti/core/cmd"
	"sync"
	"testing"
	"time"
)

func TestCmd(t *testing.T) {
	t.Run("run-server", func(t *testing.T) {
		notifier := make(chan error, 1)
		cmd.RunServer(notifier)
		time.Sleep(time.Millisecond * 200)
	})
	t.Run("run-server-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 50
		g.Add(total)
		notifier := make(chan error, total)
		for i := 0; i < total; i++ {
			go func() {
				cmd.RunServer(notifier)
				time.Sleep(time.Millisecond * 200)
				g.Done()
			}()
		}
		g.Wait()
	})
}
