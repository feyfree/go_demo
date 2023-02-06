package go20230205

import (
	"fmt"
	"log"
	"math/rand"
	"testing"
	"time"
)

func TestSayGreetings(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	go SayGreetings("hi!", 10)
	go SayGreetings("hello!", 10)
	time.Sleep(2 * time.Second)
}

func TestSayGreetingsWg(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	log.SetFlags(0)
	wg.Add(2) // 注册两个新任务。
	go SayGreetings("hi!", 10)
	go SayGreetings("hello!", 10)
	wg.Wait() // 阻塞在这里，直到所有任务都已完成。
}

//  === RUN   TestWait
// fatal error: all goroutines are asleep - deadlock!
func TestWait(t *testing.T) {
	wg.Add(1)
	go func() {
		time.Sleep(time.Second * 2)
		wg.Wait() // 阻塞在此
	}()
	wg.Wait() // 阻塞在此
}

func TestDefer(t *testing.T) {
	defer fmt.Println("9")
	fmt.Println("0")
	defer fmt.Println("8")
	fmt.Println("1")
	if false {
		defer fmt.Println("not reachable")
	}
	defer func() {
		defer fmt.Println("7")
		fmt.Println("3")
		defer func() {
			fmt.Println("5")
			fmt.Println("6")
		}()
		fmt.Println("4")
	}()
	fmt.Println("2")
	return
	defer fmt.Println("not reachable")
}

func TestTriple(t *testing.T) {
	fmt.Println(Triple(5)) // 15
}

func TestDefer1(t *testing.T) {
	Defer1()
}

func TestDefer2(t *testing.T) {
	Defer2()
}

func TestDefer3(t *testing.T) {
	Defer3()
}

func TestEvaluateValue(t *testing.T) {
	var a = 123
	go func(x int) {
		time.Sleep(time.Second)
		fmt.Println(x, a) // 123 789
	}(a)

	a = 789

	time.Sleep(2 * time.Second)
}

func TestPanicAndRecover(t *testing.T) {
	defer func() {
		fmt.Println("正常退出")
	}()
	fmt.Println("嗨！")
	defer func() {
		v := recover()
		fmt.Println("恐慌被恢复了：", v)
	}()
	panic("拜拜！") // 产生一个恐慌
	fmt.Println("执行不到这里")
}

func TestPanic(t *testing.T) {
	fmt.Println("hi!")

	go func() {
		time.Sleep(time.Second)
		panic(123)
	}()

	for {
		time.Sleep(time.Second)
	}
}
