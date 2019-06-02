package main

import (
	"fmt"
)

type customErr struct {
	info string
}

func (err customErr) Error() string {
	return fmt.Sprintf("Here is the error info: %v", err.info)
}

func main() {
	foo(customErr{"Just info"})
}

func foo(err error) {
	fmt.Println("foo ran")

	fmt.Println("err:", err, "\n", err.(customErr).info)
}
