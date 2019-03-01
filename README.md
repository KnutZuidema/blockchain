# Blockchain

Implementation of a generic blockchain in golang from scratch

Built following the [Building Blockchain in Go](https://jeiwan.cc/posts/building-blockchain-in-go-part-1/)
tutorial by Ivan Kuznetsov

## Requirements

BoltDB - https://github.com/etcd-io/bbolt

## Building

`go build cmd/blockchain.go`

## Usage

##### Adding a Block
`blockchain add -data "some data"`

##### Printing the Chain
`blockchain print`