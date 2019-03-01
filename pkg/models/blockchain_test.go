package models

import (
	"bytes"
	"os"
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

func TestBlockchainFromDb(t *testing.T) {
	type args struct {
		chainName string
		filepath  string
	}
	existingChain := NewBlockchain()
	_ = existingChain.ToDb("existing chain", "chain.db")
	tests := []struct {
		name      string
		args      args
		wantChain *Blockchain
		wantErr   bool
	}{
		{"existing chain", args{"existing chain", "chain.db"}, existingChain, false},
		{"new chain", args{"new chain", "chain.db"}, nil, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotChain, err := BlockchainFromDb(tt.args.chainName, tt.args.filepath)
			if (err != nil) != tt.wantErr {
				t.Errorf("BlockchainFromDb() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantChain != nil && !gotChain.Equals(*tt.wantChain) {
				t.Errorf("BlockchainFromDb() = %v, want %v", gotChain, tt.wantChain)
			}
		})
	}
	t.Run("remove database", func(t *testing.T) {
		err := os.Remove("chain.db")
		if err != nil {
			t.Fatal("Couldn't remove database file", err)
		}
	})
}

func TestBlockchain_ToDb(t *testing.T) {
	type args struct {
		chainName string
		filepath  string
	}
	emptyChain := NewBlockchain()
	chainWithBlock := NewBlockchain()
	chainWithBlock.AddBlock([]byte("data"))
	existingChain := NewBlockchain()
	_ = existingChain.ToDb("existing chain", "chain.db")
	tests := []struct {
		name    string
		chain   *Blockchain
		args    args
		wantErr bool
	}{
		{"empty chain", emptyChain, args{"empty chain", "chain.db"}, false},
		{"chain with block", chainWithBlock, args{"chain with block", "chain.db"}, false},
		{"existing chain", existingChain, args{"existing chain", "chain.db"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			chain := tt.chain
			if err := chain.ToDb(tt.args.chainName, tt.args.filepath); (err != nil) != tt.wantErr {
				t.Errorf("Blockchain.ToDb() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
	t.Run("remove database", func(t *testing.T) {
		err := os.Remove("chain.db")
		if err != nil {
			t.Fatal("Couldn't remove database file", err)
		}
	})
}
