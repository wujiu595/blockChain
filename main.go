package main

import (
	"time"
)

func main()  {
	//blockChain:=NewBlockChain()
	//blockChain.AddBlockChain("nihao")
	//blockChain.AddBlockChain("hangtou")
	//blockChain.AddBlockChain("shanghai")
	//blockChain.AddBlockChain("beijing")
	//blockChain.AddBlockChain("tianjin")
/*	for i,block:= range blockChain.Blocks{
		fmt.Printf("============%d===========\n",i)
		fmt.Printf("Data:%s\n",block.Data)
		fmt.Printf("PrevBlockHash:%x\n",block.PrevBlockHash)
		fmt.Printf("Hash:%x\n",block.Hash)
		fmt.Printf("Time:%v\n",FormatTime(block.TimeStamp))
		fmt.Printf("Nonce:%v\n",block.Nonce)
		pow:=NewProofOfWork(*block)
		fmt.Println(pow.IsValid())
	}*/
}

func FormatTime(timeStamp uint64)string  {
	datetime := time.Unix(int64(timeStamp), 0).Format("2006-01-02 15:04:05")
	return datetime
}