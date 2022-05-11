package strings

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func basename1(s string) string {
	// Discard last '/' and everything before.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1 if "/" not found
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

//!+
// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func TestComma(t *testing.T) {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

func TestBasename1(t *testing.T) {
	fmt.Println(basename1("a/b/c.go")) // "c"
	fmt.Println(basename1("c.d.go"))   // "c.d"
	fmt.Println(basename1("abc"))      // "abc"

	fmt.Println(basename2("a/b/c.go")) // "c"
	fmt.Println(basename2("c.d.go"))   // "c.d"
	fmt.Println(basename2("abc"))      // "abc"
}
