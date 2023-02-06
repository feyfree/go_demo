package go20230206

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {
	p0 := new(int)   // p0指向一个int类型的零值
	fmt.Println(p0)  // （打印出一个十六进制形式的地址）
	fmt.Println(*p0) // 0

	x := *p0         // x是p0所引用的值的一个复制。
	p1, p2 := &x, &x // p1和p2中都存储着x的地址。
	// x、*p1和*p2表示着同一个int值。
	fmt.Println(p1 == p2) // true
	fmt.Println(p0 == p1) // false
	p3 := &*p0            // <=> p3 := &(*p0)
	// <=> p3 := p0
	// p3和p0中存储的地址是一样的。
	fmt.Println(p0 == p3) // true
	*p0, *p1 = 123, 789
	fmt.Println(*p2, x, *p3) // 789 789 123

	fmt.Printf("%T, %T \n", *p0, x) // int, int
	fmt.Printf("%T, %T \n", p0, p1) // *int, *int
}

func TestPointerOperation(t *testing.T) {
	type MyInt int64
	type Ta *int64
	type Tb *MyInt

	// 4个不同类型的指针：
	var pa0 Ta
	var pa1 *int64
	var pb0 Tb
	var pb1 *MyInt

	// 下面这6行编译没问题。它们的比较结果都为true。
	_ = pa0 == pa1
	_ = pb0 == pb1
	_ = pa0 == nil
	_ = pa1 == nil
	_ = pb0 == nil
	_ = pb1 == nil

	// 下面这三行编译不通过。
	/*
		_ = pa0 == pb0
		_ = pa1 == pb1
		_ = pa0 == Tb(nil)
	*/
}
