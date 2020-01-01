package BLC

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

//定义挖矿的最大难度值
var (
maxNonce = math.MaxInt64
	)

const targetBits = 16


//工作量证明的结构体
type ProofOfWork struct {
	 Block *Block    //区块
	 target *big.Int //挖矿难度
}

//新建工作量证明
func NewproofOfWork(b *Block)  *ProofOfWork {

target := big.NewInt(1)
target.Lsh(target,uint(256-targetBits))
pow :=&ProofOfWork{b,target}
return  pow
}



//展示工作量证明 显示核心
func  (pow *ProofOfWork) Run() (int, []byte){
	var hashInt big.Int
	var hash [32]byte
	nonce :=0
	//挖矿
	fmt.Printf("Mining the block containing %s\n", pow.Block.Data)

  for nonce<maxNonce {
  	//准备数据
  	data :=pow.prepareData(nonce)
  	hash =sha256.Sum256(data)
  	fmt.Printf("\r %x",hash)
  	hashInt.SetBytes(hash[:])
  	if hashInt.Cmp(pow.target)==-1 {
  		break;
	}else{
		nonce ++ ;
	}
  }
	fmt.Println("\n\n")
	return  nonce,hash[:]
}




//准备数据
//join方法 将数据全部封装成bytes的类型
func (pow *ProofOfWork) prepareData(nonce int) []byte{
	data :=bytes.Join([][]byte{
		pow.Block.PrevBlockHash,
		pow.Block.Data,
		IntToHex(pow.Block.Timestamp),
		IntToHex(int64(targetBits)),
		IntToHex(int64(nonce)),
	},[]byte{})
		return data
}



//验证区块的工作量证明
func (pow *ProofOfWork) Validate() bool {
	var hashInt big.Int

	data :=pow.prepareData(pow.Block.Nonce)
	hash :=sha256.Sum256(data)
	hashInt.SetBytes(hash[:])
	isValid :=hashInt.Cmp(pow.target)==-1
	return isValid
}


