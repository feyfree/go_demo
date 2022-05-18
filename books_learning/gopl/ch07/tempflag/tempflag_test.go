// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 181.

// Tempflag prints the value of its -temp (temperature) flag.
package main

import (
	"flag"
	"fmt"
	"go_demo/books_learning/gopl/ch07/tempconv"
	"testing"
)

//!+
var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func TestTempFlag(t *testing.T) {
	flag.Parse()
	fmt.Println(*temp)
}

//!-
