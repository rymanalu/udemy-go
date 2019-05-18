package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Race
func main() {
	var wg sync.WaitGroup

	counter := 0

	for i := 0; i < 69; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			current := counter
			runtime.Gosched()
			current++
			counter = current
		}()
	}

	wg.Wait()
	fmt.Println("counter:", counter)
}
