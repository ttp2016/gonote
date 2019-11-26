/*
@Time : 19-11-26 下午1:57 
@Author : xmge
@Desc :

*/
package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2179", "localhost:2379", "localhost:2179"},
		DialTimeout: 5 * time.Second,
	})

	if err != nil {
		log.Fatal(err)
	}

	kv := clientv3.NewKV(cli)

	putResp, err := kv.Put(context.TODO(),"/test/key1", "Hello etcd!")

	fmt.Printf("PutResponse: %v, err: %v\n", putResp, err)

	getResp, err := kv.Get(context.TODO(), "/test/key1")
	fmt.Printf("PutResponse: %v, err: %v\n", getResp, err)
}