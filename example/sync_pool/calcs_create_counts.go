package sync_pool

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var numCalcsCreated int32

func createBuffer() any {
	// watch out concurrent issue
	atomic.AddInt32(&numCalcsCreated, 1)
	buffer := make([]byte, 1024)
	return &buffer
}

func ExecWithPool() {
	bufferPool := &sync.Pool{
		New: createBuffer,
	}

	numWorkers := 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			// apply for a buffer instance
			buffer := bufferPool.Get()
			_ = buffer.(*[]byte)
			defer bufferPool.Put(buffer)
		}()
	}
	wg.Wait()
	fmt.Printf("WithPool: %d buffer instance were created. \n", numCalcsCreated)

}

func ExecWithoutPool() {
	numWorkers := 1024 * 1024
	var wg sync.WaitGroup
	wg.Add(numWorkers)

	for i := 0; i < numWorkers; i++ {
		go func() {
			defer wg.Done()
			// apply for a buffer instance
			_ = createBuffer()
			//_ = buffer.(*[]byte)
		}()
	}
	wg.Wait()
	fmt.Printf("WithoutPool: %d buffer instance were created. \n", numCalcsCreated)
}
