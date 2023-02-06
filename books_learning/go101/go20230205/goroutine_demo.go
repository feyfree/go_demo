package go20230205

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func SayGreetings(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d) // 睡眠片刻（随机0到2.5秒）
	}
}

func SayGreetingsWg(greeting string, times int) {
	for i := 0; i < times; i++ {
		log.Println(greeting)
		d := time.Second * time.Duration(rand.Intn(5)) / 2
		time.Sleep(d)
	}
	wg.Done() // 通知当前任务已经完成。
}

func Triple(n int) (r int) {
	defer func() {
		r += n // 修改返回值
	}()

	return n + n // <=> r = n + n; return
}

// Defer1 第一个匿名函数中的循环打印出2、1和0这个序列，
//但是第二个匿名函数中的循环打印出三个3。
//因为第一个循环中的i是在fmt.Println函数调用被推入延迟调用队列的时候估的值，
//而第二个循环中的i是在第二个匿名函数调用的退出阶段估的值（此时循环变量i的值已经变为3）
func Defer1() {
	func() {
		for i := 0; i < 3; i++ {
			defer fmt.Println("a:", i)
		}
	}()
	fmt.Println()
	func() {
		for i := 0; i < 3; i++ {
			defer func() {
				fmt.Println("b:", i)
			}()
		}
	}()
}

func Defer2() {
	for i := 0; i < 3; i++ {
		defer func(i int) {
			// 此i为形参i，非实参循环变量i。
			fmt.Println("b:", i)
		}(i)
	}
}

func Defer3() {
	for i := 0; i < 3; i++ {
		i := i // 在下面的调用中，左i遮挡了右i。
		// <=> var i = i
		defer func() {
			// 此i为上面的左i，非循环变量i。
			fmt.Println("b:", i)
		}()
	}
}
