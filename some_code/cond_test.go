package some_code

import (
	"fmt"
	"strconv"
	"sync"
	"testing"
	"time"
)

type ShareResource struct {
	number int
	lock   sync.Locker
	c1     *sync.Cond
	c2     *sync.Cond
	c3     *sync.Cond
}

func TestCondShareResource(t *testing.T) {
	resource := &ShareResource{}
	resource.number = 1
	resource.lock = &sync.Mutex{}
	resource.c1 = sync.NewCond(resource.lock)
	resource.c2 = sync.NewCond(resource.lock)
	resource.c3 = sync.NewCond(resource.lock)
	print5 := func(operator string) {
		resource.lock.Lock()
		defer resource.lock.Unlock()
		for resource.number != 1 {
			resource.c1.Wait()
		}
		for i := 0; i < 5; i++ {
			fmt.Printf(operator + "\t" + strconv.Itoa(i) + "\n")
		}
		resource.number = 2
		resource.c2.Signal()
	}

	print10 := func(operator string) {
		resource.lock.Lock()
		defer resource.lock.Unlock()
		for resource.number != 2 {
			resource.c2.Wait()
		}
		for i := 0; i < 10; i++ {
			fmt.Printf(operator + "\t" + strconv.Itoa(i) + "\n")
		}
		resource.number = 3
		resource.c3.Signal()
	}

	print15 := func(operator string) {
		resource.lock.Lock()
		defer resource.lock.Unlock()
		for resource.number != 3 {
			resource.c3.Wait()
		}
		for i := 0; i < 15; i++ {
			fmt.Printf(operator + "\t" + strconv.Itoa(i) + "\n")
		}
		resource.number = 1
		resource.c1.Signal()
	}

	go func() {
		for i := 0; i < 5; i++ {
			print5("A")
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			print10("B")
		}
	}()

	go func() {
		for i := 0; i < 15; i++ {
			print15("C")
		}
	}()

	//go print5("A")
	//go print10("B")
	//go print15("C")

	time.Sleep(time.Second * 3)
}
