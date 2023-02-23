package some_code

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGc(t *testing.T) {
	// 假设此切片的长度很大，以至于它的元素
	// 将被开辟在堆上。
	bs := make([]byte, 1<<31)

	// 一个聪明的编译器将觉察到bs的底层元素
	// 部分已经不会再被使用，而正确地认为bs的
	// 底层元素部分在此刻可以被安全地回收了。

	fmt.Println(len(bs))
}

func TestKeepLive(t *testing.T) {
	bs := make([]int, 1000000)

	fmt.Println(len(bs))
	runtime.KeepAlive(&bs)
	// 对于这个特定的例子，也可以调用
	// runtime.KeepAlive(bs)。
}
