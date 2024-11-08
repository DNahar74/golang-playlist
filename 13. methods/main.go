package main

import "fmt"

// User struct to represent a user
type User struct {
	Username string
	Email    string
	Password string
	Age      int
}

// setUsername method to update the username, only works if you pass *User instead of User struct
func (u *User) setUsername() string {
	u.Username = "WeTheBest"
	return u.Username
}

func main() {
	hash := User{}
	hash.Username = "john_doe"
	hash.Email = "john.doe@example.com"
	hash.Password = "password123"
	hash.Age = 30
	fmt.Println("Username:", hash.setUsername())
	fmt.Printf("User: %+v\n", hash.Username)
}
