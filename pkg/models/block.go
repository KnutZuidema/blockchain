package models

import (
	"fmt"
	"time"
)

// Block is a block in a blockchain. It stores a creation timestamp, its own hash, the hash of the previous block in
// the chain, and some generic data
type Block struct {
	Timestamp    time.Time
	Data         []byte
	PreviousHash []byte
	Hash         []byte
	ProofOfWork  *pow.ProofOfWork
}

// NewBlock creates a new block with specified data and a specified hash of a previous block
func NewBlock(data []byte, previousHash []byte) *Block {
	block := &Block{
		Timestamp:          time.Now().UTC(),
		Data:               data,
		PreviousHash:       previousHash,
		ProofOfWorkCounter: big.NewInt(0),
	}
	block.createHash()
	return block
}

// createHash creates a hash for a block using SHA256. The hash is created using a proof of work method
func (block *Block) createHash() {
	pow := NewProofOfWork(block)
	block.Hash = pow.Run()
}

// String is a string representation of a block
func (block *Block) String() string {
	return fmt.Sprintf("Previous Hash: %x\nHash: %x\nTimestamp: %v\nData: %s",
		block.PreviousHash, block.Hash, block.Timestamp, block.Data)
}
