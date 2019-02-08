.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/zerjioang/etherniti/core/modules/ethfork/cmd/swarm
	go build -o bin/geth     github.com/zerjioang/etherniti/core/modules/ethfork/cmd/geth
	go build -o bin/bootnode github.com/zerjioang/etherniti/core/modules/ethfork/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/zerjioang/etherniti/core/modules/ethfork/swarm/...
