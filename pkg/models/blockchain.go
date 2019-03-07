package models

import (
	"encoding/json"
	bolt "go.etcd.io/bbolt"
)

// Blockchain represents a chain of blocks. Its only member is a slice of blocks
type Blockchain struct {
	Blocks []*Block `json:"blocks"`
}

// NewBlockchain creates a new blockchain containing a genesis block. A genesis block is required, since every block
// requires a previous block hash. The genesis block is the only block with a nil value for its previous hash
func NewBlockchain() *Blockchain {
	chain := &Blockchain{}
	chain.Blocks = append(chain.Blocks, NewBlock([]*Transaction{NewGenesisTransaction("receiver", "genesis")}, nil))
	return chain
}

func BlockchainFromDb(chainName, filepath string) (chain Blockchain, err error) {
	db, err := bolt.Open(filepath, 0600, nil)
	if err != nil {
		return
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("chains"))
		if err != nil {
			return err
		}
		bytes := bucket.Get([]byte(chainName))
		if bytes == nil {
			chain = *NewBlockchain()
		} else {
			err := json.Unmarshal(bytes, &chain)
			if err != nil {
				return err
			}
		}
		return nil
	})
	return
}

// AddBlock adds a block with the specified data to the chain using the hash of the latest block in the chain as
// previous block hash for the new block
func (chain *Blockchain) AddBlock(transactions []*Transaction) {
	block := NewBlock(transactions, chain.Blocks[len(chain.Blocks)-1].Hash)
	if !NewProofOfWork(block).Validate() {
		return
	}
	chain.Blocks = append(chain.Blocks, block)
}

func (chain Blockchain) ToDb(chainName, filepath string) error {
	bytes, err := json.Marshal(chain)
	if err != nil {
		return err
	}
	db, err := bolt.Open(filepath, 0600, nil)
	if err != nil {
		return err
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		bucket, err := tx.CreateBucketIfNotExists([]byte("chains"))
		if err != nil {
			return err
		}
		err = bucket.Put([]byte(chainName), bytes)
		if err != nil {
			return err
		}
		return nil
	})
	return nil
}

func (chain Blockchain) String() string {
	str := ""
	for _, block := range chain.Blocks {
		str += block.String() + "\n"
	}
	return str
}

func (chain Blockchain) Equals(otherChain Blockchain) bool {
	for index := range chain.Blocks {
		if !chain.Blocks[index].Equals(*otherChain.Blocks[index]) {
			return false
		}
	}
	return true
}
