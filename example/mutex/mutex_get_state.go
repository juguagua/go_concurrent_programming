package mutex

import (
	"fmt"
	"sync/atomic"
	"time"
	"unsafe"

	"go_concurrent_programming/example"
)

func (m *example.Mutex) WaiterCount() int {
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	v = v >> example.mutexWaiterShift // get number of goroutine waiting for lock
	v = v + (v & example.mutexLocked) // add the holder of lock (0 or 1)
	return int(v)
}

func (m *example.Mutex) IsLocked() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&example.mutexLocked == example.mutexLocked
}

func (m *example.Mutex) IsWoken() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&example.mutexWoken == example.mutexWoken
}

func (m *example.Mutex) IsStarving() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&example.mutexStarving == example.mutexStarving
}

func GetState() {
	var mu Mutex
	for i := 0; i < 1000; i++ {
		go func() {
			mu.Lock()
			time.Sleep(time.Second)
			mu.Unlock()
		}()
	}
	time.Sleep(time.Second)
	fmt.Printf("waitings : %d, isLocked: %t, isWoken: %t, isStarving: %t",
		mu.WaiterCount(), mu.IsLocked(), mu.IsWoken(), mu.IsStarving())
}
