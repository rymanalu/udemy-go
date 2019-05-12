package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`

	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)

	if err != nil {
		fmt.Println("err:", err)
	}

	fmt.Println("password:", s)
	fmt.Println("encrypted password:", string(bs))

	loginPass := `password123`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPass))

	if err != nil {
		fmt.Println("wrong password")
	} else {
		fmt.Println("password matched")
	}
}
