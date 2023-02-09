package go20230209

import (
	"fmt"
	"testing"
)

func TestFilterFunc_Filter(t *testing.T) {
	var f FilterFunc
	f = func(in int) bool {
		fmt.Println(in)
		return true
	}
	f(1)
	f.Filter(1)
}
