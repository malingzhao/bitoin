package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
)

//对blockchain命令行进行操作

//client对象
type CLI struct {
	BC *BlockChain //BlockChain对象
}

//用法展示
func PrintUsage() {

	fmt.Println("Usage:")
	//初始化区块链
	fmt.Println("\tcreateblockchain  -data DATA --创建区块\n")
	//添加区块
	fmt.Println("\taddblock --添加区块\n")
	//打印完整的区块信息
	fmt.Println("\tprintchain --输出区块链信息\n")
}

//初始化
func (cli *CLI) createBlockChain() {
	CreateBlockChainWithGensisBlock()
}

//添加区块
func (cli *CLI) addBlock(data string) {
	cli.BC.AddBlock([]byte(data))

}

//打印完整区块信息
func (cli *CLI) printchain() {
	cli.BC.PrintChain()
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
	//数据参数处理
	flagAddBlockArg := addBlockCmd.String("data", "send 100 btc to player", "添加区块数据")

	//判断命令
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
	default:
		//没有传递任何命令或者传递的命令不在上面的命令中
		PrintUsage()
		os.Exit(1)
	}
	//添加区块命令
	if addBlockCmd.Parsed() {
		if *flagAddBlockArg == "" {
			PrintUsage()
			os.Exit(1)
		}

		cli.addBlock(*flagAddBlockArg)
	}
	if printChainCmd.Parsed() {
		cli.printchain()
	}

	//输出区块链信息
	if CreateBLCWithGensisBlockCmd.Parsed() {
		cli.createBlockChain()
	}
}
