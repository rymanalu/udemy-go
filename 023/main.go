package main

import "fmt"

var y = 42
var z = "Shaken, no stirred"
var a = `She said, "Shaken, no stirred"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}
