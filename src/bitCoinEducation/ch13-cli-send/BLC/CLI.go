package BLC

import (
	"flag"
	"fmt"
	"github.com/boltdb/bolt"
	"log"
	"os"
)

//对blockchain命令行进行操作

//client对象
type CLI struct {
}

//用法展示
func PrintUsage() {

	fmt.Println("Usage:")
	//初始化区块链
	fmt.Println("\tcreateblockchain  -address address  --创建区块\n")
	//添加区块
	fmt.Println("\taddblock --添加区块\n")
	//打印完整的区块信息
	fmt.Println("\tprintchain --输出区块链信息\n")
	//通过命令行转账
	fmt.Printf("\t-from FROM -to TO -amount AMOUNT --发起转账\n")
	fmt.Println("\t转账参数说明")
	fmt.Printf("\t\t-from FROM   转账原地址\n")
	fmt.Printf("\t\t-to TO 转账目标地址 \n")
	fmt.Println("\t\t-amount  AMOUNT--转账金额")

}

//初始化
func (cli *CLI) createBlockChain(address string) {
	CreateBlockChainWithGensisBlock(address)
}

//添加区块
func (cli *CLI) addBlock(txs []*Transaction) {

	blockchain := BlockChainObject()
	//获取到blockchain的对象实例
	blockchain.AddBlock(txs)

}

//打印完整区块信息
func (cli *CLI) printchain() {
	blockchain := BlockChainObject()
	blockchain.PrintChain()
}

//参数数量检测函数
func IsValidArgs() {
	if len(os.Args) < 2 {
		PrintUsage()
		//直接退出
		os.Exit(1)
	}
}

//运行命令行
func (cli *CLI) Run() {
	//检测参数数量
	IsValidArgs()
	//新建相关的命令

	//添加区块
	addBlockCmd := flag.NewFlagSet("addblock", flag.ExitOnError)

	//输出区块链的额完整信息
	printChainCmd := flag.NewFlagSet("printchain", flag.ExitOnError)
	//创建区块链
	CreateBLCWithGensisBlockCmd := flag.NewFlagSet("createblockchain", flag.ExitOnError)
	sendCmd := flag.NewFlagSet("send",flag.ExitOnError)
	//数据参数处理
	flagAddBlockArg := addBlockCmd.String("data", "send 100 btc to player", "添加区块数据")
	flagCreateBlockChainArg := CreateBLCWithGensisBlockCmd.String("address", "troytan", "指定接收系统奖励的旷工地址")
	//发起交易参数
	flagSendFromArg :=sendCmd.String("from","","转账源地址")
	//判断命令
	flagSendToArg :=sendCmd.String("to","","目标地址")
	flagSendAmoutnArg :=sendCmd.String("from","","转账金额")

	switch os.Args[1] {
	case "addblock":
		if err := addBlockCmd.Parse(os.Args[2:]); nil != err {
			log.Panicf("parse addBlockCmd failed! %v\n", err)
		}
	case "printchain":
		if err := printChainCmd.Parse(os.Args[2:]); nil != err {
			log.Panicf("parse printChainCmd failed! %v\n", err)
		}
	case "createblockchain":
		if err := CreateBLCWithGensisBlockCmd.Parse(os.Args[2:]); nil != err {
			log.Panicf("parse CreateBLCWithGensisBlockCmd failed! %v\n", err)
		}
	case "send":

		if err:=sendCmd.Parse(os.Args[2:]); nil!=err{
			log.Panicf("parse sendCmd failed! %v\n", err)
		}

	default:
		//没有传递任何命令或者传递的命令不在上面的命令中
		PrintUsage()
		os.Exit(1)
	}
	//发起转账
	if sendCmd.Parsed(){
		if *flagSendFromArg == "" {
			fmt.Println("原地址不能为空")
			PrintUsage()
			os.Exit(1)
		}
		if *flagSendToArg == "" {
			fmt.Println("目标地址不能为空")
			PrintUsage()
			os.Exit(1)
		}
		if *flagSendAmoutnArg == "" {
			fmt.Println("转账金额不能为空")
			PrintUsage()
			os.Exit(1)
		}
		fmt.Printf("\tFROM:[%s]\n",*flagSendFromArg)
		fmt.Printf("\tTO:[%s]\n",*flagSendFromArg)
		fmt.Printf("\tAMOUNT:[%s]\n",*flagSendAmoutnArg)
	}
	//添加区块命令
	if addBlockCmd.Parsed() {
		if *flagAddBlockArg == "" {
			PrintUsage()
			os.Exit(1)
		}
		//调用
		cli.addBlock([]*Transaction{})
	}
	if printChainCmd.Parsed() {
		cli.printchain()
	}

	//输出区块链信息
	if CreateBLCWithGensisBlockCmd.Parsed() {
		if *flagCreateBlockChainArg == " " {
			PrintUsage()
			os.Exit(1)
		}
		cli.createBlockChain(*flagCreateBlockChainArg)
	}
}

//获取一个blockchain对象
func BlockChainObject() *BlockChain {
	//获取命令
	//获取Tip
	db, err := bolt.Open(dbName, 0600, nil)
	if nil != err {
		log.Panicf("open the db[%s] failed %v", dbName, err)
	}
	//获取tip
	var tip []byte

	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(blockTableName))
		if nil != b {
			tip = b.Get([]byte("l"))
		}
		return nil

	})
	if nil != err {
		log.Panicf("get the blockchain  object failed! %v\n ", err)

	}
	return &BlockChain{Db: db, Tip: tip}
}

//发起交易的
func(cli *CLI) send(){

}