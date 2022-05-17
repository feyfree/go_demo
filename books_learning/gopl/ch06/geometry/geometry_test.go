package geometry

import (
	"fmt"
	"testing"
)

func TestDistance(t *testing.T) {
	p := &Point{X: 0, Y: 0}
	q := &Point{X: 1, Y: 1}
	pq := p.Distance(*q)
	fmt.Println(pq)
	path := Path{*p, *q}
	fmt.Println(path.Distance())
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance()) // "12"
}
