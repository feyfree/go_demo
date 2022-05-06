package echo2


import (
	"fmt"
	"os"
)

// 0 一般是运行的脚本本身
func echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}