package main

import "fmt"

func main()  {
	blockChain:=NewBlockChain()
	blockChain.AddBlockChain("nihao")
	blockChain.AddBlockChain("hangtou")
	blockChain.AddBlockChain("shanghai")
	blockChain.AddBlockChain("beijing")
	blockChain.AddBlockChain("tianjin")
	for i,block:= range blockChain.Blocks{
		fmt.Printf("============%d===========\n",i)
		fmt.Printf("Data:%s\n",block.Data)
		fmt.Printf("PrevBlockHash:%x\n",block.PrevBlockHash)
		fmt.Printf("Hash:%x\n",block.Hash)
	}
}