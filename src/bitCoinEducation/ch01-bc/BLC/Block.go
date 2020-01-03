package BLC

import (
	"bytes"
	"crypto/sha256"
	"time"
)

//区块的基本结构以及功能管理文件
//实现一个最基本的区块结构
type Block struct {
	TimeStamp int64 //区块的时间戳 代表区块时间
	Hash  []byte  //当前区块哈希
	PrevBlockHash []byte  //前区块哈希
	Height int64   //区块的高度
	Data []byte    //交易数据
}

//新建区块
func NewBlock(height int64,prevBlockHash ,data []byte) *Block{
	var block Block
	block = Block{
		TimeStamp:time.Now().Unix(),
		Hash :nil ,
		PrevBlockHash:prevBlockHash,
		Height:height,
		Data:data,
	}
	//生成哈希

	block.SetHash()
	//.....
	return &block
}

//计算区块哈希  为了不传参
func (b *Block)SetHash(){

	//设置hash根据什么设置 时间戳.......
	//调用sha256实现hash生成
	timeStampBytes := IntToHex(b.TimeStamp)
	heightBytes :=IntToHex(b.Height)
	blockBytes:= bytes.Join([][]byte{
		heightBytes,
		timeStampBytes,
		b.PrevBlockHash,
		b.Data,
	}, []byte{})
	hash:=sha256.Sum256(blockBytes)
	b.Hash = hash[:]
}
