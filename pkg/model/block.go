package model

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp    int64
	Data         []byte
	PreviousHash []byte
	Hash         []byte
}

func NewBlock(data []byte, previousHash []byte) *Block {
	block := &Block{
		Timestamp:    time.Now().Unix(),
		Data:         data,
		PreviousHash: previousHash,
	}
	block.createHash()
	return block
}

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

func (block *Block) String() string {
	return fmt.Sprintf("Previous Hash: %v\nHash: %v\n Data: %v",
		block.PreviousHash, block.Hash, block.Data)
}