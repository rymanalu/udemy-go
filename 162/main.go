package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go populate(c1)

	go fatOutIn(c1, c2)

	for v := range c2 {
		fmt.Println(v)
	}

	fmt.Println("exit")
}

func populate(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}

	close(c)
}

func fatOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup

	for v := range c1 {
		wg.Add(1)

		go func(x int) {
			c2 <- timeConsumingWork(x)
			wg.Done()
		}(v)
	}

	wg.Wait()
	close(c2)
}

func timeConsumingWork(x int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))

	return x + rand.Intn(1000)
}
