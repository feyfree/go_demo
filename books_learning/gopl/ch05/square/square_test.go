package square

import (
	"fmt"
	"testing"
)

func TestSquare(t *testing.T) {
	// f 相当于是一个
	f := squares()
	fmt.Println(f()) // "1"
	fmt.Println(f()) // "4"
	fmt.Println(f()) // "9"
	fmt.Println(f()) // "16"
}
