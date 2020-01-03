package BLC
//交易书瑞皇冠梨
//输入结构

type TXInput struct{
  //交易的哈希
	TxHash    []byte
	//引用的上一交易的输出的索引号
	Vout int

	//用户名
	ScriptSig string
}