package main

import (
	"crypto/sha256"
	"fmt"
	"math/big"
)

type ProofOfWork struct {
	block Block
	targetStr big.Int
}

func NewProofOfWork(block Block) *ProofOfWork {
	targetStr:="0000100000000000000000000000000000000000000000000000000000000000"
	bigIntTemp:=big.Int{}
	bigIntTemp.SetString(targetStr,16)
	pow:=ProofOfWork{
		block:block,
		targetStr:bigIntTemp,
	}

	return &pow
}

func (this *ProofOfWork)Run() (uint64,[]byte){
	b:=this.block
	var nonce uint64
	var hash [32]byte
	for{
		fmt.Printf("%x\n",hash)
		var blockInfo []byte
		blockInfo=append(blockInfo,Uint2Bytes(b.Version)...)
		blockInfo = append(blockInfo,b.PrevBlockHash...)
		blockInfo = append(blockInfo, b.MerkelRoot...)
		blockInfo = append(blockInfo, Uint2Bytes(b.TimeStamp)...)
		blockInfo = append(blockInfo, Uint2Bytes(b.Bits)...)
		blockInfo = append(blockInfo, Uint2Bytes(nonce)...)
		blockInfo = append(blockInfo,b.Data...)
		hash=sha256.Sum256(blockInfo)
		var bigIntTmp big.Int
		bigIntTmp.SetBytes(hash[:])
		if bigIntTmp.Cmp(&this.targetStr) ==-1{
			break
		}
		nonce++
	}
	return nonce,hash[:]
}