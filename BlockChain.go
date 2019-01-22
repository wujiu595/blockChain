package main

import (
	"github.com/boltdb/bolt"
	"log"
)

const blockChainFile="./Block.db"
const blockBucket="blockBucket"
const lastHashKey="lastHashKey"

type BlockChain struct {
	db *bolt.DB
	tail []byte
}
//新建区块链
func NewBlockChain()*BlockChain  {
	db,err:=bolt.Open(blockChainFile,0600,nil)
	bc:=BlockChain{}
	if err!=nil{
		log.Fatal(err)
		return nil
	}
	err=db.Update(func(tx *bolt.Tx) error {
		b:=tx.Bucket([]byte(blockBucket))
		//第一次调用
		if b==nil{
			//新建存储桶
			b,err=tx.CreateBucket([]byte(blockBucket))
			if err!=nil{
				return err
			}
			//新建创世块
			genesisBlock:=NewBlock(genesisInfo,[]byte{})
			//初始化
			err=b.Put(genesisBlock.Hash,genesisBlock.ToBytes())
			if err!=nil{
				log.Fatal(err)
				return nil
			}
			err=b.Put([]byte(lastHashKey),genesisBlock.Hash)
			if err!=nil{
				log.Fatal(err)
				return nil
			}
			bc.tail=genesisBlock.Hash
			return nil
		}
		bc.tail=b.Get([]byte(lastHashKey))
		return nil
	})
	if err!=nil{
		log.Fatal(err)
		return nil
	}
	bc.db=db
	return &bc
}
//添加区块到链中
//func (bc *BlockChain)AddBlockChain(data string)  {
//	newBlock := NewBlock(data,bc.Blocks[len(bc.Blocks)-1].Hash)
//	bc.Blocks=append(bc.Blocks, newBlock)
//}