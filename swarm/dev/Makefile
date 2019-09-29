.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    geth-timing/cmd/swarm
	go build -o bin/geth     geth-timing/cmd/geth
	go build -o bin/bootnode geth-timing/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v geth-timing/swarm/...
