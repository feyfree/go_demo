package structs

// T 如果内部field是小写的话， field是不能被外部包直接引用的
type T struct {
	a, b int
}

type E struct {
	X, Y int
}

type Point struct {
	X, Y int
}

//type Circle struct {
//	Center Point
//	Radius int
//}
//type Wheel struct {
//	Circle Circle
//	Spokes int
//}

// Circle 匿名写法
type Circle struct {
	Point
	Radius int
}

type Wheel struct {
	Circle
	Spokes int
}
