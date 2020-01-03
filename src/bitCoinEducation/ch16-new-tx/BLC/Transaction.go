package BLC

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"log"
)

//交易管理文件

//定义一个基本的交易结构
type Transaction struct {
	//交易的哈希（表示）
	TxHash []byte
	//输入列表
	Vins []*TxInput
	//输出列表
	Vouts []*TxOutput
}

//实现coinbase交易
func NewCoinBaseTransaction(address string) *Transaction {
	//var txCoinBase *Transaction
	//输入
	//coinbase特点
	//txHashnil
	//vout
	//ScriptSig 系统奖励

	txInput := &TxInput{[]byte{}, -1, "sysytem reward"}
	//输出
	//value
	//address

	txOutput := &TxOutput{10, ""}

	//输入输出组装交易
	txCoinBase := &Transaction{
		nil,
		[]*TxInput{txInput},
		[]*TxOutput{txOutput},
	}

	//交易哈希生成
	txCoinBase.HashTransaction()

	return txCoinBase

}


//生成交易哈希(交易序列化)

func(tx *Transaction) HashTransaction(){
	var result bytes.Buffer
	//设置编码对象
	encoder :=gob.NewEncoder(&result)

	if err:=encoder.Encode(tx); err!=nil{
		log.Panicf("tx Hash encoded failed")
	}
	hash:=sha256.Sum256(result.Bytes())
	tx.TxHash = hash[:]
}

//生成普通转账交易

func NewSimpleTransaction(from string ,to string,amount int) *Transaction{
	var txInputs []*TxInput  //输入列表
	var txOutputs []*TxOutput//输出列表
	//输入
	txInput :=&TxInput{[]byte("1ff7657984e884e11a7440adb8394f78454c1c1275690a7adde3f287edc742e5"),0, from}
	txInputs = append(txInputs,txInput)
	//输出 (转账源)
	txOutput :=&TxOutput{amount,to}
	txOutputs = append(txOutputs,txOutput)

    //输出 找零
if amount <10 {
	txOutput = &TxOutput{10-amount,from}
	txOutputs =append(txOutputs,txOutput)

}

	tx :=Transaction{nil,txInputs,txOutputs}
	tx.HashTransaction()
	return  &tx
}