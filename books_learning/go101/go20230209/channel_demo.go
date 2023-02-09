package go20230209

import (
	"fmt"
)

func Fibonacci() {
	fibonacci := func() chan uint64 {
		c := make(chan uint64)
		go func() {
			var x, y uint64 = 0, 1
			for ; y < (1 << 63); c <- y { // 步尾语句
				x, y = y, x+y
			}
			close(c)
		}()
		return c
	}
	c := fibonacci()
	//for x, ok := <-c; ok; x, ok = <-c { // 初始化和步尾语句
	//	time.Sleep(time.Second)
	//	fmt.Println(x)
	//}

	// 等价上面的注释
	for x := range c {
		fmt.Println(x)
	}
}
