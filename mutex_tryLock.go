package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
	"unsafe"
)

// copy mutex's constant
const (
	mutexLocked      = 1 << iota // locked position
	mutexWoken                   // lock awake position
	mutexStarving                // lock starving position
	mutexWaiterShift = iota      // start's position of waiter
)

// Mutex : Expand mutex structure
type Mutex struct {
	sync.Mutex
}

// TryLock : try to get lock
func (m *Mutex) TryLock() bool {
	// if it can get the lock
	if atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), 0, mutexLocked) {
		return true
	}
	// if it's waken, locked or starving, this request is out of competition and return false
	old := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	if old&(mutexLocked|mutexStarving|mutexWoken) != 0 {
		return false
	}

	// try to request lock in competition
	new := old | mutexLocked
	return atomic.CompareAndSwapInt32((*int32)(unsafe.Pointer(&m.Mutex)), old, new)
}

func main() {
	try()
}

func try() {
	var mu Mutex
	go func() {
		mu.Lock()
		time.Sleep(time.Duration(rand.Intn(2)) * time.Second)
		mu.Unlock()
	}()
	time.Sleep(time.Second)

	ok := mu.TryLock()
	if ok {
		fmt.Println("got the lock")
		mu.Unlock()
		return
	}

	fmt.Println("can't get the lock")
}
