package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

const dbFile = "blockchain.db"
const blockBucket = "blocks"

//区块链的结构体
type BlockChain struct {
	//	Blocks []*Block
	Tip []byte   //存储最后一个区块的哈希
	Db  *bolt.DB //存储了一个数据库的连接
}

//向区块链中添加区块
func (bc *BlockChain) AddBlock(data string) {
	var lastHash []byte

	//在数据库中获取最后一个块的哈希，
	err := bc.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))

		lastHash = b.Get([]byte("l"))
		return nil
	})

	if err != nil {
		log.Panic(err)
	}

	// 然后用它挖出一个新的块的哈希
	newBlock := NewBlock(data, lastHash)

	err = bc.Db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		err := b.Put(newBlock.Hash, newBlock.Serialize())
		if err != nil {
			log.Panic(err)
		}
		err = b.Put([]byte("l"), newBlock.Hash)
		return nil
	})

}

//创建一个新的区块链的方法
func NewBlockChain() *BlockChain {

	/*
		1. 打开一个数据库文文件
		2. 检查文件是否已经存储了一个区块链
		3. 如果已经存储了一个区块链
		   1. 创建一个新的BlockChain实例
		    2. 设置Blockchain实例的tip为数据库存储的最后的一个区块的哈希
		4. 如果没有区块链
		   1. 创建创世快
		   2. 存储到数据库
		   3. 将创世快保存成最后一个快的哈希
		   4. 创建一个新的BlockChain实例
	*/

	var tip []byte
	db, err := bolt.Open(dbFile, 0600, nil)
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))

		if b == nil {
			//如果不存在 生成创世快
			gensis := NewGensisBlock()
			//创建bucket
			b, err := tx.CreateBucket([]byte(blockBucket))
			if err != nil {
				log.Panic(err)
			}
			//将区块保存在里面
			err = b.Put(gensis.Hash, gensis.Serialize())
			//更新";"键存储最后一个快的哈希
			err = b.Put([]byte("l"), gensis.Hash)
		} else {
			//如果存在就读取"l"键
			tip = b.Get([]byte("l"))
		}
		return nil
	})
	if err != nil {
		log.Panic(err)
	}

	bc := BlockChain{Tip:tip,Db: db}

	return &bc

}

//区块链的迭代器
type BlockchainIterator struct {
	CurrentHash []byte
	Db          *bolt.DB
}

//返回链中的下一个区块的方法
func (i *BlockchainIterator) Next() *Block {
	var block *Block
	err := i.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockBucket))
		encodeedBlock := b.Get(i.CurrentHash)
		block = DeserializeBlock(encodeedBlock)
		return nil
	})
	i.CurrentHash = block.PrevBlockHash
	if err!=nil{
		log.Panic(err)
	}
	return block
}
