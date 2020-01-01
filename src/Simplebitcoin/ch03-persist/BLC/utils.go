package BLC

import (
	"bytes"
	"encoding/binary"
	"log"
)

//将int类型转换成16进制
func IntToHex(num int64) []byte{
	buff := new(bytes.Buffer)
	err := binary.Write(buff,binary.BigEndian,num)
	if err!=nil {
		log.Panic(err)
	}
	return buff.Bytes()
}