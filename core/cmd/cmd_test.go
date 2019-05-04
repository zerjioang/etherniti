// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package cmd

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCmd(t *testing.T) {
	t.Run("test-server", func(t *testing.T) {
		notifier := make(chan error, 1)
		RunServer(notifier)
		err := <-notifier
		assert.Nil(t, err)
	})
	t.Run("test-server-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 50
		g.Add(total)
		notifier := make(chan error, total)
		for i := 0; i < total; i++ {
			go func() {
				RunServer(notifier)
				<-notifier
				g.Done()
			}()
		}
		g.Wait()
	})
}
