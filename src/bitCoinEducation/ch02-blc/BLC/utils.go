package BLC

import (
	"bytes"
	"encoding/binary"
	"log"
)

// 实现int64转换成byte字节数组
func IntToHex(data int64) []byte{

	buffer :=new(bytes.Buffer )
	err :=binary.Write(buffer,binary.BigEndian,data)
	if nil !=err{
		log.Panicf("int transact to []byte failedf!%v\n",err)
	}
	return buffer.Bytes()
}
