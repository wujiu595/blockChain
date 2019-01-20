package main

import (
	"crypto/sha256"
)

var genesisInfo = "MyFirstBlock"

type Block struct {
	PrevBlockHash []byte
	Data []byte
	//当前区块hash本应该存在，为方便计算使用
	Hash []byte
}

func NewBlock(data string,prevBlockHas []byte)*Block  {
	block:=Block{
		PrevBlockHash:prevBlockHas,
		Data:[]byte(data),
	}
	block.setHas()
	return &block
}


func (b *Block)setHas()  {
	var blockInfo []byte
	blockInfo = append(blockInfo,b.PrevBlockHash...)
	blockInfo = append(blockInfo,b.Data...)
	Hash:=sha256.Sum256(blockInfo)
	b.Hash = Hash[:]
}
