package server

import "testing"

func TestNewEthernitiContext(t *testing.T) {
	t.Run("instantiate", func(t *testing.T) {
		_ = NewEthernitiContext()
	})
}
