// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package cmd

import (
	"sync"
	"testing"

	"github.com/zerjioang/etherniti/core/logger"
)

func BenchmarkCmd(b *testing.B) {

	b.Run("run-server-goroutines", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			var g sync.WaitGroup
			total := 50
			g.Add(total)
			notifier := make(chan error, total)
			for i := 0; i < total; i++ {
				go func() {
					RunServer(notifier)
					g.Done()
				}()
			}
			g.Wait()
		}
	})
}
