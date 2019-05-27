package project

import (
	"testing"

	"github.com/zerjioang/etherniti/core/logger"
)

func BenchmarkProjectInteractionController(b *testing.B) {
	b.Run("create-interaction-controller-struct", func(b *testing.B) {
		logger.Enabled(false)
		p := NewProjectController()
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = NewProjectInteractionController(p)
		}
	})
	b.Run("create-interaction-controller-ptr", func(b *testing.B) {
		logger.Enabled(false)
		p := NewProjectControllerPtr()
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = NewProjectInteractionControllerPtr(p)
		}
	})
}
