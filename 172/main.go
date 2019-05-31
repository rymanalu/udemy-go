package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	// 1
	// n, err := fmt.Println("Hello")

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(n)

	// 2
	// var name string
	// var age int

	// fmt.Print("Name: ")
	// _, err := fmt.Scan(&name)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Print("Age: ")
	// _, err = fmt.Scan(&age)

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println(name, age)

	// 3
	f, err := os.Create("greatest-band.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := strings.NewReader("The Beatles")

	io.Copy(f, r)
}
