package main

import (
	"fmt"
	"github.com/boltdb/bolt"
	"log"
)

func main() {

	db, err := bolt.Open("my.db", 0600, nil)

	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	db.Update(func(tx *bolt.Tx) error {
		//创建一个桶
		b, err := tx.CreateBucket([]byte("MyBucket"))
		if err != nil {
			return fmt.Errorf("create bucket :%s", err)
		}

		//写入数据
		if nil != b {
			err := b.Put([]byte("l"), []byte("11"))
			if nil != err {
				return  err
			}
		}
		return nil
	})

//read
db.View(func(tx *bolt.Tx) error {

	//获取通
	b :=tx.Bucket([]byte("MyBucket"))
	if nil!=b {
		value :=b.Get([]byte("l"))
		fmt.Printf("value:%s\n",string(value))
	}
	return  nil
})


}
