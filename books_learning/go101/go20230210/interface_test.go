package go20230210

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInterface(t *testing.T) {
	// 一个*Book值被包裹在了一个Aboutable值中。
	var a Aboutable = &Book{"Go语言101"}
	fmt.Println(a) // &{Go语言101}

	fmt.Println(reflect.TypeOf(a))

	// i是一个空接口值。类型*Book实现了任何空接口类型。
	var i interface{} = &Book{"Rust 101"}
	fmt.Println(i) // &{Rust 101}
	fmt.Println(reflect.TypeOf(i))

	// Aboutable实现了空接口类型interface{}。
	i = a
	fmt.Println(i) // &{Go语言101}
	fmt.Println(reflect.TypeOf(i))

}

func TestEmptyInterface(t *testing.T) {
	var i interface{}
	i = []int{1, 2, 3}
	fmt.Println(i) // [1 2 3]
	fmt.Println(reflect.TypeOf(i))

	i = map[string]int{"Go": 2012}
	fmt.Println(i) // map[Go:2012]
	fmt.Println(reflect.TypeOf(i))

	i = true
	fmt.Println(i) // true
	fmt.Println(reflect.TypeOf(i))

	i = 1
	fmt.Println(i) // 1
	fmt.Println(reflect.TypeOf(i))

	i = "abc"
	fmt.Println(i) // abc
	fmt.Println(reflect.TypeOf(i))

	// 将接口值i中包裹的值清除掉。
	i = nil
	fmt.Println(i) // <nil>
	fmt.Println(reflect.TypeOf(i))

}

func TestInterface2(t *testing.T) {
	greetings := []Greeting{
		FriendGreeting{}, BrotherGreeting{},
	}
	for _, greeting := range greetings {
		greeting.SayHello()
	}
}

func TestType01(t *testing.T) {
	// 编译器将把123的类型推断为内置类型int。
	var x interface{} = 123

	// 情形一：
	n, ok := x.(int)
	fmt.Println(n, ok) // 123 true
	n = x.(int)
	fmt.Println(n) // 123

	// 情形二：
	a, ok := x.(float64)
	fmt.Println(a, ok) // 0 false

	// 情形三：
	a = x.(float64) // 将产生一个恐慌
}

type Writer interface {
	Write(buf []byte) (int, error)
}

type DummyWriter struct{}

func (DummyWriter) Write(buf []byte) (int, error) {
	return len(buf), nil
}

func TestWriter(t *testing.T) {
	var x interface{} = DummyWriter{}
	// y的动态类型为内置类型string。
	var y interface{} = "abc"
	var w Writer
	var ok bool

	// DummyWriter既实现了Writer，也实现了interface{}。
	w, ok = x.(Writer)
	fmt.Println(w, ok) // {} true
	x, ok = w.(interface{})
	fmt.Println(x, ok) // {} true

	// y的动态类型为string。string类型并没有实现Writer。
	w, ok = y.(Writer)
	fmt.Println(w, ok) //  false
	w = y.(Writer)     // 将产生一个恐慌
}

func TestSwitchType(t *testing.T) {
	values := []interface{}{
		456, "abc", true, 0.33, int32(789),
		[]int{1, 2, 3}, map[int]bool{}, nil,
	}
	for _, x := range values {
		// 这里，虽然变量v只被声明了一次，但是它在
		// 不同分支中可以表现为多个类型的变量值。
		switch v := x.(type) {
		case []int: // 一个类型字面表示
			// 在此分支中，v的类型为[]int。
			fmt.Println("int slice:", v)
		case string: // 一个类型名
			// 在此分支中，v的类型为string。
			fmt.Println("string:", v)
		case int, float64, int32: // 多个类型名
			// 在此分支中，v的类型为x的类型interface{}。
			fmt.Println("number:", v)
		case nil:
			// 在此分支中，v的类型为x的类型interface{}。
			fmt.Println(v)
		default:
			// 在此分支中，v的类型为x的类型interface{}。
			fmt.Println("others:", v)
		}
		// 注意：在最后的三个分支中，v均为接口值x的一个复制。
	}
}

func TestNilWrapper(t *testing.T) {
	var a, b, c interface{} = "abc", 123, "a" + "b" + "c"
	fmt.Println(a == b) // 第二步的情形。输出"false"。
	fmt.Println(a == c) // 第三步的情形。输出"true"。

	var x *int = nil
	var y *bool = nil
	var ix, iy interface{} = x, y
	var i interface{} = nil
	fmt.Println(ix == iy) // 第二步的情形。输出"false"。
	fmt.Println(ix == i)  // 第一步的情形。输出"false"。
	fmt.Println(iy == i)  // 第一步的情形。输出"false"。

	var s []int = nil // []int为一个不可比较类型。
	i = s
	fmt.Println(i == nil) // 第一步的情形。输出"false"。
	fmt.Println(i == i)   // 第三步的情形。将产生一个恐慌。
}

type I interface {
	m(int) bool
}

type T string

func (t T) m(n int) bool {
	return len(t) > n
}

func TestImplicit(t *testing.T) {
	var i I = T("gopher")
	fmt.Println(i.m(5))    // true
	fmt.Println(I.m(i, 5)) // true
	//fmt.Println(interface{ m(int) bool }.m(i, 5)) // true

	// 下面这几行被执行的时候都将会产生一个恐慌。
	I(nil).m(5)
	I.m(nil, 5)
	interface{ m(int) bool }.m(nil, 5)
}
