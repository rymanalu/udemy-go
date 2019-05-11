package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func main() {
	response := `[{"first_name":"John","last_name":"Lennon","age":40},{"first_name":"Paul","last_name":"McCartney","age":76}]`
	bs := []byte(response)

	var people []person

	err := json.Unmarshal(bs, &people)

	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println(people)

	for _, v := range people {
		fmt.Println("Name:", v.FirstName, v.LastName)
		fmt.Println("Age:", v.Age)
	}
}
