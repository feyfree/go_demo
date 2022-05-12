package ch04

import (
	"fmt"
	"sort"
	"testing"
)

func TestMaps(t *testing.T) {

	ages := make(map[string]int) // mapping from strings to ints
	fmt.Println(ages)
	ages = map[string]int{
		"alice":   31,
		"charlie": 34,
	}
	fmt.Println(ages)
	delete(ages, "alice")
	fmt.Println(ages)
	ages["bob"] = ages["bob"] + 1
	fmt.Println(ages)
	// 用来判定是否存在一个 key
	_, ok := ages["allen"]
	if !ok {
		fmt.Println("Not contain allen")
	}
	for name, age := range ages {
		fmt.Printf("%s\t%d\n", name, age)
	}
	// 会报编译错误 Cannot take the address of 'ages["bob"]', 原因是map element 中的value 不能取地址
	// i := &ages["bob"]

	var names []string
	for name := range ages {
		names = append(names, name)
	}
	fmt.Println("----Next print sorted names with age----")
	sort.Strings(names)
	for _, name := range names {
		fmt.Printf("%s\t%d\n", name, ages[name])
	}
}

func TestMaps2(t *testing.T) {
	var ages map[string]int
	fmt.Println(ages == nil)    // "true"
	fmt.Println(len(ages) == 0) // "true"
	// nil map 不能分配entry
	//ages["carol"] = 21          // panic: assignment to entry in nil map
}

func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

func TestEqual(t *testing.T) {
	x := map[string]int{"a": 1, "b": 2}
	y := map[string]int{"a": 1, "c": 2}
	fmt.Println(equal(x, y))
}
