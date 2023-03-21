package main

import (
	"fmt"
	"sync/atomic"
	"time"
	"unsafe"
)

func (m *Mutex) WaiterCount() int {
	v := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	v = v >> mutexWaiterShift // get number of goroutine waiting for lock
	v = v + (v & mutexLocked) // add the holder of lock (0 or 1)
	return int(v)
}

func (m *Mutex) IsLocked() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexLocked == mutexLocked
}

func (m *Mutex) IsWoken() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexWoken == mutexWoken
}

func (m *Mutex) IsStarving() bool {
	state := atomic.LoadInt32((*int32)(unsafe.Pointer(&m.Mutex)))
	return state&mutexStarving == mutexStarving
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
