package numbers

import (
	"fmt"
	"testing"
)

func TestUInt(t *testing.T) {
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)
}

// result -- (注意这个是 int 是有符号位的)
// 127 -128 1
func TestInt8(t *testing.T) {
	var i int8 = 127
	fmt.Println(i, i+1, i*i)
}

func TestBitOperation(t *testing.T) {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("%08b\n", x)    // "00100010", the set {1, 5}
	fmt.Printf("%08b\n", y)    // "00000110", the set {1, 2}
	fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}
	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}
	fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}
}

func TestBitClear(t *testing.T) {
	var x uint8 = 2<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<3
	fmt.Printf("%08b\n", x) // "00100100", the set {2, 5}
	fmt.Printf("%08b\n", y) // "00001010", the set {1, 3}
	// bit clear means： remove the bit both in (x, y) from x, return the value
	fmt.Printf("%08b\n", x&^y) // "00100100", the difference {2, 5}
}

func TestLoopVariable(t *testing.T) {
	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i]) // "bronze", "silver", "gold"
	}
}

func TestShortVariable(t *testing.T) {
	// 10进制默认不为 0
	a := 10
	// 8 进制默认首位是 0
	b := 06
	// 16 进制默认首位是0x
	c := 0xff
	fmt.Printf("%d  %b\n", a, a)
	fmt.Printf("%o  %b\n", b, b)

	// %[1]o 表示还是使用的是 b
	// %#[1]o 表示还是使用的是 b, 并且添加对应的默认前缀 (16 进制对应可能是 %x 或者是 %X -> 对应是 0x 或者是 0X)
	fmt.Printf("%d %[1]o %#[1]o\n", b)

	fmt.Printf("%x  %b\n", c, c)

}

// Usually a Printf format string containing multiple % verbs
// would require the same number of extra operands, but the [1] ‘‘adverbs’’ after % tell Printf to
// use the first operand over and over again. Second, the # adverb for %o or %x or %X tells Printf
// to emit a 0 or 0x or 0X prefix respectively.
func TestFormat1(t *testing.T) {
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
	// Output:
	// 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF
}

func TestFormat2(t *testing.T) {
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
	fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
	fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"
}
