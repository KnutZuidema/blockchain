package models

import (
	"bytes"
	"testing"
)

func TestNewBlockchain(t *testing.T) {
	t.Run("simple constructor", func(t *testing.T) {
		got := NewBlockchain()
		if len(got.Blocks) != 1 {
			t.Error("no genesis block")
		}
		if !bytes.Equal(got.Blocks[0].Data, []byte("Genesis Block")) {
			t.Errorf("did not set data of genesis block correctly, got %v, want %v", got.Blocks[0].Data,
				[]byte("Genesis Block"))
		}
	})
}

func TestBlockchain_AddBlock(t *testing.T) {
	tests := []struct {
		name string
		args []byte
	}{
		{"first", []byte("0")},
		{"second", []byte("1")},
		{"third", []byte("2")},
	}
	chain := NewBlockchain()
	for index, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			beforeSize := len(chain.Blocks)
			chain.AddBlock(tt.args)
			if !(len(chain.Blocks) == beforeSize+1) {
				t.Error("did not add element to chain")
			}
			if !(bytes.Equal(chain.Blocks[index].Hash, chain.Blocks[index+1].PreviousHash)) {
				t.Error("hashes do not match")
			}
		})
	}
}
