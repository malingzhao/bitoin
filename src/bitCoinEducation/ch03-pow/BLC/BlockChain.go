package BLC

//区块链的管理文件
//区块链的基本结构
type BlockChain struct {
	Blocks []*Block
}

//初始化区块链
func CreateBlockChainWithGensisBlock() *BlockChain {
	//生成创世区块
	block := CreateGensisBlock([]byte("init blockchain"))
	return &BlockChain{[]*Block{block}}
}


//添加区块到区块链
func (bc *BlockChain) AddBlock(height int64, prevBlockHash []byte, data []byte) {
	//var newBlock *Block
	newBlock := NewBlock(height, prevBlockHash, data)
	bc.Blocks = append(bc.Blocks, newBlock)
}

//
