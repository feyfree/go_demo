package snippets

import (
	"fmt"
	"testing"
)

func TestFib(t *testing.T) {
	fmt.Println(fib(10))
}

func TestGcd(t *testing.T) {
	fmt.Println(gcd(2, 4))
}
