package go20230208

import (
	"bytes"
	"fmt"
	"testing"
)

func TestStringDemo(t *testing.T) {
	s := "颜色感染是一个有趣的游戏。"
	bs := []byte(s) // string -> []byte
	fmt.Println(bs)
	s = string(bs)  // []byte -> string
	rs := []rune(s) // string -> []rune
	fmt.Println(rs)
	s = string(rs)       // []rune -> string
	rs = bytes.Runes(bs) // []byte -> []rune
	fmt.Println(rs)
	bs = Runes2Bytes(rs) // []rune -> []byte}
	fmt.Println(bs)
}

func TestForAndForRange(t *testing.T) {
	s := "asc"
	fmt.Printf("%T  ", s[0])
	fmt.Println()
	for i := range s {
		fmt.Printf("%T ", s[i])
	}
	fmt.Println()
	for _, i2 := range s {
		fmt.Printf("%T ", i2)
	}
	fmt.Println()
	for i := 0; i < len(s); i++ {
		fmt.Printf("%T ", s[i])
	}

}

func TestAppend(t *testing.T) {
	hello := []byte("Hello ")
	world := "world!"

	// helloWorld := append(hello, []byte(world)...) // 正常的语法
	helloWorld := append(hello, world...) // 语法糖
	fmt.Println(string(helloWorld))

	helloWorld2 := make([]byte, len(hello)+len(world))
	copy(helloWorld2, hello)
	// copy(helloWorld2[len(hello):], []byte(world)) // 正常的语法
	copy(helloWorld2[len(hello):], world) // 语法糖
	fmt.Println(string(helloWorld2))
}
