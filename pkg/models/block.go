package models

import (
	"bytes"
	"fmt"
	"math/big"
	"time"
)

// Block is a block in a blockchain. It stores a creation timestamp, its own hash, the hash of the previous block in
// the chain, and some generic data
type Block struct {
	Timestamp          time.Time `json:"timestamp"`
	Data               []byte    `json:"data"`
	PreviousHash       []byte    `json:"previous_hash"`
	Hash               []byte    `json:"hash"`
	ProofOfWorkCounter *big.Int  `json:"proof_of_work_counter"`
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
func (block Block) String() string {
	return fmt.Sprintf("Previous Hash: %x\nHash: %x\nTimestamp: %v\nData: %s",
		block.PreviousHash, block.Hash, block.Timestamp, block.Data)
}

func (block Block) Equals(otherBlock Block) bool {
	return bytes.Equal(NewProofOfWork(&block).createHash(), NewProofOfWork(&otherBlock).createHash())
}
