package BLC

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"math/big"
	"os"
)

//区块链的管理文件
//数据库名称
const dbName = "./block.db"

//表名称
const blockTableName = "blocks"

//区块链的基本结构
type BlockChain struct {
	//Blocks []*Block //区块的切片
	Db  *bolt.DB //数据库对象
	Tip []byte   //保存最新区块的哈希值
}


//判断数据库文件是否存在
func dbExit() bool{

	if _,err:=os.Stat(dbName); os.IsNotExist(err){
		//数据库文件不存在
		return  false
	}
	return  true
}

//初始化区块链
func CreateBlockChainWithGensisBlock() *BlockChain {

	if dbExit(){
		//文件已存在 说明创世区块已经存在
		fmt.Println("文件已存在........")

		os.Exit(1)
	}
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
	fmt.Println("\n打印区块链完整信息")

	var curBlock *Block
	bcit := bc.Iterator() //获取迭代器对象

	//读取数据库

	//循环读取
	//什么时候退出
	for {

		curBlock = bcit.Next()
		fmt.Println("--------------------")
		fmt.Printf("\tHash:%x\n", curBlock.Hash)
		fmt.Printf("\tPrevBlockHash:%x\n", curBlock.PrevBlockHash)
		fmt.Printf("\tTimeStamp:%s\n", curBlock.TimeStamp)
		fmt.Printf("\tData:%s\n", curBlock.Data)
		fmt.Printf("\tHeight:%d\n", curBlock.Height)
		fmt.Printf("\tNonce:%d\n", curBlock.Nonce)

		//退出条件

		var hashInt big.Int
		hashInt.SetBytes(curBlock.PrevBlockHash)

		if big.NewInt(0).Cmp(&hashInt) == 0 {
			//遍历到创世区块
			break
		}


		//if len(curBlock.PrevBlockHash) == 0 {
		//	break
		//}
	}

	//func (bc *BlockChain) PrintChain() {
	//
	//	bci := bc.Iterator()
	//
	//	for {
	//		block := bci.Next()
	//
	//		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
	//		fmt.Printf("Data: %s\n", block.Data)
	//		fmt.Printf("Hash: %x\n", block.Hash)
	//		fmt.Println()
	//
	//		if len(block.PrevBlockHash) == 0 {
	//			break
	//		}
	//	}



	}
