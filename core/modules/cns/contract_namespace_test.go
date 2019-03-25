package cns

import (
	"fmt"
	"runtime"
)

func ExampleContractNameSystem() {
	fmt.Println("Hello, playground")

	fmt.Println(runtime.NumCPU())

	cns := NewContractNameSystem()

	contract := ContractInfo{}
	contract.name = "test"
	contract.description = "this is a demo contract"
	contract.address = "0xf17f52151EbEF6C7334FAD080c5704D77216b732"
	contract.version = "1.2"

	cns.Register(contract)

	fmt.Println(cns.Resolve("test-1.2"))
	fmt.Println(cns.Resolve("test-02"))
}