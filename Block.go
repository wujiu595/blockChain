package main

import (
	"bytes"
	"encoding/binary"
	"time"
)

var genesisInfo = "MyFirstBlock"

type Block struct {
	Version uint64
	//上一个区块的Hash
	PrevBlockHash []byte
	//默克尔根
	MerkelRoot []byte
	//时间
	TimeStamp uint64
	//难度值
	Bits uint64
	//工作量证明
	Nonce uint64
	//数据
	Data []byte
	//当前区块hash本应该存在，为方便计算使用
	Hash []byte
}

//新建区块
func NewBlock(data string,prevBlockHas []byte)*Block  {
	block:=Block{
		Version:0.0,
		PrevBlockHash:prevBlockHas,
		MerkelRoot:[]byte{},
		TimeStamp:uint64(time.Now().Unix()),
		Bits:0,
		Nonce :0,
		Data:[]byte(data),
	}
	pow:=NewProofOfWork(block)
	nonce,hash:=pow.Run()
	block.Nonce = nonce
	block.Hash = hash
	return &block
}
//
func (this *Block)ToBytes()[]byte  {
	return []byte{}
}


//字节化
func Uint2Bytes(src interface{})[]byte{
	bin_buf := bytes.NewBuffer([]byte{})
	binary.Write(bin_buf, binary.BigEndian, src)
	return bin_buf.Bytes()
}
