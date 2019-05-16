package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	// Race condition
	fmt.Println("CPUs:", runtime.NumCPU())
	fmt.Println("Goroutines:", runtime.NumGoroutine())

	counter := 0
	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	var mtx sync.Mutex

	for i := 0; i < gs; i++ {
		go func() {
			mtx.Lock()
			v := counter
			runtime.Gosched()
			v++
			counter = v
			mtx.Unlock()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println("count:", counter)
}
