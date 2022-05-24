package ch08

import (
	"fmt"
	"testing"
	"time"
)

func TestBufferedChannelLength(t *testing.T) {
	channel := make(chan int, 4)
	fmt.Println(len(channel), cap(channel))
	channel <- 1
	channel <- 2
	channel <- 3
	fmt.Println(len(channel), cap(channel))
	close(channel)
}

var done = make(chan int)

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func TestCancellation(t *testing.T) {
	go func() {
		for i := 0; i < 10; i++ {
			if cancelled() {
				fmt.Println("It is cancelled")
			}
		}
	}()
	close(done)
	time.Sleep(10 * time.Second)
}

func TestClose(t *testing.T) {
	ch := make(chan struct{})
	close(ch)
	x, ok := <-ch

	fmt.Println("received: ", x)

	x, ok = <-ch
	if !ok {
		fmt.Println("channel closed, data invalid.")
	}
}
