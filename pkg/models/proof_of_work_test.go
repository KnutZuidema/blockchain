package models

import (
	"bytes"
	"crypto/sha256"
	"math/big"
	"reflect"
	"strconv"
	"testing"
	"time"
)

func TestNewProofOfWork(t *testing.T) {
	type args struct {
		block *Block
	}
	block := &Block{Timestamp: time.Now(), Data: []byte("data"), PreviousHash: nil}
	target := big.NewInt(1)
	target.Lsh(target, uint(256-LeadingZeros))
	tests := []struct {
		name string
		args args
		want *ProofOfWork
	}{
		{"simple", args{block}, &ProofOfWork{target, block, big.NewInt(0)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewProofOfWork(tt.args.block); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewProofOfWork() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProofOfWork_createHash(t *testing.T) {
	type fields struct {
		target  *big.Int
		block   *Block
		Counter *big.Int
	}
	block := NewBlock([]byte("data"), nil)
	target := big.NewInt(1)
	target.Lsh(target, uint(256-LeadingZeros))
	hashValue := bytes.Join(
		[][]byte{
			block.Data,
			block.PreviousHash,
			[]byte(strconv.FormatInt(block.Timestamp.Unix(), 10)),
			big.NewInt(0).Bytes(),
		}, []byte{})
	hash := sha256.Sum256(hashValue)
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{"first", fields{target, block, big.NewInt(0)}, hash[:]},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pow := ProofOfWork{
				target:  tt.fields.target,
				block:   tt.fields.block,
				Counter: tt.fields.Counter,
			}
			if got := pow.createHash(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProofOfWork.createHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProofOfWork_Run(t *testing.T) {
	tests := []struct {
		name string
		pow  *ProofOfWork
	}{
		{"genesis block", NewProofOfWork(NewBlock([]byte("data"), nil))},
		{"generic block", NewProofOfWork(NewBlock([]byte("data"), []byte("previous")))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.pow.Run(); !tt.pow.Validate() {
				t.Error("hash value does not have required amount of leading zeros")
			}
		})
	}
}

func TestProofOfWork_Validate(t *testing.T) {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-LeadingZeros))
	mockBlock1 := &Block{}
	mockBlock2 := NewBlock([]byte("data"), nil)
	tests := []struct {
		name string
		pow  *ProofOfWork
		want bool
	}{
		{"false", &ProofOfWork{target, mockBlock1, big.NewInt(0)}, false},
		{"true", mockBlock2.ProofOfWork, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.pow.Validate(); got != tt.want {
				t.Errorf("ProofOfWork.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
