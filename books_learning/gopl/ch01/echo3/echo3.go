package echo3

import (
	"fmt"
	"os"
	"strings"
)

func echo3() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
