package BLC

import (
	"github.com/boltdb/bolt"
	"log"
)

//区块链迭代器管理文件

//实现迭代器基本结构
type BlockChainIterator struct {
	Db          *bolt.DB //迭代目标
	CurrentHash []byte   //当前的迭代目标的哈希
}

//next()函数
//创建迭代器对象
//实现迭代函数next 获取到每一个区块
func (blc *BlockChain) Iterator() *BlockChainIterator {
	return &BlockChainIterator{blc.Db, blc.Tip}
}

//实现迭代函数next 读取到每一个区块
func (bcit  *BlockChainIterator) Next() *Block {
	var block *Block
	err := bcit.Db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))

		currentBlockBytes := b.Get(bcit.CurrentHash)
		block = DeserializeBlock(currentBlockBytes)

		return nil
	})

	if nil != err {
		log.Panicf("iterator the db failed!%v\n", err)
	}
	//更新迭代器中的哈希值
	bcit.CurrentHash = block.PrevBlockHash
	return block
}
