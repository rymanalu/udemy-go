package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	fmt.Println("error check 1:", ctx.Err())
	fmt.Println("num of goroutines 1:", runtime.NumGoroutine())

	go func() {
		n := 0

		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("working", n)
			}
		}
	}()

	time.Sleep(time.Second * 2)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("num of goroutines 2:", runtime.NumGoroutine())

	fmt.Println("about to cancel")
	cancel()
	fmt.Println("context cancelled")

	time.Sleep(time.Second * 2)
	fmt.Println("error check 2:", ctx.Err())
	fmt.Println("num of goroutines 2:", runtime.NumGoroutine())
}
