package main

import (
	"fmt"

	"github.com/rymanalu/udemy-go/196/quote"
	"github.com/rymanalu/udemy-go/196/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))

	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}
}
