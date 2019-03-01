package models

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
	block := NewBlock(data, chain.Blocks[len(chain.Blocks)-1].Hash)
	if !block.ProofOfWork.Validate() {
		return
	}
	chain.Blocks = append(chain.Blocks, block)
}

func (chain *Blockchain) String() string {
	str := "Blockchain:\n\n"
	for _, block := range chain.Blocks {
		str += block.String() + "\n\n"
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
