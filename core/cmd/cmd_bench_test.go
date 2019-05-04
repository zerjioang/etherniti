// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package cmd

import (
	"sync"
	"testing"
)

func BenchmarkCmd(b *testing.B) {
	b.Run("run-server", func(b *testing.B) {
		// pre-required data
		notifier := make(chan error, 1)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			RunServer(notifier)
			<-notifier
		}
	})
	b.Run("run-server-goroutines", func(b *testing.B) {
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
					<-notifier
					g.Done()
				}()
			}
			g.Wait()
		}
	})
}
