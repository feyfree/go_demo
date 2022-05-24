package ch09

import (
	"sync"
	"testing"
)

// 测试可冲入锁 会报错
// === RUN   TestReentrant
// fatal error: all goroutines are asleep - deadlock!
func TestReentrant(t *testing.T) {
	var mutex sync.Mutex
	mutex.Lock()
	mutex.Lock()
}
