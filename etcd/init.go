package etcd

import (
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"time"
)

var (
	client *clientv3.Client
	dbKvHandler clientv3.KV
)


func init() {
	//certF *os.File
	// 客户端配置
	// rootCertPool := x509.NewCertPool()
	// rootCertPool.AppendCertsFromPEM([]byte(certFile))
	//Trace.Printf("DB: %v %d", dbas, len(dbas))
	config := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
		// TLS: &tls.Config{
		// 	RootCAs:      rootCertPool,
		// 	Certificates: []tls.Certificate{etcdCert},
	}

	c, err := clientv3.New(config)
    client = c
	if err != nil {
		panic(err)
	}
	fmt.Println(client)
	fmt.Println("connect success!")

	dbKvHandler = clientv3.NewKV(client)
	fmt.Println(dbKvHandler)
}
