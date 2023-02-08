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

	x := make([]int, 10, 100)
	fmt.Println(x[50:60])
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

func TestArrayRange(t *testing.T) {
	type Person struct {
		name string
		age  int
	}
	persons := [2]Person{{"Alice", 28}, {"Bob", 25}}
	for i, p := range persons {
		fmt.Println(i, p)
		// 此修改将不会体现在这个遍历过程中，
		// 因为被遍历的数组是persons的一个副本。
		persons[1].name = "Jack"

		// 此修改不会反映到persons数组中，因为p
		// 是persons数组的副本中的一个元素的副本。
		p.age = 31
	}
	fmt.Println("persons:", &persons)
}

func TestSliceRange(t *testing.T) {
	type Person struct {
		name string
		age  int
	}
	// 改为一个切片。
	persons := []Person{{"Alice", 28}, {"Bob", 25}}
	for i, p := range persons {
		fmt.Println(i, p)
		// 这次，此修改将反映在此次遍历过程中。
		persons[1].name = "Jack"
		// 这个修改仍然不会体现在persons切片容器中。
		p.age = 31
	}
	fmt.Println("persons:", &persons)

	for i := range persons {
		persons[i].name = "Happy"
	}
	fmt.Println("persons:", &persons)
}

func TestMapRange(t *testing.T) {
	languages := map[struct{ dynamic, strong bool }]map[string]int{
		{true, false}:  {"JavaScript": 1995},
		{false, true}:  {"Go": 2009},
		{false, false}: {"C": 1972},
	}
	// 此映射的键值和元素类型均为指针类型。
	// 这有些不寻常，只是为了讲解目的。
	m0 := map[*struct{ dynamic, strong bool }]*map[string]int{}
	for category, langInfo := range languages {
		m0[&category] = &langInfo
		// 下面这行修改对映射 languages 没有任何影响。
		category.dynamic, category.strong = true, true
	}
	for category, langInfo := range languages {
		fmt.Println(category, langInfo)
	}

	m1 := map[struct{ dynamic, strong bool }]map[string]int{}
	for category, langInfo := range m0 {
		m1[*category] = *langInfo
	}
	// 映射m0和m1中均只有一个条目。
	fmt.Println(len(m0), len(m1)) // 1 1
	fmt.Println(m1)               // map[{true true}:map[C:1972]]
}

func TestArrayPointer(t *testing.T) {
	// 未初始化的数组不是 nil
	var a [100]int

	fmt.Println(a)

	for i, n := range &a { // 复制一个指针的开销很小
		fmt.Println(i, n)
	}

	for i, n := range a[:] { // 复制一个切片的开销很小
		fmt.Println(i, n)
	}

	var p *[5]int // nil

	for i, _ := range p { // okay
		fmt.Println(i)
	}

	for i := range p { // okay
		fmt.Println(i)
	}

	// p 指针类型 默认是nil
	for i, n := range p { // panic
		fmt.Println(i, n)
	}

	var pa *[5]int                // == nil
	fmt.Println(len(pa), cap(pa)) // 5 5
}

func TestAppendOrder(t *testing.T) {
	a := []int{1, 2}
	b := []int{3, 4}
	s := append(a, b...)
	fmt.Println(s)
	d := append([]int(nil), a...)
	fmt.Println(d)

	x := []int{1, 2, 3}
	y := []int{4, 5, 6}
	m := append(x[:1], y[:1]...)
	fmt.Println(m)
}
