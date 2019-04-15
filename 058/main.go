package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this should no print")
	case 2 == 4:
		fmt.Println("this should no print too")
	case 3 == 3:
		fmt.Println("prints")
	case 4 == 4:
		fmt.Println("also true, does it print?")
	}

	switch {
	case false:
		fmt.Println("this should no print")
	case 2 == 4:
		fmt.Println("this should no print too")
	case 3 == 3:
		fmt.Println("prints")
		fallthrough
	case 4 == 4:
		fmt.Println("also true, does it print?")
		fallthrough
	case 7 == 9:
		fmt.Println("not true 1")
		fallthrough
	case 11 == 14:
		fmt.Println("not true 2")
		fallthrough
	case 15 == 15:
		fmt.Println("true 15")
	default:
		fmt.Println("this is default")
	}

	switch {
	case false:
		fmt.Println("this should no print")
	case 2 == 4:
		fmt.Println("this should no print too")
	default:
		fmt.Println("this is default")
	}

	switch "Bond" {
	case "Moneypenny":
		fmt.Println("miss money")
	case "Bond":
		fmt.Println("bond james")
	case "Q":
		fmt.Println("this is q")
	default:
		fmt.Println("this is default")
	}

	n := "Bond"
	switch n {
	case "Moneypenny":
		fmt.Println("miss money")
	case "Bond":
		fmt.Println("bond james")
	case "Q":
		fmt.Println("this is q")
	default:
		fmt.Println("this is default")
	}

	switch n {
	case "Moneypenny", "Bond", "Do No":
		fmt.Println("miss money or bond or do no")
	case "M":
		fmt.Println("m")
	case "Q":
		fmt.Println("this is q")
	default:
		fmt.Println("this is default")
	}
}
