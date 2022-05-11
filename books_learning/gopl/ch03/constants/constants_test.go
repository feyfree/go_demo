package constants

import (
	"fmt"
	"math"
	"reflect"
	"testing"
	"time"
)

const noDelay time.Duration = 0
const timeout = 5 * time.Minute

func TestConstants0(t *testing.T) {
	fmt.Printf("%T %[1]v\n", noDelay)     // "time.Duration 0"
	fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration 5m0s
	fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration 1m0s"
}

const (
	a = 1
	b
	c = 2
	d
)

func TestConstants1(t *testing.T) {
	fmt.Println(a, b, c, d) // "1 1 2 2"
}

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

func TestConstants2(t *testing.T) {
	fmt.Println(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday) // 0 1 2 3 4 5 6
}

type Flags uint

const (
	FlagUp           Flags = 1 << iota // is up
	FlagBroadcast                      // supports broadcast access capability
	FlagLoopBack                       // is a loopback interface
	FlagPointToPoint                   // belongs to a point-to-point link
	FlagMulticast                      // supports multicast access capability
)

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }

func TestFlags(t *testing.T) {
	fmt.Println(FlagUp, FlagBroadcast, FlagLoopBack, FlagPointToPoint, FlagMulticast)
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
}

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776 (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424 (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)

func TestPower(t *testing.T) {
	// debug 的时候可以发现这几个常量都是 untyped int 类型
	fmt.Println(KiB, MiB, GiB, TiB, PiB)
	fmt.Println(YiB / ZiB)
}

// only constants can be untyped
// 一旦被复制给变量，就会有确切的类型
func TestUntypedConstants(t *testing.T) {
	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi
	fmt.Println(x, y, z)
	fmt.Println(math.Pi)

	// reflect.TypeOf 是无法反应untyped constants 的
	fmt.Println(reflect.TypeOf(math.Pi))
	var f float64 = 3 + 0i // untyped complex -> float64
	fmt.Println(reflect.TypeOf(f))
	f = 2 // untyped integer -> float64
	fmt.Println(reflect.TypeOf(f))
	f = 1e123 // untyped floating-point -> float64
	fmt.Println(reflect.TypeOf(f))
	f = 'a' // untyped rune -> float64
	fmt.Println(reflect.TypeOf(f))

}

func TestOperands(t *testing.T) {
	var f float64 = 212
	fmt.Println((f - 32) * 5 / 9)     // "100"; (f - 32) * 5 is a float64
	fmt.Println(5 / 9 * (f - 32))     // "0"; 5/9 is an untyped integer, 0
	fmt.Println(5.0 / 9.0 * (f - 32)) // "100"; 5.0/9.0 is an untyped float
}

func TestImplicit(t *testing.T) {
	i := 0      // untyped integer; implicit int(0)
	r := '\000' // untyped rune; implicit rune('\000')
	f := 0.0    // untyped floating-point; implicit float64(0.0)
	c := 0i     // untyped complex; implicit complex128(0i)
	fmt.Println(i, r, f, c)
}

func TestExplicitType(t *testing.T) {
	// short variable 依赖 go compiler
	// 指明类型的话还是主动显示声明类型
	var i = int8(0)
	// 或者
	// var i int8 = 0
	fmt.Printf("%T\n", i)

	fmt.Printf("%T\n", 0)      // "int"
	fmt.Printf("%T\n", 0.0)    // "float64"
	fmt.Printf("%T\n", 0i)     // "complex128"
	fmt.Printf("%T\n", '\000') // "int32" (rune)

}
