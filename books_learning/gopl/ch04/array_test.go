package ch04

// Go 默认是值传递 所以如果用数组的话， 尽量用指针， 避免复制的开销

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func TestArray0(t *testing.T) {
	symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	fmt.Println(RMB, symbol[RMB])

	// index = 3, value = 1 说明长度是 4， 并且 在 index < 3 的地方都是 0
	r := [...]int{3: 1} // [0 0 0 1]
	fmt.Println(r)

	a := [2]int{1, 2}
	b := [...]int{1, 2}
	c := [2]int{1, 3}
	fmt.Println(a == b, a == c, b == c) // "true false false"
	fmt.Println(&a[0], &b[0])
	d := [3]int{1, 2}
	// 这个会报编译错误 compile error: cannot compare [2]int == [3]int }
	// fmt.Println(a == d) //
	fmt.Println(d)
}

func TestArraySha256(t *testing.T) {
	// 注意这个是小写的 x 和 大写的 X 的sha256 的比较
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	// Output:
	// 2d711642b726b04401627ca9fbac32f5c8530fb1903cc4db02258717921a4881
	// 4b68ab3847feda7d6c62c1fbcbeebfa35eab7351ed5e78f4ddadea5df64b8015
	// false
	// [32]uint8
}
