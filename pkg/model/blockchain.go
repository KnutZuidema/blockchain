package model

type Blockchain struct {
	Blocks []*Block
}

func NewBlockchain() *Blockchain {
	chain := &Blockchain{}
	chain.Blocks = append(chain.Blocks, NewBlock([]byte("Genesis Block"), nil))
	return chain
}

func (chain *Blockchain) AddBlock(data []byte) {
	chain.Blocks = append(chain.Blocks, NewBlock(data, chain.Blocks[len(chain.Blocks) - 1].Hash))
}
