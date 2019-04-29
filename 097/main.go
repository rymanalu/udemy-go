package main

import "fmt"

type person struct {
	firstName          string
	lastName           string
	favIceCreamFlavors []string
}

func main() {
	p1 := person{
		firstName:          "George",
		lastName:           "Harrison",
		favIceCreamFlavors: []string{"Almond", "Banana"},
	}

	p2 := person{
		firstName:          "Ringo",
		lastName:           "Starr",
		favIceCreamFlavors: []string{"Cherry", "Raspberry", "Caramel"},
	}

	fmt.Println("p1:")
	fmt.Printf("\tFirstname:\t\t\t%v\n", p1.firstName)
	fmt.Printf("\tLastname:\t\t\t%v\n", p1.lastName)
	fmt.Printf("\tFavorite Ice Cream Flavors:\n")
	for _, v := range p1.favIceCreamFlavors {
		fmt.Printf("\t\t\t\t\t%v\n", v)
	}

	fmt.Println("p2:")
	fmt.Printf("\tFirstname:\t\t\t%v\n", p2.firstName)
	fmt.Printf("\tLastname:\t\t\t%v\n", p2.lastName)
	fmt.Printf("\tFavorite Ice Cream Flavors:\n")
	for _, v := range p2.favIceCreamFlavors {
		fmt.Printf("\t\t\t\t\t%v\n", v)
	}
}
