package etcd_demo

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
	"time"
)

func TestEtcd(t *testing.T) {
	// 创建一个etcd客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"}, // etcd服务地址
		DialTimeout: 5 * time.Second,            // 连接超时时间
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cli.Close()

	// 设置一个键值对
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	deadline, ok := ctx.Deadline()
	if ok {
		fmt.Println(deadline)
	}
	_, err = cli.Put(ctx, "foo", "bar")
	cancel()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 获取一个键的值
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	resp, err := cli.Get(ctx, "foo")
	cancel()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, ev := range resp.Kvs {
		fmt.Printf("%s : %s\n", ev.Key, ev.Value)
	}

	// 删除一个键
	ctx, cancel = context.WithTimeout(context.Background(), time.Second)
	_, err = cli.Delete(ctx, "foo")
	cancel()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestLease(t *testing.T) {
	//
	// 创建etcd客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer cli.Close()

	// 创建一个租约
	resp, err := cli.Grant(context.Background(), 10)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 将一个键值对与租约绑定
	_, err = cli.Put(context.Background(), "key", "value", clientv3.WithLease(resp.ID))
	if err != nil {
		fmt.Println(err)
		return
	}

	// 每秒检查一次租约是否还有效
	for {
		time.Sleep(time.Second)
		result, err := cli.Get(context.Background(), "key")
		if err != nil {
			fmt.Println(err)
			return
		}
		if result.Count == 0 {
			fmt.Println("key has been deleted")
			return
		}
	}
}
