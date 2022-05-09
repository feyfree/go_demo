package cf

import (
	"fmt"
	"go_demo/books_learning/gopl/ch02/tempconv"
	"os"
	"strconv"
	"testing"
)

func TestImport(t *testing.T) {
	// 这地方注意 goland 的testing 的 command 有效的不是从1开始
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
