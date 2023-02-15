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

func TestBook(t *testing.T) {
	var book Book

	book.SetPages(1)

	fmt.Printf("%T \n", book.Pages)       // func() int
	fmt.Printf("%T \n", (&book).SetPages) // func(int)
	// &book值有一个隐式方法Pages。
	fmt.Printf("%T \n", (&book).Pages) // func() int

	// 调用这三个方法。
	(&book).SetPages(123)
	book.SetPages(123)           // 等价于上一行
	fmt.Println(book.Pages())    // 123
	fmt.Println((&book).Pages()) // 123
}

func TestAge(t *testing.T) {
	_ = (StringSet(nil)).Has   // 不会产生恐慌
	_ = ((*Age)(nil)).IsNil    // 不会产生恐慌
	_ = ((*Age)(nil)).Increase // 不会产生恐慌

	_ = (StringSet(nil)).Has("key") // 不会产生恐慌
	_ = ((*Age)(nil)).IsNil()       // 不会产生恐慌

	// 下面这行将产生一个恐慌，但是此恐慌不是在调用方法的时
	// 候产生的，而是在此方法体内解引用空指针的时候产生的。
	((*Age)(nil)).Increase()
}

type Story struct {
	pages int
}

func (b Story) SetPages(pages int) {
	b.pages = pages
}

type Stories []Story

func (stories Stories) Modify() {
	// 对属主参数的间接部分的修改将反映到方法之外。
	stories[0].pages = 500
	// 对属主参数的直接部分的修改不会反映到方法之外。
	stories = append(stories, Story{789})
}

// 属主传递都是值复制
func TestStory(t *testing.T) {
	var story Story
	story.SetPages(10)
	fmt.Println(story.pages) // 0
}

func TestMyAge(t *testing.T) {
	var x MyInt = 3
	_ = x.IsOdd() // okay

	var y Age = 36
	// _ = y.IsOdd() // error: y.IsOdd undefined
	_ = y

	var m MyInfo
	fmt.Println(m.IsOdd())
}
