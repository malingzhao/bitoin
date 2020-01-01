package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
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


//将block类型序列化成字节
func (b *Block) Serialize() []byte {
   //定义一个buffer存储序列化之后的数据
	var result bytes.Buffer
	//初始化一个encoder
	encoder := gob.NewEncoder(&result)
    //对block进行编码
	err :=encoder.Encode(b)
    if err!=nil{
    	log.Panic(err)
	}
    return result.Bytes()
}


//将字节类型反序列化成Block类型
func DeserializeBlock(d []byte) *Block {
	//定义block存储反序列化之后的数据
	var block Block
	//通过构建decoder对象 并将字节传递进去
	decoder :=gob.NewDecoder(bytes.NewReader(d))
	//对之进行解码的操作
	err := decoder.Decode(&block)
	if err!=nil {
		log.Panic(err)
	}
	return  &block
}
