// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 357.

// Package unsafeptr demonstrates basic use of unsafe.Pointer.
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//!+main
	var x struct {
		a bool
		b int16
		c []int
	}

	// equivalent to pb := &x.b
	pb := (*int16)(unsafe.Pointer(
		uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)))
	*pb = 42

	fmt.Println(x.b) // "42"
	//!-main
}

/*
//!+wrong
	// NOTE: subtly incorrect!
	// 因为 uintptr 仅仅是一个number， 万一出现GC 的时候， 它还是指向原来的地址
	tmp := uintptr(unsafe.Pointer(&x)) + unsafe.Offsetof(x.b)
	pb := (*int16)(unsafe.Pointer(tmp))
	*pb = 42
//!-wrong
*/
