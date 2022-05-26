package ch12

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflection0(t *testing.T) {
	a := 3
	typeA := reflect.TypeOf(a)
	fmt.Printf("%T", typeA)
	valueA := reflect.ValueOf(a)
	fmt.Println(valueA)

}

func TestAddressable(t *testing.T) {
	x := 2                   // value type variable?
	a := reflect.ValueOf(2)  // 2 int no
	b := reflect.ValueOf(x)  // 2 int no
	c := reflect.ValueOf(&x) // &x *int no
	d := c.Elem()            // 2 int yes (x)

	fmt.Println(a.CanAddr()) // "false"
	fmt.Println(b.CanAddr()) // "false"
	fmt.Println(c.CanAddr()) // "false"
	fmt.Println(d.CanAddr()) // "true"
}

func TestModify(t *testing.T) {
	x := 2
	d := reflect.ValueOf(&x).Elem()   // d refers to the variable x
	px := d.Addr().Interface().(*int) // px := &x
	*px = 3                           // x = 3
	fmt.Println(x)                    // "3"
	d.Set(reflect.ValueOf(4))
	fmt.Println(x)
	d.Set(reflect.ValueOf(int64(5))) // panic: int64 is not assignable to int

	b := reflect.ValueOf(x)
	b.Set(reflect.ValueOf(3)) // panic: Set using unaddressable value
}

type Result struct {
	Data    string
	ErrCode int
}

func TestVisitField(t *testing.T) {
	ptr := &Result{Data: "123", ErrCode: 0}
	v := reflect.ValueOf(ptr).Elem()
	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i)
		fmt.Println(field.Type)
	}
}

func TestReflection1(t *testing.T) {
	type A = [16]int16
	var c <-chan map[A][]byte
	tc := reflect.TypeOf(c)
	fmt.Println(tc.Kind())    // chan
	fmt.Println(tc.ChanDir()) // <-chan
	fmt.Println(reflect.TypeOf(tc).Kind())
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

func TestReflection2(t *testing.T) {
	tp := reflect.TypeOf(new(interface{}))
	tt := reflect.TypeOf(T{})
	fmt.Println(tp.Kind(), tt.Kind()) // ptr slice

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

func TestReflection3(t *testing.T) {
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
	fmt.Println(ok1, ok2) // fa
}
