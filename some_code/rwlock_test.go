package some_code

import (
	"sync"
	"testing"
)

var lock sync.RWMutex

func TestRWLock(t *testing.T) {
	//TestLockUp(lock)
	//TestLockDown(lock)
	//TestReLock(lock)
	TestReRLock(lock)
}
