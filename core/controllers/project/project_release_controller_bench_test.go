package project

import (
	"testing"

	"github.com/zerjioang/etherniti/core/logger"
)

func BenchmarkProjectReleaseController(b *testing.B) {
	b.Run("create-release-controller-struct", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = NewProjectReleaseController()
		}
	})
	b.Run("create-release-controller-ptr", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = NewProjectReleaseControllerPtr()
		}
	})
}
