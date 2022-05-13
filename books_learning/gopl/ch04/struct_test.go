package ch04

import (
	"fmt"
	"go_demo/books_learning/gopl/ch04/structs"
	"testing"
)

func TestStruct(t *testing.T) {
	// 下面会编译报错， 因为 a, b 小写是不能被exported的
	//t := structs.T{1, 1}

	// 这样写也是可以的， 但是不推荐
	//e := structs.E{1, 2}
	e := &structs.E{X: 1, Y: 2}
	fmt.Println(*e)
}

// 匿名的话 struct 继承内部的 匿名部分的 sub fields
// 使用匿名 可以使用类似直接的sub fields 赋值， 但是无法进行直接的 构造
func TestAnonymousFields(t *testing.T) {
	var w structs.Wheel
	w.X = 8      // equivalent to w.Circle.Point.X = 8
	w.Y = 8      // equivalent to w.Circle.Point.Y = 8
	w.Radius = 5 // equivalent to w.Circle.Radius = 5
	w.Spokes = 20
	fmt.Println(w)

	w = structs.Wheel{
		Circle: structs.Circle{
			Point:  structs.Point{X: 8, Y: 8},
			Radius: 5,
		},
		Spokes: 20, // NOTE: trailing comma necessary here (and at Radius)
	}
	fmt.Printf("%#v\n", w)
	w.X = 42
	fmt.Printf("%#v\n", w)

}
