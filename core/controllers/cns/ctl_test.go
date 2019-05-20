package cns

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCnsController(t *testing.T) {
	t.Run("create-cns-controller", func(t *testing.T) {
		pc := NewContractNameServiceController()
		assert.NotNil(t, pc)
	})
}
