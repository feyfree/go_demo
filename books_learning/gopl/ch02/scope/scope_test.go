package scope

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestVariableScope1(t *testing.T) {
	x := "hello"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x)
	}
}

func f() int {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	return random.Intn(100)
}

func g(x int) int {
	return x ^ x
}

func TestVariableScope2(t *testing.T) {
	if x := f(); x == 0 {
		fmt.Println("--1--")
		fmt.Println(x)
	} else if y := g(x); x == y {
		fmt.Println("--2--")
		fmt.Println(x, y)
	} else {
		fmt.Println("--3--")
		fmt.Println(x, y)
	}
	// if-else 之外实际是访问不到的
	// fmt.Println(x, y)
}
