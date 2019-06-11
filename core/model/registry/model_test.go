package registry

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRegistryModel(t *testing.T) {
	t.Run("create-model", func(t *testing.T) {
		contract := NewEmptyRegistry()
		assert.NotNil(t, contract)

		contract.Name = "test"
		contract.Description = "this is a demo contract"
		contract.Address = "0xf17f52151EbEF6C7334FAD080c5704D77216b732"
	})
}

func ExampleRegistry() {
	fmt.Println("Hello, playground")

	fmt.Println(runtime.NumCPU())

	contract := NewEmptyRegistry()

	contract.Name = "test"
	contract.Description = "this is a demo contract"
	contract.Address = "0xf17f52151EbEF6C7334FAD080c5704D77216b732"
}
