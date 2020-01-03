package BLC

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
)

// 实现int64转换成byte字节数组
func IntToHex(data int64) []byte {

	buffer := new(bytes.Buffer)
	err := binary.Write(buffer, binary.BigEndian, data)
	if nil != err {
		log.Panicf("int transact to []byte failedf!%v\n", err)
	}
	return buffer.Bytes()
}

//标准json格式转切片
//windows下的特殊的格式

//bc.exe send -from "["\troytan\"]" -to "["\Alice\"]" -value "["\10\"]"
func JSONToSlice(jsonString string) []string {

	var strSlice []string
	//jjosn
	if err := json.Unmarshal([]byte(jsonString), &strSlice); nil != err {
		log.Panicf("json to []string failed !%v\n", err)
	}
	return strSlice
}
