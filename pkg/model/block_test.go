package model

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"testing"
)

func TestNewBlock(t *testing.T) {
	block1 := NewBlock([]byte("first"), []byte(""))
	block2 := NewBlock([]byte("second"), block1.Hash)
	t.Run("new block", func(tt *testing.T) {
		if !(bytes.Equal(block1.Data, []byte("first"))) {
			t.Errorf("did not set data correctly, got %v, want %v", block1.Data, []byte("first"))
		}
		if !(bytes.Equal(block1.Hash, block2.PreviousHash)) {
			t.Errorf("did not set previous hash correctly, %v != %v", block1.Hash, block2.PreviousHash)
		}
	})
}

func TestBlock_createHash(t *testing.T) {
	block := NewBlock([]byte("hash"), []byte(""))
	hashValue := bytes.Join(
		[][]byte{
			[]byte("hash"),
			[]byte(""),
			[]byte(strconv.FormatInt(block.Timestamp.Unix(), 10)),
			block.ProofOfWork.Counter.Bytes(),
		}, []byte{})
	hash := sha256.Sum256(hashValue)
	hashValue = hash[:]
	t.Run("creating hash", func(tt *testing.T) {
		if !(bytes.Equal(hashValue, block.Hash)) {
			t.Errorf("incorrectly set hash value, got %v, want %v", block.Hash, hashValue)
		}
	})
}

func TestBlock_String(t *testing.T) {
	t.Run("string", func(t *testing.T) {
		block := NewBlock([]byte("str"), []byte(""))
		str := fmt.Sprintf("Previous Hash: %x\nHash: %x\nTimestamp: %v\nData: %s",
			block.PreviousHash, block.Hash, block.Timestamp, block.Data)
		if got := block.String(); got != str {
			t.Errorf("Block.String() = %v, want %v", got, str)
		}
	})
}
