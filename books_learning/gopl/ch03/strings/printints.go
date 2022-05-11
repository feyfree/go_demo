package strings

import (
	"bytes"
	"fmt"
)

// intsToString is like fmt.Sprintf(values) but adds commas.

// When appending the UTF-8 encoding of an arbitrary rune to a bytes.Buffer, it’s best to use
// bytes.Buffer’s WriteRune method, but WriteByte is fine for ASCII characters such as '['
// and ']'.
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	// 这种写法会编译报错， byte 是放不下 这个 rune的数据的
	//buf.WriteByte('国')
	buf.WriteByte(']')
	return buf.String()
}
