package main

import "fmt"

type user struct {
	username string
	email    string
	password string
	age      int
}

func main() {
	user1 := user{}

	user1.username = "john_doe"
	user1.email = "john.doe@example.com"
	user1.password = "password123"
	user1.age = 30

	fmt.Printf("User 1: %+v\n", user1)
}
