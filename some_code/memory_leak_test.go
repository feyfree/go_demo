package some_code

import (
	"strings"
	"testing"
)

func TestCloneString(t *testing.T) {

	a := "123"
	// 相对高效, 可以防止底层内存块 因为共享而无法被回收， 导致非substring 占用内存
	strings.Clone(a[:1])
}

var s0 []int

func g(s1 []int) {
	// 假设s1的长度远大于30。
	s0 = s1[len(s1)-30:]
}

func gWithoutLeak(s1 []int) {
	s0 = make([]int, 30)
	copy(s0, s1[len(s1)-30:])
}
