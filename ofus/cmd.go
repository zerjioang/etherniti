package main

import (
	"os"

	"github.com/zerjioang/etherniti/ofus/lib"
)

func main() {
	gopath := os.Getenv("GOPATH")
	scan := gopath + "/src/github.com/zerjioang/etherniti"
	mainPath := scan + "/main.go"
	of := lib.NewOfuscator()
	of.Start(scan, mainPath)
}
