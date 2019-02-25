package model

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

// Block is a block in a blockchain. It stores a creation timestamp, its own hash, the hash of the previous block in
// the chain, and some generic data
type Block struct {
	Timestamp    int64
	Data         []byte
	PreviousHash []byte
	Hash         []byte
}

// NewBlock creates a new block with specified data and a specified hash of a previous block
func NewBlock(data []byte, previousHash []byte) *Block {
	block := &Block{
		Timestamp:    time.Now().Unix(),
		Data:         data,
		PreviousHash: previousHash,
	}
	block.createHash()
	return block
}

// createHash creates a hash for a block using SHA256. The hash consists of the timestamp, the data and the hash
// of the previous block
func (block *Block) createHash() {
	hashValue := bytes.Join(
		[][]byte{
			block.Data,
			block.PreviousHash,
			[]byte(strconv.FormatInt(block.Timestamp, 10)),
		}, []byte{})
	hash := sha256.Sum256(hashValue)
	block.Hash = hash[:]
}

// String is a string representation of a block
func (block *Block) String() string {
	return fmt.Sprintf("Previous Hash: %v\nHash: %v\n Data: %v",
		block.PreviousHash, block.Hash, block.Data)
}
