package handlers

import "testing"

func TestUnixSocketListener(t *testing.T) {
	t.Run("instantiation", func(t *testing.T) {
		NewUnixSocketDeployer()
	})
}
