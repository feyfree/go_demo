package coloredpoint

import (
	"fmt"
	"image/color"
	"testing"
)

func TestColoredPoint(t *testing.T) {

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = ColoredPoint{Point{1, 1}, red}
	fmt.Printf("Point-P:(%v, %v) \n", p.X, p.Y)
	var q = ColoredPoint{Point{5, 4}, blue}

	// 实际上这种 embedding struct 是能直接调用内部 嵌入对象的 相关方法的
	// 这个实际是在编译的时候， 会将这些方法 按照封装的对象的类型重写一份给这个封装的对象 (入参和原来的函数入参是一致的)
	// 但是注意类型和签名保持一致
	fmt.Println(p.Distance(q.Point)) // "5"
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point)) // "10"

	// 也可以这么写
	pd := p.Distance
	fmt.Println(pd(q.Point))

}
