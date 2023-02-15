package go20230212

import (
	"fmt"
	"math"
	"reflect"
	"testing"
	"unsafe"
)

func TestUnsafeMethods(t *testing.T) {
	var x struct {
		a int64
		b bool
		c string
	}
	const M, N = unsafe.Sizeof(x.c), unsafe.Sizeof(x)
	fmt.Println(M, N) // 16 32

	fmt.Println(unsafe.Alignof(x.a)) // 8
	fmt.Println(unsafe.Alignof(x.b)) // 1
	fmt.Println(unsafe.Alignof(x.c)) // 8

	fmt.Println(unsafe.Offsetof(x.a)) // 0
	fmt.Println(unsafe.Offsetof(x.b)) // 8
	fmt.Println(unsafe.Offsetof(x.c)) // 16
}

func TestTypeEmbedded(t *testing.T) {
	type T struct {
		c string
	}
	type S struct {
		b bool
	}
	var x struct {
		a int64
		*S
		T
	}

	fmt.Println(unsafe.Offsetof(x.a)) // 0

	fmt.Println(unsafe.Offsetof(x.S)) // 8
	fmt.Println(unsafe.Offsetof(x.T)) // 16

	// 此行可以编译过，因为选择器x.c中的隐含字段T为非指针。
	fmt.Println(unsafe.Offsetof(x.c)) // 16

	// 此行编译不过，因为选择器x.b中的隐含字段S为指针。
	//fmt.Println(unsafe.Offsetof(x.b)) // error

	// 此行可以编译过，但是它将打印出字段b在x.S中的偏移量.
	fmt.Println(unsafe.Offsetof(x.S.b)) // 0
}

func TestUnsafeOperations(tt *testing.T) {
	a := [16]int{3: 3, 9: 9, 11: 11}
	fmt.Println(a)
	eleSize := int(unsafe.Sizeof(a[0]))
	p9 := &a[9]
	up9 := unsafe.Pointer(p9)
	p3 := (*int)(unsafe.Add(up9, -6*eleSize))
	fmt.Println(*p3) // 3
	s := unsafe.Slice(p9, 5)[:3]
	fmt.Println(s)              // [9 0 11]
	fmt.Println(len(s), cap(s)) // 3 5

	t := unsafe.Slice((*int)(nil), 0)
	fmt.Println(t == nil) // true

	// 下面是两个不正确的调用。因为它们
	// 的返回结果引用了未知的内存块。
	_ = unsafe.Add(up9, 7*eleSize)
	_ = unsafe.Slice(p9, 8)
}

func TestFloat64bits(t *testing.T) {
	var f float64 = 1.1
	p := Float64bits(f)
	fmt.Println(reflect.TypeOf(p), p)
	fmt.Println(math.Float64bits(f))
}

func TestFloat64FromBits(t *testing.T) {
	var f uint64 = 10
	s := Float64FromBits(f)
	fmt.Println(s, reflect.TypeOf(s))
	sm := math.Float64frombits(f)
	fmt.Println(sm, reflect.TypeOf(sm))
	fmt.Println(uint64(10))

}

func TestConversion(t *testing.T) {
	type MyString string
	ms := []MyString{"C", "C++", "Go"}
	fmt.Printf("%s \n", ms)
	// ss := ([]string)(ms) 编译错误
	ss := *(*[]string)(unsafe.Pointer(&ms))
	ss[1] = "Java"
	fmt.Printf("%s \n", ms)
	// ms = []MyString(ss) 编译错误
	ms = *(*[]MyString)(unsafe.Pointer(&ss))
	ms[1] = "Javascript"
	fmt.Printf("%s \n", ms)
}

func TestPointer(tt *testing.T) {
	type T struct{ a int }
	var t T
	fmt.Printf("%p\n", &t)                          // 0xc6233120a8
	println(&t)                                     // 0xc6233120a8
	fmt.Printf("%x\n", uintptr(unsafe.Pointer(&t))) // c6233120a8
}

func TestTT(t *testing.T) {
	const N = unsafe.Offsetof(TT{}.y)
	const M = unsafe.Sizeof(TT{}.y[0])
	tt := TT{y: [3]int16{123, 456, 789}}
	p := unsafe.Pointer(&tt)
	// "uintptr(p) + N + M + M"为t.y[2]的内存地址。
	ty2 := (*int16)(unsafe.Pointer(uintptr(p) + N + M + M))
	fmt.Println(*ty2) // 789
}

func TestTT2(t *testing.T) {
	const N = unsafe.Offsetof(TT{}.y)
	const M = unsafe.Sizeof(TT{}.y[0])
	tt := TT{y: [3]int16{123, 456, 789}}
	p := unsafe.Pointer(&tt)
	// ty2 := (*int16)(unsafe.Pointer(uintptr(p)+N+M+M))
	addr := uintptr(p) + N + M + M

	// ...（一些其它操作）

	// 从这里到下一行代码执行之前，t值将不再被任何值
	// 引用，所以垃圾回收器认为它可以被回收了。一旦
	// 它真地被回收了，下面继续使用t.y[2]值的曾经
	// 的地址是非法和危险的！另一个危险的原因是
	// t的地址在执行下一行之前可能改变（见事实三）。
	// 另一个潜在的危险是：如果在此期间发生了一些
	// 操作导致协程堆栈大小改变的情况，则记录在addr
	// 中的地址将失效。
	ty2 := (*int16)(unsafe.Pointer(addr))
	fmt.Println(*ty2)
}
