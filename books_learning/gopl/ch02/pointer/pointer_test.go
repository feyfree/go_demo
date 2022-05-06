package pointer

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {
	fmt.Println(f() == f())
	fmt.Println(*f() == *f())
}

func TestIncr(t *testing.T) {
	v := 1
	fmt.Printf("%p\n", &v)
	incr(&v)
	fmt.Println(incr(&v))
	fmt.Printf("%p", &v)

}

func TestNewInt(t *testing.T) {
	p := newInt()
	fmt.Println(*p)
}
