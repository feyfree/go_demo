package go20230210

import "fmt"

type Aboutable interface {
	About() string
}

// 类型*Book实现了接口类型Aboutable。
type Book struct {
	name string
}

func (book *Book) About() string {
	return "Book: " + book.name
}

type Greeting interface {
	SayHello()
	SayGoodBye()
}

type BrotherGreeting struct{}

func (BrotherGreeting) SayHello() {
	fmt.Println("你好, 兄弟")
}

func (BrotherGreeting) SayGoodBye() {
	fmt.Println("兄弟, 再见")
}

type FriendGreeting struct {
}

func (FriendGreeting) SayHello() {
	fmt.Println("你好, 朋友")
}

func (FriendGreeting) SayGoodBye() {
	fmt.Println("朋友， 再见")
}
