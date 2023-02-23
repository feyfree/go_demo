package some_code

import (
	"fmt"
	"sync"
)

func TestLockUp(lock sync.RWMutex) {
	lock.RLock()
	fmt.Println("read lock sucess")
	lock.Lock()
	fmt.Println("write lock success")
	lock.Unlock()
	lock.RUnlock()
}

// fatal error: all goroutines are asleep - deadlock!
func TestLockDown(lock sync.RWMutex) {
	lock.Lock()
	fmt.Println("write lock success")
	lock.RUnlock()
	fmt.Println("read lock sucess")
	lock.RUnlock()
	lock.Unlock()
}

// 相当于加了两遍读锁
func TestReRLock(lock sync.RWMutex) {
	lock.RLock()
	fmt.Println("read lock 1 success")
	lock.RLock()
	fmt.Println("read lock 2 sucess")
	lock.RUnlock()
	lock.RUnlock()
}

// fatal error: all goroutines are asleep - deadlock!
func TestReLock(lock sync.RWMutex) {
	lock.Lock()
	fmt.Println("write lock 1 success")
	lock.Lock()
	fmt.Println("write lock 2 sucess")
	lock.RUnlock()
	lock.RUnlock()
}
