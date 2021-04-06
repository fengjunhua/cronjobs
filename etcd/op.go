package etcd

import (
	"context"
	"log"
	"time"
)

/*
const (
	dbOpPut       opType = 0
	dbOpDelete    opType = 1
	dbOpDeleteAll opType = 2 //with prefix
	dbOpGet       opType = 3
	dbOpGetAll    opType = 4 //with prefix
)

 */

// 获取
func dbGet(key string) (value []byte, modRev int64, err error) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Duration(2)*time.Second)
	resp, err := dbKvHandler.Get(ctx, key)
	cancelFunc()
	if err != nil {
		log.Println(err)
		return []byte(""), 0, err
	}
	return resp.Kvs[0].Value, resp.Kvs[0].ModRevision, nil
}

func Put() {

}

func del()  {

}

func delAll()  {

}

func get()  {

}

func getAll() {

}
func kvPut()  {

}

func kvDel()  {

}