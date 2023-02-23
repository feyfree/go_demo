package go20230217

import (
	rand "crypto/rand"
	"fmt"
	"os"
	"sort"
	"testing"
)

// 向一个通道发送一个值来实现单对单通知

func TestNotice(t *testing.T) {
	values := make([]byte, 32*1024*1024)
	if _, err := rand.Read(values); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	done := make(chan struct{}) // 也可以是缓冲的

	// 排序协程
	go func() {
		sort.Slice(values, func(i, j int) bool {
			return values[i] < values[j]
		})
		done <- struct{}{} // 通知排序已完成
	}()

	// 并发地做一些其它事情...

	<-done // 等待通知
	fmt.Println(values[0], values[len(values)-1])
}
