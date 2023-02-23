package go20230217

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
)

// 实现互斥
func TestMutex(t *testing.T) {
	mutex := make(chan struct{}, 1) // 容量必须为1

	counter := 0
	increase := func() {
		mutex <- struct{}{} // 加锁
		counter++
		<-mutex // 解锁
	}

	increase1000 := func(done chan<- struct{}) {
		for i := 0; i < 1000; i++ {
			increase()
		}
		done <- struct{}{}
	}

	done := make(chan struct{})
	go increase1000(done)
	go increase1000(done)
	<-done
	<-done
	fmt.Println(counter) // 2000
}

// 实现互斥
func TestMutex2(t *testing.T) {
	mutex := make(chan struct{}, 1)
	sum := 0
	counter := 10
	done := make(chan struct{}, 1)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			mutex <- struct{}{}
			sum += i
			counter--
			if counter == 0 {
				done <- struct{}{}
			}
			<-mutex
		}()
	}
	<-done
	fmt.Println(sum)
}

func TestSync(t *testing.T) {
	flag := make(chan struct{}, 1)
	done := make(chan struct{}, 2)
	go func() {
		<-flag
		fmt.Println("后")
		done <- struct{}{}
	}()
	go func() {
		flag <- struct{}{}
		fmt.Println("前")
		done <- struct{}{}
	}()
	<-done
	<-done
}

type Seat int
type Bar chan Seat

func (bar Bar) ServeCustomer(c int) {
	log.Print("顾客#", c, "进入酒吧")
	seat := <-bar // 需要一个位子来喝酒
	log.Print("++ customer#", c, " drinks at seat#", seat)
	log.Print("++ 顾客#", c, "在第", seat, "个座位开始饮酒")
	time.Sleep(time.Second * time.Duration(2+rand.Intn(6)))
	log.Print("-- 顾客#", c, "离开了第", seat, "个座位")
	bar <- seat // 释放座位，离开酒吧
}
func TestCountingSemaphore(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	bar24x7 := make(Bar, 10) // 此酒吧有10个座位
	// 摆放10个座位。
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		bar24x7 <- Seat(seatId) // 均不会阻塞
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		go bar24x7.ServeCustomer(customerId)
	}
	for {
		time.Sleep(time.Second)
	} // 睡眠不属于阻塞状态
}
