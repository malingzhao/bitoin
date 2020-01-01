package BLC

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

//区块的结构体
type Block struct {
	Timestamp     int64  //区块创建的时间
	Data          []byte //区块存储的实际的有效信息
	PrevBlockHash []byte // 前一个区块的哈希 父哈希
	Hash          []byte //当前区块的哈希
	Nonce          int
}


//创建一个新的区块
func NewBlock(data string , prevBlockHash []byte) *Block{
	block :=&Block{time.Now().Unix(),
		[]byte(data),prevBlockHash,[]byte{},0}
	pow :=NewproofOfWork(block)
	nonce,hash :=pow.Run()
	block.Hash = hash[:]
	block.Nonce = nonce
    return block
}



//设置hash
func(b *Block)  SetHash(){
	timestamp:=[]byte(strconv.FormatInt(b.Timestamp,10))
	headers:= bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]

}




//创世区块的创建
func NewGensisBlock() *Block{
	return NewBlock("Gensis Block",[]byte{})
}


