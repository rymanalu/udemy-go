package main

import "fmt"

var y = 420

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Printf("%b\n", y)
	fmt.Printf("%x\n", y)
	fmt.Printf("%#x\n", y)
	y = 69
	s := fmt.Sprintf("%#x\t%b\t%x", y, y, y)
	fmt.Println(s)
}
