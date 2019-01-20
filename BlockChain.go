package main

type BlockChain struct {
	Blocks []*Block
}

func NewBlockChain()*BlockChain  {
	genesisBlock:=NewBlock(genesisInfo,[]byte{})
	bc:=BlockChain{
		Blocks:[]*Block{genesisBlock},
	}
	return &bc
}

func (bc *BlockChain)AddBlockChain(data string)  {
	newBlock := NewBlock(data,bc.Blocks[len(bc.Blocks)-1].Hash)
	bc.Blocks=append(bc.Blocks, newBlock)
}