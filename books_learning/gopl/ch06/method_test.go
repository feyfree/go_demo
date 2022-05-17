package ch06

import (
	"fmt"
	"go_demo/books_learning/gopl/ch06/geometry"
	"net/url"
	"testing"
	"time"
)

func TestDistance(t *testing.T) {
	perim := geometry.Path{{1, 1}, {5, 1}, {5, 4}, {1, 1}}
	fmt.Println(geometry.PathDistance(perim)) // "12", standalone function
	fmt.Println(perim.Distance())             // "12", method of geometry.Path
}

type IntList struct {
	Value int
	Tail  *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

func TestIntList(t *testing.T) {
	a := &IntList{Value: 1}
	b := &IntList{Value: 2, Tail: a}
	fmt.Println(b.Sum())
}

func TestUrlValues(t *testing.T) {
	m := url.Values{"lang": {"en"}} // direct construction
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang")) // "en"
	fmt.Println(m.Get("q"))    // ""
	fmt.Println(m.Get("item")) // "1"      (first value)
	fmt.Println(m["item"])     // "[1 2]"  (direct map access)

	m = nil
	fmt.Println(m.Get("item")) // ""
	m.Add("item", "3")         // panic: assignment to entry in nil map
}

type Rocket struct {
}

func (r *Rocket) Launch() func() {
	var times int
	return func() {
		times++
		fmt.Printf("Rocket Launch now: %d\n", times)
	}
}

func TestExpressionAndValue(t *testing.T) {
	p := geometry.Point{X: 1, Y: 2}
	q := geometry.Point{X: 4, Y: 6}
	distanceFromP := p.Distance        // method value
	fmt.Println(distanceFromP(q))      // "5"
	var origin geometry.Point          // {0, 0}
	fmt.Println(distanceFromP(origin)) // "2.23606797749979", ;5

	r := new(Rocket)
	f := r.Launch()
	time.AfterFunc(10*time.Second, func() {
		f()
	}) // Rocket Launch now: 2
	time.AfterFunc(1*time.Second, f) // Rocket Launch now: 1
	time.Sleep(15 * time.Second)
}
