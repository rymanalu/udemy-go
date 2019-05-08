package main

import "fmt"

func main() {
	func() {
		fmt.Println("called from anonymous func")
	}()
}
