package go20230209

import (
	"fmt"
	"testing"
)

func TestFunction(t *testing.T) {
	// 这几行编译没问题。
	AddSub(HalfAndNegative(6))
	AddSub(AddSub(AddSub(7, 5)))
	AddSub(AddSub(HalfAndNegative(6)))
	Dummy(HalfAndNegative(6))
	_, _ = AddSub(7, 5)

	// 下面这几行编译不通过。
	/*
		_, _, _ = 6, AddSub(7, 5)
		Dummy(AddSub(7, 5), 9)
		Dummy(AddSub(7, 5), HalfAndNegative(6))
	*/
}

func TestFunctionAsParams(t *testing.T) {
	fmt.Printf("%T\n", Double) // func(int) int
	// Double = nil // error: Double是不可修改的

	var f func(n int) int // 默认值为nil
	f = Double
	g := Apply
	fmt.Printf("%T\n", g) // func(int, func(int) int) int

	fmt.Println(f(9))         // 18
	fmt.Println(g(6, Double)) // 12
	fmt.Println(Apply(6, f))  // 12
}
