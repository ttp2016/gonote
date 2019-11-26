package main

import (
	"github.com/coreos/etcd/clientv3"
	"time"
	"fmt"
)

func main(){
	var (
		config clientv3.Config
		err error
		client *clientv3.Client
	)
	//配置
	config = clientv3.Config{
		Endpoints:[]string{"192.168.1.188:2379"},
		DialTimeout:time.Second*5,
	}
	//连接
	if client,err = clientv3.New(config);err != nil{
		fmt.Println(err)
		return
	}
	client=client
}