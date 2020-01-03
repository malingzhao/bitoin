package BLC
//交易的输出管理

//输出结构
type TxOutput struct {
	//金额（大写)
	Value int
	//用户米(UTXO的所有者)
	ScriptPubKey string
}


