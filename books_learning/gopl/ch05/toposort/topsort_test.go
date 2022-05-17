package main

import (
	"fmt"
	"testing"
)

//!+main
func TestTopsort(t *testing.T) {
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}
