package go20230217

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func longTimeRequest() <-chan int32 {
	r := make(chan int32)

	go func() {
		time.Sleep(time.Second * 3) // 模拟一个工作负载
		r <- rand.Int31n(100)
	}()

	return r
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

// 返回单向接收通道做为函数返回结果
func TestChannel01(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	a, b := longTimeRequest(), longTimeRequest()
	fmt.Println(sumSquares(<-a, <-b))
}

func source(c chan<- int32) {
	ra, rb := rand.Int31(), rand.Intn(3)+1
	// 睡眠1秒/2秒/3秒
	time.Sleep(time.Duration(rb) * time.Second)
	c <- ra
}

// 采用最快回应
func TestEarliestOne(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	startTime := time.Now()
	c := make(chan int32, 5) // 必须用一个缓冲通道
	for i := 0; i < cap(c); i++ {
		go source(c)
	}
	rnd := <-c // 只有第一个回应被使用了
	fmt.Println(time.Since(startTime))
	fmt.Println(rnd)
}
