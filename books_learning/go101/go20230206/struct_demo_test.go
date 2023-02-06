package go20230206

import (
	"fmt"
	"testing"
)

func TestStruct(t *testing.T) {
	book := Book{"Go语言101", "老貘", 256}
	fmt.Println(book) // {Go语言101 老貘 256}

	// 使用带字段名的组合字面量来表示结构体值。
	book = Book{author: "老貘", pages: 256, title: "Go语言101"}
	// title和author字段的值都为空字符串""，pages字段的值为0。
	book = Book{}
	// title字段空字符串""，pages字段为0。
	book = Book{author: "老貘"}

	// 使用选择器来访问和修改字段值。
	var book2 Book // <=> book2 := Book{}
	book2.author = "Tapir"
	book2.pages = 300
	fmt.Println(book2.pages) // 300

	book3 := new(Book)
	book3.title = "GO"
	book3.author = "GO"
	book3.pages = 10
	fmt.Println(*book3)
}

func TestStruct2(t *testing.T) {
	var _ = Book{
		author: "老貘",
		pages:  256,
		title:  "Go语言101", // 这里行尾的逗号不可省略
	}

	// 下行}前的逗号可以省略。
	var _ = Book{author: "老貘", pages: 256, title: "Go语言101"}
}

func TestStruct3(t *testing.T) {
	book1 := Book{pages: 300}
	book2 := Book{"Go语言101", "老貘", 256}

	book2 = book1
	// 上面这行和下面这三行是等价的。
	book2.title = book1.title
	book2.author = book1.author
	book2.pages = book1.pages

	fmt.Printf("%p\n", &book2)
	fmt.Printf("%p\n", &book1)
	fmt.Println(book1 == book2)
}

func TestStructAddressable(t *testing.T) {
	type Book struct {
		Pages int
	}
	var book = Book{} // 变量值book是可寻址的
	p := &book.Pages
	*p = 123
	fmt.Println(book) // {123}

	// 下面这两行编译不通过，因为Book{}是不可寻址的，
	// 继而Book{}.Pages也是不可寻址的。
	/*
		Book{}.Pages = 123
		p = &Book{}.Pages // <=> p = &(Book{}.Pages)
	*/
}

type S0 struct {
	y int "foo"
	x bool
}

type S1 = struct { // S1是一个无名类型
	x int "foo"
	y bool
}

type S2 = struct { // S2也是一个无名类型
	x int "bar"
	y bool
}

type S3 S2 // S3是一个定义类型（因而具名）。
type S4 S3 // S4是一个定义类型（因而具名）。
// 如果不考虑字段标签，S3（S4）和S1的底层类型一样。
// 如果考虑字段标签，S3（S4）和S1的底层类型不一样。

var v0, v1, v2, v3, v4 = S0{}, S1{}, S2{}, S3{}, S4{}

func TestConversion(t *testing.T) {
	v1 = S1(v2)
	v2 = S2(v1)
	v1 = S1(v3)
	v3 = S3(v1)
	v1 = S1(v4)
	v4 = S4(v1)
	v2 = v3
	v3 = v2 // 这两个转换可以是隐式的
	v2 = v4
	v4 = v2 // 这两个转换也可以是隐式的
	v3 = S3(v4)
	v4 = S4(v3)
}
