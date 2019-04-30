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

	people := map[string]person{
		p1.lastName: p1,
		p2.lastName: p2,
	}

	fmt.Println(people)

	for k, v := range people {
		fmt.Printf("%v:\n", k)
		fmt.Printf("\tFirstname:\t\t\t%v\n", v.firstName)
		fmt.Printf("\tLastname:\t\t\t%v\n", v.lastName)
		fmt.Printf("\tFavorite Ice Cream Flavors:\n")
		for _, iceCream := range v.favIceCreamFlavors {
			fmt.Printf("\t\t\t\t\t%v\n", iceCream)
		}
	}
}
