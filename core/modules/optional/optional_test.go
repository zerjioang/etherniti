package optional

import "testing"

// example of functional programming in go
func TestOptional(t *testing.T) {
	t.Run("optional-1", func(t *testing.T) {
		t.Run("standard", func(t *testing.T) {
			a := "foo"
			b := "bar"
			c := -1
			if a == b {
				t.Log("does match")
				c = 1
			} else {
				t.Log("does not match")
				c = 0
			}
			t.Log("c value:", c)
		})
		t.Run("functional", func(t *testing.T) {
			a := "foo"
			b := "bar"
			c := -1
			Equal(a, b).Map(func() { t.Log("does match"); c = 1 }).OrElse(func() { t.Log("does not match"); c = 0 })
			t.Log("c value:", c)
		})
	})
}
