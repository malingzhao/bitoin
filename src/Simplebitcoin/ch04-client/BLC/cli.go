package BLC

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

type CLI struct {
	Bc *BlockChain
}



//Run 函数
func (cli *CLI) Run(){
	cli.validateArgs()

	//使用标准库里面的flag来解析命令行参数
	addBlockCmd :=flag.NewFlagSet("addblock",flag.ExitOnError)

	printChainCmd:=flag.NewFlagSet("printchain",flag.ExitOnError)
	//给addblock添加‘-data’参数
	addBlockData := addBlockCmd.String("data","","Block Data")

	switch os.Args[1]{
	case "addblock":
		err:=addBlockCmd.Parse(os.Args[2:])
        if err!=nil{
        	log.Panic(err)
		}
	case "printchain":
		err:=printChainCmd.Parse(os.Args[2:])
		if err!=nil{
			log.Panic(err)
		}
	default:
		cli.printUsage( )
		os.Exit(1)
	}
	if addBlockCmd.Parsed(){
		//如果没有接收到任何的参数
		if *addBlockData == ""{
			//打印用法
			addBlockCmd.Usage()
			os.Exit(1)
		}
		cli.addBlock(*addBlockData)
	}
	if printChainCmd.Parsed(){
		cli.printChain()
	}
}



func(cli *CLI) addBlock(data string){
	cli.Bc.AddBlock(data)
	fmt.Println("success")
}


func(cli *CLI) printChain(){
	bci:=cli.Bc.Iterator()
	for {
		block := bci.Next()
		fmt.Printf("Prev.hash：%x\n", block.PrevBlockHash)
		fmt.Printf("Data:%s\n", block.Data)
		fmt.Printf("Hash:%x\n", block.Hash)
		pow := NewproofOfWork(block)
		fmt.Printf("Pow:%s\n", strconv.FormatBool(pow.Validate()))
		fmt.Println()
		if len(block.PrevBlockHash) == 0 {
			break
		}

	}
}

func(cli *CLI) printUsage(){
fmt.Println("Usage:")
fmt.Println("\taddblock -data BLOCK_DATA  - add  a block to the blockchain")
fmt.Println("\tprintchain - print all the blocks of the  blockchain ")
}

func (cli *CLI) validateArgs(){
	if len(os.Args)<2 {
		cli.printUsage()
		os.Exit(1)
	}
}


