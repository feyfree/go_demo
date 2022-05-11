package strings

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestRune(t *testing.T) {
	s := "Hello, 世界"
	fmt.Println(len(s))                    // "13"
	fmt.Println(utf8.RuneCountInString(s)) // "9"

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	for i, r := range s {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	n := 0
	for _, _ = range s {
		n++
	}
	fmt.Println("loop times:", n) // 9

	n = 0
	// 相当于 n = utf8.RuneCountInString(s)
	for range s {
		n++
	}
	fmt.Println("loop times:", n) // 9
}

func TestRuneAndString(t *testing.T) {
	// "program" in Japanese katakana
	s := "プログラム"
	// % x 中间有个空格， 表示每个 digits 中间有个空格
	fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	r := []rune(s)
	fmt.Printf("%x\n", r) // "[30d7 30ed 30b0 30e9 30e0]"

	fmt.Println(string(r))

	// Conversion from int to string interprets an integer value as a code point
	// Inspection info: Reports conversions of the form string(x) where x is an integer (but not byte or rune) type.
	// ide 会飘黄
	//fmt.Println(string(65))     // "A", not "65"
	//fmt.Println(string(0x4eac)) // "C"

	fmt.Println(string(rune(65)))     // "A", not "65"
	fmt.Println(string(rune(0x4eac))) // "京"

	// 无效rune 这地方会报compiler failure
	//fmt.Println(string(1234567))
}
