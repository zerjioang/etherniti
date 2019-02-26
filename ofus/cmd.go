package main

import (
	"github.com/zerjioang/etherniti/ofus/lib"
	"os"
)

func main() {
	gopath := os.Getenv("GOPATH")
	scan := gopath+"/src/github.com/zerjioang/etherniti"
	lib.NewOfuscator().Start(scan)
}
