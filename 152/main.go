package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// Race (Atomic)
func main() {
	var wg sync.WaitGroup

	var counter int32

	for i := 0; i < 69; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			runtime.Gosched()
			atomic.AddInt32(&counter, 1)
		}()
	}

	wg.Wait()
	fmt.Println("counter:", atomic.LoadInt32(&counter))
}
