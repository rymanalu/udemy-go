package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p := person{
		First: "John",
		Last:  "Lennon",
		Sayings: []string{
			"I'm an Elvis fan because it was Elvis who really got me out of Liverpool.",
			"You're born in pain and pain is what we're in most of the time. And I think that the bigger the pain, the more gods we need.",
			"Paul & I enjoyed writing the music for the film, but there were times when we honestly thought we'd never get time to write all the material. We managed to get a couple finished while we were in Paris, and 3 more completed soaking up sun on Miami",
		},
	}

	bs, err := json.Marshal(p)

	if err != nil {
		log.Fatalln(err)

		return
	}

	fmt.Println(string(bs))
}
