package model

// Blockchain represents a chain of blocks. Its only member is a slice of blocks
type Blockchain struct {
	Blocks []*Block
}

// NewBlockchain creates a new blockchain containing a genesis block. A genesis block is required, since every block
// requires a previous block hash. The genesis block is the only block with a nil value for its previous hash
func NewBlockchain() *Blockchain {
	chain := &Blockchain{}
	chain.Blocks = append(chain.Blocks, NewBlock([]byte("Genesis Block"), nil))
	return chain
}

// AddBlock adds a block with the specified data to the chain using the hash of the latest block in the chain as
// previous block hash for the new block
func (chain *Blockchain) AddBlock(data []byte) {
	chain.Blocks = append(chain.Blocks, NewBlock(data, chain.Blocks[len(chain.Blocks)-1].Hash))
}
