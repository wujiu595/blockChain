package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math/big"
)

type ProofOfWork struct {
	block Block
	targetStr big.Int
}

//初始化工作量证明
func NewProofOfWork(block Block) *ProofOfWork {
	targetStr:="0010000000000000000000000000000000000000000000000000000000000000"
	bigIntTemp:=big.Int{}
	bigIntTemp.SetString(targetStr,16)
	pow:=ProofOfWork{
		block:block,
		targetStr:bigIntTemp,
	}

	return &pow
}
//运行工作量证明
func (this *ProofOfWork)Run() (uint64,[]byte){
	var nonce uint64
	var hash [32]byte
	for{
		var bigIntTmp big.Int
		//进行sha256计算，获得hash值
		hash=sha256.Sum256(this.PrepareData(nonce))
		fmt.Printf("%x\n",hash)
		bigIntTmp.SetBytes(hash[:])
		if bigIntTmp.Cmp(&this.targetStr) ==-1{
			break
		}
		nonce++
	}
	return nonce,hash[:]
}

//构建前一个区块
func (this *ProofOfWork)PrepareData(nonce uint64)[]byte{
	b:=this.block
	tempByte:=[][]byte{
		Uint2Bytes(b.Version),
		b.PrevBlockHash,
		b.MerkelRoot,
		Uint2Bytes(b.TimeStamp),
		Uint2Bytes(b.Bits),
		Uint2Bytes(nonce),
		b.Data,
	}
	blockInfo:=bytes.Join(tempByte,[]byte(""))
	return blockInfo
}
