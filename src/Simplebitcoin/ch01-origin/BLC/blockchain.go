package BLC

//区块链的结构体
type BlockChain struct {
	Blocks []*Block
}


//向区块链中添加区块
func(bc *BlockChain)  AddBlock(data string ){
	prevBlock :=bc.Blocks[len(bc.Blocks)-1]
   newBlock := NewBlock(data,prevBlock.Hash)
   bc.Blocks = append(bc.Blocks,newBlock)
}




//创建一个新的区块链的方法
func NewBlockChain() *BlockChain{
	return &BlockChain{[]*Block{NewGensisBlock()}}
}