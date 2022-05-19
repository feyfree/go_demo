package ch07

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestInterfaceSatisfy(t *testing.T) {
	var w io.Writer
	w = os.Stdout
	w.Write([]byte("hello"))
	// 注意  w 是 io.Writer 类型， 没有Close方法
	// 如果调用了 w.Close() 会无法编译
	// w.Close()
}

// empty interface mean that we can assign any value to the empty interface
func TestEmptyInterface(t *testing.T) {
	var any interface{}
	any = true
	any = 12.34
	any = "hello"
	any = map[string]int{"one": 1}
	fmt.Println(any)
	any = new(bytes.Buffer)
	fmt.Println(any)
}

func TestInterface(t *testing.T) {
	// 这地方注意不是指针变量
	var w io.Writer
	fmt.Printf("%T\n", w) // "<nil>"
	// os.Stdout 是 concrete type
	// w 是 interface type
	// 这个实际上存在一个 隐含的转换
	w = os.Stdout
	w.Write([]byte("hello"))
	fmt.Printf("%T\n", w) // "*os.File"

	// 显示转换
	w = io.Writer(os.Stdout)
	w.Write([]byte("hello"))

	w = new(bytes.Buffer)
	fmt.Printf("%T\n", w) // "*bytes.Buffer"
	w.Write([]byte("hello"))

	var x interface{} = []int{1, 2, 3}
	fmt.Println(x == x) // panic: comparing uncomparable type []int
}

// assign the nil pointer to the interface was a mistake
// nil pointer is not nil
func TestNilPointer(t *testing.T) {
	// 这地方注意是指针变量, 最好还是用 bytes.Buffer 然后 new
	var buf *bytes.Buffer
	fmt.Printf("%T\n", buf) // *bytes.Buffer
	//buf = new(bytes.Buffer)

	// 如果不通过 new(bytes.Buffer)操作的话， 会panic
	// 相当于是 有 type 无 value
	f(buf)
}

func f(out io.Writer) {
	fmt.Printf("%T\n", out) // *bytes.Buffer
	if out != nil {
		out.Write([]byte("Hello"))
	}
}

// type assertion like x.(T)
// 1. 如果 T 是具体类型， type assertion 会检查 x 的 dynamic value 是不是 T， 如果是会返回 x 的 dynamic value, 否则会报错(使用二元组可以避免 panic)
// 2. 如果 T 是interface 类型, 会检查 x 的 dynamic value 会不会 satisfy T
//
func TestTypeAssertions(t *testing.T) {
	var w io.Writer
	w = os.Stdout
	_, ok := w.(*os.File)
	if ok {
		fmt.Println("w holds : os.Stdout")
	}
	// panic: interface holds *os.File, not *bytes.Buffer
	//c := w.(*bytes.Buffer)

	_, right := w.(*bytes.Buffer)
	if !right {
		fmt.Println("w holds : not bytes.Buffer")
	}

	//y := os.Stdout
	// 这样写的话会编译不通过 invalid type assertion: y.(io.Writer) (non-interface type *os.File on left)
	// Type switches require an interface to introspect.
	// If you are passing a value of known type to it it bombs out.
	// If you make a function that accepts an interface as a parameter, it will work
	//_, isWriter := y.(io.Writer)
	//fmt.Println(isWriter)
}

func TestTypeAssertions2(t *testing.T) {

	var w io.Writer
	w = os.Stdout
	_ = w.(io.ReadWriter) // success: *os.File has both Read and Write

	//w = new(bytecounter.ByteCounter)
	// panic: *ByteCounter has no Read method
	//_ = w.(io.ReadWriter)

	rw := w.(io.ReadWriter)
	fmt.Printf("%T \n", rw)
}
