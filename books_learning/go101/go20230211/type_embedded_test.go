package go20230211

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestEmbedded(t *testing.T) {
	type P = *bool
	type M = map[int]int
	var x struct {
		string // 一个具名非指针类型
		error  // 一个具名接口类型
		*int   // 一个无名指针类型
		P      // 一个无名指针类型的别名
		M      // 一个无名类型的别名

		http.Header // 一个具名映射类型
	}
	x.string = "Go"
	x.error = nil
	x.int = new(int)
	x.P = new(bool)
	x.M = make(M)
	x.Header = http.Header{}
}

func TestPerson(t *testing.T) {
	var gaga = Singer{Person: Person{"Gaga", 30}}
	gaga.PrintName()
	gaga.Name = "Lady, Gaga"
	gaga.SetAge(15)
	gaga.PrintName()
	fmt.Println(gaga.Age)
}

func TestReflection(t *testing.T) {
	st := reflect.TypeOf(Singer{}) // the Singer type
	fmt.Println(st, "has", st.NumField(), "fields:")
	for i := 0; i < st.NumField(); i++ {
		fmt.Print(" field#", i, ": ", st.Field(i).Name, "\n")
	}
	fmt.Println(st, "has", st.NumMethod(), "methods:")
	for i := 0; i < st.NumMethod(); i++ {
		fmt.Print(" method#", i, ": ", st.Method(i).Name, "\n")
	}

	pt := reflect.TypeOf(&Singer{}) // th e *Singer type
	fmt.Println(pt, "has", pt.NumMethod(), "methods:")
	for i := 0; i < pt.NumMethod(); i++ {
		fmt.Print(" method#", i, ": ", pt.Method(i).Name, "\n")
	}

}

type I interface {
	m()
}

type T struct {
	I
}

func TestDeadLoop(a *testing.T) {
	var t T
	// i := &t
	// t.I = i
	// i.m()
	t.I = &t
	t.m() // 将调用t.m()，然后再次调用i.m()，......
}
