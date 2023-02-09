package go20230209

import (
	"fmt"
	"testing"
	"time"
)

func TestChannelA(t *testing.T) {
	c := make(chan int) // 一个非缓冲通道
	go func(ch chan<- int, x int) {
		time.Sleep(time.Second)
		// <-ch    // 此操作编译不通过
		ch <- x * x // 阻塞在此，直到发送的值被接收
	}(c, 3)
	done := make(chan struct{})
	go func(ch <-chan int) {
		n := <-ch      // 阻塞在此，直到有值发送到c
		fmt.Println(n) // 9
		// ch <- 123   // 此操作编译不通过
		time.Sleep(time.Second)
		done <- struct{}{}
	}(c)
	<-done // 阻塞在此，直到有值发送到done
	fmt.Println("bye")
}

func TestBufferChannel(t *testing.T) {
	c := make(chan int, 2) // 一个容量为2的缓冲通道
	c <- 3
	c <- 5
	close(c)
	fmt.Println(len(c), cap(c)) // 2 2
	x, ok := <-c
	fmt.Println(x, ok)          // 3 true
	fmt.Println(len(c), cap(c)) // 1 2
	x, ok = <-c
	fmt.Println(x, ok)          // 5 true
	fmt.Println(len(c), cap(c)) // 0 2
	x, ok = <-c
	fmt.Println(x, ok) // 0 false
	x, ok = <-c
	fmt.Println(x, ok)          // 0 false
	fmt.Println(len(c), cap(c)) // 0 2
	close(c)                    // 此行将产生一个恐慌
	c <- 7                      // 如果上一行不存在，此行也将产生一个恐慌。
}

func TestPlayFootball(t *testing.T) {
	var ball = make(chan string)
	kickBall := func(playerName string) {
		for {
			fmt.Print(<-ball, "传球", "\n")
			time.Sleep(time.Second)
			ball <- playerName
		}
	}
	go kickBall("张三")
	go kickBall("李四")
	go kickBall("王二麻子")
	go kickBall("刘大")
	ball <- "裁判"    // 开球
	var c chan bool // 一个零值nil通道
	<-c             // 永久阻塞在此
}

func TestFibonacci(t *testing.T) {
	Fibonacci()
}

func TestSelectA(t *testing.T) {
	var c chan struct{} // nil
	select {
	case <-c: // 阻塞操作
	case c <- struct{}{}: // 阻塞操作
	default:
		fmt.Println("Go here.")
	}
}

func TestSelectB(t *testing.T) {
	c := make(chan string, 2)
	trySend := func(v string) {
		select {
		case c <- v:
		default: // 如果c的缓冲已满，则执行默认分支。
		}
	}
	tryReceive := func() string {
		select {
		case v := <-c:
			return v
		default:
			return "-" // 如果c的缓冲为空，则执行默认分支。
		}
	}
	trySend("Hello!") // 发送成功
	trySend("Hi!")    // 发送成功
	trySend("Bye!")   // 发送失败，但不会阻塞。
	// 下面这两行将接收成功。
	fmt.Println(tryReceive()) // Hello!
	fmt.Println(tryReceive()) // Hi!
	// 下面这行将接收失败。
	fmt.Println(tryReceive()) // -
}

func TestPotentialPanic(t *testing.T) {
	c := make(chan struct{})
	close(c)
	select {
	case c <- struct{}{}: // 若此分支被选中，则产生一个恐慌
	case <-c:
	}
}
