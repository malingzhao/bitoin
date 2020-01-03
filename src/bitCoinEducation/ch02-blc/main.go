package main

import (
	"bitCoinEducation/ch02-blc/BLC"
	"fmt"
)

//启动
func main() {

	//block := BLC.NewBlock(1, nil, []byte("the first  block testing"))
	//fmt.Printf("the first block:%v\n", block)

	bc := BLC.CreateBlockChainWithGensisBlock()
	fmt.Printf("blockchain :%v\n", bc.Blocks[0])
	//上链
	bc.AddBlock(bc.Blocks[len(bc.Blocks)-1].Height+1,
		bc.Blocks[len(bc.Blocks)-1].Hash, []byte("a send 10 eth to b "))

	bc.AddBlock(bc.Blocks[len(bc.Blocks)-1].Height+1,
		bc.Blocks[len(bc.Blocks)-1].Hash, []byte("c send 10 eth to d"))

	for _, block := range bc.Blocks {
		fmt.Printf("prevBlockHash:%x\t currentHash:%x\n",block.PrevBlockHash , block.Hash)
	}
}
