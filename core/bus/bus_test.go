package bus

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBus(t *testing.T) {
	t.Run("get-shared-bus", func(t *testing.T) {
		b := SharedBus()
		assert.NotNil(t, b)
	})
	t.Run("get-shared-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 200
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				b := SharedBus()
				assert.NotNil(t, b)
				g.Done()
			}()
		}
		g.Wait()
	})
}
