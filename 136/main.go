package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string `json:"first"`
	Age   int    `json:"age"`
}

func main() {
	u1 := user{"John", 40}
	u2 := user{"Paul", 77}
	u3 := user{"George", 58}
	u4 := user{"Ringo", 80}

	users := []user{u1, u2, u3, u4}

	fmt.Println(users)

	bs, err := json.Marshal(users)

	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println(string(bs))
}
