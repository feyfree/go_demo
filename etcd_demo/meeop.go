package etcd_demo

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"os"
	"time"
)

// 实现文件 两端同步
// directory A -> directory B 使用 etcd 存储元信息, A 文件更新操作，会同步更新到 B 文件中

var (
	FileA = ""
	FileB = ""
)

type FileEtcdInfo struct {
	CheckSum   string `json:"check_sum"`
	Path       string `json:"path"`
	CreateTime string `json:"create_time"`
}

func addFile(filepath string, key string) {
	// 创建一个etcd客户端
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"}, // etcd服务地址
		DialTimeout: 5 * time.Second,            // 连接超时时间
	})
	defer cli.Close()
	// 读取文件
	file, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Println("读取文件异常")
		panic(err)
	}
	fmt.Println(file)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	marshal, _ := json.Marshal(&FileEtcdInfo{CheckSum: digest(filepath), Path: filepath, CreateTime: time.Now().Format(time.RFC3339)})
	_, err = cli.Put(ctx, key, string(marshal))
	if err != nil {
		fmt.Println("更新数据异常")
		panic(err)
	}

}

func listen(key string) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379"}, // etcd服务地址
		DialTimeout: 5 * time.Second,            // 连接超时时间
	})
	defer cli.Close()
	if err != nil {
		fmt.Println("链接etcd异常")
		panic(err)
	}
	// 创建一个watcher
	watcher := cli.Watch(context.Background(), key)

	// 循环监听事件
	for {
		select {
		case resp := <-watcher:
			for _, ev := range resp.Events {
				fmt.Printf("Type: %v, Key: %s, Value: %s\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
				// 读取源文件的内容
				v := new(FileEtcdInfo)
				err := json.Unmarshal(ev.Kv.Value, v)
				if err != nil {
					fmt.Println("解析错误")
					panic(err)
				}
				filepath := v.Path
				fileData, _ := os.ReadFile(filepath)
				fmt.Println(fileData)
			}
		case <-time.After(time.Second * 1000):
			fmt.Println("timeout")
			return
		}
	}
}

func digest(filepath string) string {
	file, _ := os.ReadFile(filepath)
	hash := sha256.Sum256(file)
	return fmt.Sprintf("%x", hash)
}
