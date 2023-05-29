package go20230212

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect01(t *testing.T) {
	type A = [16]int16
	var c <-chan map[A][]byte
	tc := reflect.TypeOf(c)
	fmt.Println(tc.Kind())    // chan
	fmt.Println(tc.ChanDir()) // <-chan
	tm := tc.Elem()
	ta, tb := tm.Key(), tm.Elem()
	fmt.Println(tm.Kind(), ta.Kind(), tb.Kind()) // map array slice
	tx, ty := ta.Elem(), tb.Elem()

	// byte是uint8类型的别名。
	fmt.Println(tx.Kind(), ty.Kind()) // int16 uint8
	fmt.Println(tx.Bits(), ty.Bits()) // 16 8
	fmt.Println(tx.ConvertibleTo(ty)) // true
	fmt.Println(tb.ConvertibleTo(ta)) // false

	// 切片类型和映射类型都是不可比较类型。
	fmt.Println(tb.Comparable()) // false
	fmt.Println(tm.Comparable()) // false
	fmt.Println(ta.Comparable()) // true
	fmt.Println(tc.Comparable()) // true
}

type T []interface{ m() }

func (T) m() {}

func TestReflect02(t *testing.T) {

	tp := reflect.TypeOf(new(interface{}))
	tt := reflect.TypeOf(T{})
	fmt.Println(tp.Kind(), tt.Kind()) // ptr slice

	fmt.Println(reflect.ValueOf(new(interface{})).Kind())

	// 使用间接的方法得到表示两个接口类型的reflect.Type值。
	ti, tim := tp.Elem(), tt.Elem()
	fmt.Println(ti.Kind(), tim.Kind()) // interface interface

	fmt.Println(tt.Implements(tim))  // true
	fmt.Println(tp.Implements(tim))  // false
	fmt.Println(tim.Implements(tim)) // true

	// 所有的类型都实现了任何空接口类型。
	fmt.Println(tp.Implements(ti))  // true
	fmt.Println(tt.Implements(ti))  // true
	fmt.Println(tim.Implements(ti)) // true
	fmt.Println(ti.Implements(ti))  // true
}

type F func(string, int) bool

func (f F) m(s string) bool {
	return f(s, 32)
}
func (f F) M() {}

type I interface {
	m(s string) bool
	M()
}

func TestReflect03(t *testing.T) {
	var x struct {
		F F
		i I
	}
	tx := reflect.TypeOf(x)
	fmt.Println(tx.Kind())        // struct
	fmt.Println(tx.NumField())    // 2
	fmt.Println(tx.Field(1).Name) // i
	// 包路径（PkgPath）是非导出字段（或者方法）的内在属性。
	fmt.Println(tx.Field(0).PkgPath) //
	fmt.Println(tx.Field(1).PkgPath) // main

	tf, ti := tx.Field(0).Type, tx.Field(1).Type
	fmt.Println(tf.Kind())               // func
	fmt.Println(tf.IsVariadic())         // false
	fmt.Println(tf.NumIn(), tf.NumOut()) // 2 1
	t0, t1, t2 := tf.In(0), tf.In(1), tf.Out(0)
	// 下一行打印出：string int bool
	fmt.Println(t0.Kind(), t1.Kind(), t2.Kind())

	fmt.Println(tf.NumMethod(), ti.NumMethod()) // 1 2
	fmt.Println(tf.Method(0).Name)              // M
	fmt.Println(ti.Method(1).Name)              // m
	_, ok1 := tf.MethodByName("m")
	_, ok2 := ti.MethodByName("m")
	fmt.Println(ok1, ok2) // false true
}

func TestReflect04(t *testing.T) {
	ta := reflect.ArrayOf(5, reflect.TypeOf(123))
	fmt.Println(ta) // [5]int
	tc := reflect.ChanOf(reflect.SendDir, ta)
	fmt.Println(tc) // chan<- [5]int
	tp := reflect.PtrTo(ta)
	fmt.Println(tp) // *[5]int
	ts := reflect.SliceOf(tp)
	fmt.Println(ts) // []*[5]int
	tm := reflect.MapOf(ta, tc)
	fmt.Println(tm) // map[[5]int]chan<- [5]int
	tf := reflect.FuncOf([]reflect.Type{ta},
		[]reflect.Type{tp, tc}, false)
	fmt.Println(tf) // func([5]int) (*[5]int, chan<- [5]int)
	tt := reflect.StructOf([]reflect.StructField{
		{Name: "Age", Type: reflect.TypeOf("abc")},
	})
	fmt.Println(tt)            // struct { Age string }
	fmt.Println(tt.NumField()) // 1
}

func TestReflect05(t *testing.T) {
	n := 123
	p := &n
	vp := reflect.ValueOf(&p)
	fmt.Println(vp.CanSet(), vp.CanAddr()) // false false
	vn := vp.Elem()                        // 取得vp的底层指针值引用的值的代表值
	fmt.Println(vn.CanSet(), vn.CanAddr()) // true true
	vn.Set(reflect.ValueOf(789))           // <=> vn.SetInt(789)
	fmt.Println(n)                         // 789
}

// 一个结构体值的非导出字段不能通过反射来修改
func TestReflect06(t *testing.T) {
	var s struct {
		X interface{} // 一个导出字段
		y interface{} // 一个非导出字段
	}
	vp := reflect.ValueOf(&s)
	// 如果vp代表着一个指针，下一行等价于"vs := vp.Elem()"。
	vs := reflect.Indirect(vp)
	// vx和vy都各自代表着一个接口值。
	vx, vy := vs.Field(0), vs.Field(1)
	fmt.Println(vx.CanSet(), vx.CanAddr()) // true true
	// vy is addressable but not modifiable.
	fmt.Println(vy.CanSet(), vy.CanAddr()) // false true
	vb := reflect.ValueOf(123)
	vx.Set(vb) // okay, 因为vx代表的值是可修改的。
	// vy.Set(vb)  // 会造成恐慌，因为vy代表的值是不可修改的。
	fmt.Println(s)                      // {123
	fmt.Println(vx.IsNil(), vy.IsNil()) // false true
}
