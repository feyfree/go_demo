package strings

import (
	"fmt"
	"testing"
)

func TestPrintints(t *testing.T) {
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
}
