package go20230202

import (
	"fmt"
	"testing"
)

func TestVariables(t *testing.T) {
	//const a uint8 = 256

	// 编译正常 最大的uint
	const MaxUint = ^uint(0)

	// 编译正常 最大的int
	const MaxInt = int(^uint(0) >> 1)

	// 我们可以声明一个常量来表示当前操作系统的位数，或者检查当前操作系统是32位的还是64位的
	const NativeWordBits = 32 << (^uint(0) >> 63) // 64 or 32
	const Is64bitOS = ^uint(0)>>63 != 0
	const Is32bitOS = ^uint(0)>>32 == 0

	/*------------------------------------------------------*/
	// 编译失败 constant -1 overflows uint
	//const MaxUintA = uint(^0)

	// 编译失败 constant -1 overflows uint
	//const MaxUintB uint = ^0

	/*------------------------------------------------------*/

}

func TestIota(t *testing.T) {
	const (
		k = 3 // 在此处，iota == 0

		m float32 = iota + .5 // m float32 = 1 + .5
		n                     // n float32 = 2 + .5

		p    = 9          // 在此处，iota == 3
		q    = iota * 2   // q = 4 * 2
		_                 // _ = 5 * 2
		r                 // r = 6 * 2
		s, a = iota, iota // s, t = 7, 7
		u, v              // u, v = 8, 8
		_, w              // _, w = 9, 9
	)

	const x = iota // x = 0 （iota == 0）
	const (
		y = iota // y = 0 （iota == 0）
		z        // z = 1
	)

	println(m)             // +1.500000e+000
	println(n)             // +2.500000e+000
	println(q, r)          // 8 12
	println(s, a, u, v, w) // 7 7 8 8 9
	println(x, y, z)       // 0 0 1
}

func TestVar(t *testing.T) {
	var a int
	fmt.Println(a == 0)
	var b *int
	fmt.Println(b == nil)
	//fmt.Println(*b) // panic: runtime error: invalid memory address or nil pointer dereference [recovered]
	// new 申请了一段空间， 并返回指针
	c := new(int)
	fmt.Println(c == nil)

}
