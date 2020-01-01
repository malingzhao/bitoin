package main

import (
	"Simplebitcoin/ch04-client/BLC"
)

func main() {

	bc := BLC.NewBlockChain()
	defer bc.Db.Close()

	cli :=BLC.CLI{bc}
	cli.Run()
}
