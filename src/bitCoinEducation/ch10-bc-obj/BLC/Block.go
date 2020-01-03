package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
	"time"
)

//区块的基本结构以及功能管理文件
//实现一个最基本的区块结构
type Block struct {
	TimeStamp     int64  //区块的时间戳 代表区块时间
	Hash          []byte //当前区块哈希
	PrevBlockHash []byte //前区块哈希
	Height        int64  //区块的高度
	Data          []byte //交易数据
	Nonce         int64  //在运行pow时生成的哈希变化值，也代表pow运行时动态修改的数据
}

//新建区块
func NewBlock(height int64, prevBlockHash, data []byte) *Block {
	var block Block
	block = Block{
		TimeStamp:     time.Now().Unix(),
		Hash:          nil,
		PrevBlockHash: prevBlockHash,
		Height:        height,
		Data:          data,
	}
	//生成哈希

	block.SetHash()
	//.....

	pow := NewProofOfWork(&block)
	//执行工作量证明算法
	hash,nonce:=pow.Run()
	block.Hash=hash
	block.Nonce = int64(nonce)
	return &block
}

//计算区块哈希  为了不传参
func (b *Block) SetHash() {

	//设置hash根据什么设置 时间戳.......
	//调用sha256实现hash生成
	timeStampBytes := IntToHex(b.TimeStamp)
	heightBytes := IntToHex(b.Height)
	blockBytes := bytes.Join([][]byte{
		heightBytes,
		timeStampBytes,
		b.PrevBlockHash,
		b.Data,
	}, []byte{})
	hash := sha256.Sum256(blockBytes)
	b.Hash = hash[:]
}

//生成创世区块
func CreateGensisBlock(data []byte) *Block {
	return NewBlock(1, nil, data)
}

//区块结构序列化
func (block *Block) Serialize() []byte{

var buffer bytes.Buffer
//新建编码对象
encoder:=gob.NewEncoder(&buffer)
//编码（序列化)
if err :=encoder.Encode(block) ; nil!=err {
	log.Panicf("serialize the block to []byte failed%v\n",err)
}
	return  buffer.Bytes()
}

//区块结构反序列化
func DeserializeBlock(blockBytes []byte) *Block{
	var block Block
	//fmt.Println("blockBytes:%v\n",blockBytes)
	//新建decoder对象
	decoder :=gob.NewDecoder(bytes.NewReader(blockBytes))
	if err:=decoder.Decode(&block); nil !=err{
		log.Panicf("deserialize the []byte to the block failed",err)
	}
	return  &block
}
