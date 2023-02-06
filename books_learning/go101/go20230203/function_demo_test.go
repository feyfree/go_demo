package go20230203

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestFunction1(t *testing.T) {
	m, n := SquaresOfSumAndDiff(1, 2)
	fmt.Println(m, n)
}

func TestLambda(t *testing.T) {
	Lambda()
}

func TestFormat(t *testing.T) {
	rand.Seed(time.Now().UnixNano()) // 设置随机数种子
	fmt.Printf("下一个伪随机数是%v。\n", rand.Uint32())

	a, b := 123, "Go"
	fmt.Printf("a == %v == 0x%x, b == %s\n", a, a, b)
	fmt.Printf("type of a: %T, type of b: %T\n", a, b)
	fmt.Printf("1%% 50%% 99%%\n")
}
