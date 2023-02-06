package go20230206

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	// var a uint = 1

	// var _ = []int{a: 100}  // error: 下标必须为常量
	// var _ = [5]int{a: 100} // error: 下标必须为常量
}

func TestMake(t *testing.T) {
	// 创建映射。
	fmt.Println(make(map[string]int)) // map[]
	m := make(map[string]int, 3)
	fmt.Println(m, len(m)) // map[] 0
	m["C"] = 1972
	m["Go"] = 2009
	fmt.Println(m, len(m)) // map[C:1972 Go:2009] 2

	// 创建切片。
	s := make([]int, 3, 5)
	fmt.Println(s, len(s), cap(s)) // [0 0 0] 3 5
	s = make([]int, 2)
	fmt.Println(s, len(s), cap(s)) // [0 0] 2 2
}

func TestNew(t *testing.T) {
	m := *new(map[string]int)   // <=> var m map[string]int
	fmt.Println(m == nil)       // true
	s := *new([]int)            // <=> var s []int
	fmt.Println(s == nil)       // true
	a := *new([5]bool)          // <=> var a [5]bool
	fmt.Println(a == [5]bool{}) // true

	n := new(map[string]int)
	fmt.Println(n)
	// n 是个 nil 指针
	(*n)["hello"] = 1
	fmt.Println(*n)
}

func TestMap(t *testing.T) {
	type T struct{ age int }
	mt := map[string]T{}
	mt["John"] = T{age: 29} // 整体修改是允许的
	ma := map[int][5]int{}
	ma[1] = [5]int{1: 789} // 整体修改是允许的

	// 这两个赋值编译不通过，因为部分修改一个映射
	// 元素是非法的。这看上去确实有些反直觉。
	/*
		ma[1][1] = 123      // error
		mt["John"].age = 30 // error
	*/
	// 读取映射元素的元素或者字段是没问题的。
	fmt.Println(ma[1][1])       // 789
	fmt.Println(mt["John"].age) // 29

}
