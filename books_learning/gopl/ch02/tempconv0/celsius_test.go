package tempconv0

import (
	"fmt"
	"testing"
)

func TestOne(t *testing.T) {
	{
		//!+arith
		fmt.Printf("%g\n", BoilingC-FreezingC) // "100" °C
		boilingF := CToF(BoilingC)
		fmt.Printf("%g\n", boilingF-CToF(FreezingC)) // "180" °F
		//!-arith
	}
	/*
		//!+arith
		fmt.Printf("%g\n", boilingF-FreezingC)       // compile error: type mismatch
		//!-arith
	*/

	// Output:
	// 100
	// 180
}

func TestTwo(t *testing.T) {
	//!+printf
	c := FToC(212.0)
	fmt.Println(c.String()) // "100°C"
	fmt.Printf("%v\n", c)   // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", c)   // "100°C"
	fmt.Println(c)          // "100°C"
	fmt.Printf("%g\n", c)   // "100"; does not call String
	fmt.Println(float64(c)) // "100"; does not call String
	//!-printf

	// Output:
	// 100°C
	// 100°C
	// 100°C
	// 100°C
	// 100
	// 100
}


// 会有type miss 编译就会报错
func Test3(t *testing.T) {
	//c := FToC(212.0)
	//e := CToF(100)
	//fmt.Println(c == e)
}
