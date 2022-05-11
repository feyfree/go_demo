package constants

import (
	"fmt"
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
	fmt.Println(KiB, MiB, GiB, TiB, PiB)
}
