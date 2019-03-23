package models

import (
	"encoding/json"
	"errors"
	"os"

	bolt "go.etcd.io/bbolt"
)

var DbFilePath = "blockchain.db"
var BucketName = "blocks"
var LastBlockKey = "last"

// Blockchain represents a chain of blocks. Its only member is a slice of blocks
type Blockchain struct {
	Tail     []byte   `json:"tail"`
	Database *bolt.DB `json:"database"`
}

// NewBlockchain creates a new blockchain containing a genesis block. A genesis block is required, since every block
// requires a previous block hash. The genesis block is the only block with a nil value for its previous hash
func NewBlockchain(address string) (chain *Blockchain, err error) {
	err = os.Remove(DbFilePath)
	if err != nil {
		return
	}
	db, err := bolt.Open(DbFilePath, 0600, nil)
	if err != nil {
		return
	}
	var tail []byte
	err = db.Update(func(tx *bolt.Tx) error {
		genesis := NewGenesisBlock([]*Transaction{NewGenesisTransaction(address)})
		bucket, err := tx.CreateBucket([]byte(BucketName))
		bytes, err := json.Marshal(genesis)
		if err != nil {
			return err
		}
		err = bucket.Put(genesis.Hash, bytes)
		if err != nil {
			return err
		}
		err = bucket.Put([]byte(LastBlockKey), genesis.Hash)
		if err != nil {
			_ = bucket.Delete(genesis.Hash)
			return err
		}
		tail = genesis.Hash
		return nil
	})
	if err != nil {
		return
	}
	chain = &Blockchain{tail, db}
	return
}

func GetBlockchain() (chain *Blockchain, err error) {
	db, err := bolt.Open(DbFilePath, 0600, nil)
	if err != nil {
		return
	}
	var tail []byte
	err = db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		if bucket == nil {
			return errors.New("there is no existing blockchain")
		}
		tail = bucket.Get([]byte(LastBlockKey))
		return nil
	})
	if err != nil {
		return
	}
	chain = &Blockchain{tail, db}
	return
}

// AddBlock adds a block with the specified data to the chain using the hash of the latest block in the chain as
// previous block hash for the new block
func (chain *Blockchain) AddBlock(transactions []*Transaction) (err error) {
	block := NewBlock(transactions, chain.Tail)
	bytes, err := json.Marshal(block)
	if err != nil {
		return
	}
	err = chain.Database.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		err := bucket.Put(block.Hash, bytes)
		if err != nil {
			return err
		}
		err = bucket.Put([]byte(LastBlockKey), block.Hash)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return
	}
	chain.Tail = block.Hash
	return
}

func (chain *Blockchain) MineBlock(address string) (err error) {
	err = chain.AddBlock([]*Transaction{NewGenesisTransaction(address)})
	return
}

func (chain Blockchain) End() (block Block, err error) {
	return chain.Get(chain.Tail)
}

func (chain Blockchain) Get(hash []byte) (block Block, err error) {
	err = chain.Database.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		err := json.Unmarshal(bucket.Get(hash), block)
		if err != nil {
			return err
		}
		return nil
	})
	return
}

func (chain Blockchain) Range(action func(Block) bool) (err error) {
	nextKey := chain.Tail
	var block Block
	err = chain.Database.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		for {
			bytes := bucket.Get(nextKey)
			if bytes == nil {
				return nil
			}
			err := json.Unmarshal(bytes, block)
			if err != nil {
				return err
			}
			nextKey = block.PreviousHash
			if !action(block) {
				break
			}
		}
		return nil
	})
	return
}

func (chain Blockchain) String() (str string) {
	_ = chain.Range(func(block Block) bool {
		str = block.String() + "\n" + str
		return true
	})
	return
}
