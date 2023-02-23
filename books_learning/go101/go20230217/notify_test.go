package go20230217

import (
	"log"
	"testing"
)
import "time"

type T = struct{}

func worker(id int, ready <-chan T, done chan<- T) {
	<-ready // 阻塞在此，等待通知
	log.Print("Worker#", id, "开始工作")
	// 模拟一个工作负载。
	time.Sleep(time.Second * time.Duration(id+1))
	log.Print("Worker#", id, "工作完成")
	done <- T{} // 通知主协程（N-to-1）
}

func TestNotify(t *testing.T) {
	log.SetFlags(0)

	ready, done := make(chan T), make(chan T)
	go worker(0, ready, done)
	go worker(1, ready, done)
	go worker(2, ready, done)

	// 模拟一个初始化过程
	time.Sleep(time.Second * 3 / 2)
	// 单对多通知
	ready <- T{}
	ready <- T{}
	ready <- T{}
	// 等待被多对单通知
	<-done
	<-done
	<-done
}
