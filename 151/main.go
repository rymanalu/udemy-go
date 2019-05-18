package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Race (Mutex)
func main() {
	var wg sync.WaitGroup

	var mtx sync.Mutex

	counter := 0

	for i := 0; i < 69; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			mtx.Lock()
			current := counter
			runtime.Gosched()
			current++
			counter = current
			mtx.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println("counter:", counter)
}
