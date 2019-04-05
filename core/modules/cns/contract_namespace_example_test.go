// Copyright etherniti
// SPDX-License-Identifier: Apache License 2.0

package cns

import (
	"fmt"
	"runtime"
	"testing"
)

func ExampleContractNameSystem() {
	fmt.Println("Hello, playground")

	fmt.Println(runtime.NumCPU())

	cns := NewContractNameSystem()

	contract := ContractInfo{}
	contract.Name = "test"
	contract.Description = "this is a demo contract"
	contract.Address = "0xf17f52151EbEF6C7334FAD080c5704D77216b732"
	contract.Version = "1.2"

	cns.Register(contract)

	fmt.Println(cns.Resolve("test-1.2"))
	fmt.Println(cns.Resolve("test-02"))
}

func TestExampleContractNameSystem(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ExampleContractNameSystem()
		})
	}
}
