package entropy

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestAvailableEntropy(t *testing.T) {
	t.Run("initial-entropy", func(t *testing.T) {
		v := InitialEntropy()
		t.Log(v)
		assert.True(t, v != -1)
	})
	t.Run("initial-entropy-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 50
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				v := InitialEntropy()
				assert.True(t, v != -1)
				g.Done()
			}()
		}
		g.Wait()
	})
	t.Run("current-entropy", func(t *testing.T) {
		v := AvailableEntropy()
		t.Log(v)
		assert.True(t, v != -1)
	})
	t.Run("current-entropy-goroutines", func(t *testing.T) {
		var g sync.WaitGroup
		total := 50
		g.Add(total)
		for i := 0; i < total; i++ {
			go func() {
				v := AvailableEntropy()
				assert.True(t, v != -1)
				g.Done()
			}()
		}
		g.Wait()
	})
}
