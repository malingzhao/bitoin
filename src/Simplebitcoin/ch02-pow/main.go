package main

import (
	"Simplebitcoin/ch02-pow/BLC"
	"fmt"

)

func main() {

	bc := BLC.NewBlockChain()

	bc.AddBlock("send 1 BTC to Ivan")
	bc.AddBlock("send 3 more BTC to Ivan")


	for _,block :=range bc.Blocks{
		fmt.Printf("Prev.Hash:%x\n",block.PrevBlockHash)
		fmt.Printf("Data:%s\n",block.Data)
		fmt.Printf("Hash:%x\n",block.Hash)
		fmt.Println()
	}

}
