package main

// Blockchain contains all blocks
type Blockchain struct {
	blocks []*Block
}

// AddBlock attaches new block to chain
func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// NewBlockchain creates a branh new chain
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}

// NewGenesisBlock creates the Genesisblock
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}
