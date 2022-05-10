package numbers

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

func TestType0(t *testing.T) {
	const e = 2.71828
	fmt.Println(reflect.TypeOf(e)) // float64

	var f float32 = 16777216 // 1 << 24
	// ide 判定是false, 但是实际上是相等的 float32 可能存在溢出
	fmt.Println(f == f+1)
}

func TestFormat(t *testing.T) {
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d ex = %8.3f\n", x, math.Exp(float64(x)))
	}
}
