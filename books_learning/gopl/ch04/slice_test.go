package ch04

import (
	"fmt"
	"testing"
)

func TestSliceStruct(t *testing.T) {
	a := [...]int{1, 2, 3}
	b := a[:1]
	fmt.Println(len(b), cap(b)) // 1 3

}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func TestReverse(t *testing.T) {
	array := [...]int{1, 2, 3}
	slice := array[:]
	reverse(slice)
	fmt.Println(slice) // [3 2 1]
	fmt.Println(array) // [3 2 1]
}

// 判断一个slice是不是 empty, 用 len(s) 去判断
func TestExpression(t *testing.T) {
	var s []int // len(s) == 0, s == nil
	fmt.Println(len(s), s == nil)
	s = nil // len(s) == 0, s == nil
	fmt.Println(len(s), s == nil)
	s = []int(nil) // len(s) == 0, s == nil
	fmt.Println(len(s), s == nil)
	s = []int{} // len(s) == 0, s != nil
	fmt.Println(len(s), s == nil)
}

// 使用make 构建指定类型的slice
func TestMake(t *testing.T) {
	a := make([]int, 10)
	fmt.Println(len(a), cap(a))
	fmt.Println(len(a), cap(a))
	b := make([]string, 10, 16)
	fmt.Println(len(b), cap(b))
}

// 一些操作
func TestOperations(t *testing.T) {
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o' ',''''世' '界']"
}

func appendSlice(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// There is room to expand the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}

func TestAppendSlice(t *testing.T) {
	x := []int{1, 2, 3}
	slice := appendSlice(x, 1, 2, 3)
	fmt.Println(slice)
}

//!+append
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// There is room to grow.  Extend the slice.
		z = x[:zlen]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}

//!+growth
func TestAppendGrowth(t *testing.T) {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d  cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}

//!-growth

/*
//!+output
0  cap=1   [0]
1  cap=2   [0 1]
2  cap=4   [0 1 2]
3  cap=4   [0 1 2 3]
4  cap=8   [0 1 2 3 4]
5  cap=8   [0 1 2 3 4 5]
6  cap=8   [0 1 2 3 4 5 6]
7  cap=8   [0 1 2 3 4 5 6 7]
8  cap=16  [0 1 2 3 4 5 6 7 8]
9  cap=16  [0 1 2 3 4 5 6 7 8 9]
//!-output
*/

func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0] // zero-length slice of original
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func TestNonempty(t *testing.T) {
	data := []string{"one", "", "three"}
	fmt.Printf("%q\n", nonempty(data))  // `["one" "three"]`
	fmt.Printf("%q\n", data)            // `["one" "three" "three"]`
	fmt.Printf("%q\n", nonempty2(data)) // `["one" "three" "three"]`
}
