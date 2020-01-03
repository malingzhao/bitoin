package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
)

//区块链的管理文件
//数据库名称
const dbName = "block.db"

//表名称
const blockTableName = "blocks"

//区块链的基本结构
type BlockChain struct {
	//Blocks []*Block //区块的切片
	Db  *bolt.DB //数据库对象
	Tip []byte   //保存最新区块的哈希值
}

//初始化区块链
func CreateBlockChainWithGensisBlock() *BlockChain {
	//保存最新区块的哈希值
	var blockHash []byte

	//生成创世区块
	//block := CreateGensisBlock([]byte("init blockchain"))
	//1. 创建或者打开一个数据库
	db, err := bolt.Open(dbName, 0600, nil)
	if err != nil {
		log.Panicf("create db [%s] failed %v\n", dbName, err)
	}
	//2. 创建桶 把生成的创世块放到数据库中
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if b == nil {
			//没找到
			b, err := tx.CreateBucket([]byte(blockTableName))
			if nil != err {
				log.Panicf("create bucket [%s] failed %v\n", blockTableName, err)
			}
			//生成创区块
			gensisBlock := CreateGensisBlock([]byte("init blockchain "))
			//存储
			//1 key value 分别以什么数据代表 --hash
			//2 如何把block结构存储到数据库里面去

			err = b.Put(gensisBlock.Hash, gensisBlock.Serialize())
			if nil != err {
				log.Panicf("insert the gensis block failed %v\n", err)
			}
			blockHash = gensisBlock.Hash
			//保存最新区块的哈希值
			// l:lastest
			err = b.Put([]byte("l"), gensisBlock.Hash)
			if nil != err {
				log.Panicf("save the hash of gensis block failed %v\n", err)
			}
		}
		return nil
	})
	//3. 把创世区块传到数据库中

	return &BlockChain{Db: db, Tip: blockHash}
}

//添加区块到区块链
func (bc *BlockChain) AddBlock(data []byte) {

	//更新区块数据(insert)
	err := bc.Db.Update(func(tx *bolt.Tx) error {
		//1. 获取数据库桶
		b := tx.Bucket([]byte(blockTableName))
		if nil != b {
			//2. 获取最后插入的区块
			blockBytes := b.Get(bc.Tip)
			//fmt.Printf("Add--%v", blockBytes)

			//区块数据反序列化
			lastest_block := DeserializeBlock(blockBytes)
			//3.新建区块
			//
			newBlock := NewBlock(lastest_block.Height+1, lastest_block.Hash, data)
			//4 存入数据库
			err := b.Put(newBlock.Hash, newBlock.Serialize())
			if nil != err {
				log.Panicf("insert the new block to db failed %v", err)
			}
			err = b.Put([]byte("l"), newBlock.Hash)
			if nil != err {
				log.Panicf("update the lastest block hash to db  failed %v", err)
			}
			//更新区块链对象中的最新区块哈希
			bc.Tip = newBlock.Hash
		}
		return nil
	})
	if err != nil {
		log.Panicf("insert tblock to db failed %v", err)
	}
}

//遍历数据库，输出所有区块信息
func (bc *BlockChain) PrintChain() {

	var curBlock *Block
	var currentHash []byte = bc.Tip
	//读取数据库
	fmt.Println("\n打印区块链完整信息")
	//循环读取
	//什么时候退出
	for {
		fmt.Println("--------------------")
		bc.Db.View(func(tx *bolt.Tx) error {
			b := tx.Bucket([]byte(blockTableName))
			if nil != b {
				blockBytes := b.Get(currentHash)
				curBlock := DeserializeBlock(blockBytes)
				//输出区块信息
				fmt.Printf("\tHash:%x\n", curBlock.Hash)
				fmt.Printf("\tPrevBlockHash:%x\n", curBlock.PrevBlockHash)
				fmt.Printf("\tTimeStamp:%s\n", curBlock.TimeStamp)
				fmt.Printf("\tData:%v\n", curBlock.Data)
				fmt.Printf("\tHeight:%d\n", curBlock.Height)
				fmt.Printf("\tNonce:%d\n", curBlock.Nonce)
			}
			return nil
		})
		//退出条件
		var hashInt big.Int
		hashInt.SetBytes(curBlock.PrevBlockHash)
		//比较
		//if big.NewInt(0).Cmp(&hashInt)==0{
		//	//遍历到创世区块
		//	break
		//}
		//if len(curBlock.PrevBlockHash)==0 {
		//	break
		//}
		if curBlock.PrevBlockHash==nil {
			break
		}
		//更新当前要获取的区块哈希值
		currentHash =curBlock.PrevBlockHash
	}

}
