package main

import "fmt"

func main() {
	printHello()
	greeting("Ringo")
	s := sprintHello("Paul")
	fmt.Println(s)

	x, y := strlenmerge("John", "George")
	fmt.Println(x)
	fmt.Println(y)
}

func printHello() {
	fmt.Println("Hello from printHello function")
}

func greeting(subject string) {
	fmt.Println("Hello,", subject)
}

func sprintHello(subject string) string {
	return fmt.Sprint("Hello from sprintHello, ", subject)
}

func strlenmerge(s1 string, s2 string) (string, int) {
	r1 := fmt.Sprint(s1, s2)
	r2 := len(s1) + len(s2)
	return r1, r2
}
