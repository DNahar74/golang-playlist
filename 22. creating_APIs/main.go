package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//? Models

// User represents a user in the system that can be an owner of multiple furnitures
type User struct {
	Username   string       `json:"username"`
	Email      string       `json:"email"`
	Password   string       `json:"-"`
	Age        int          `json:"age"`
	Furnitures []*Furniture `json:"furnitures"`
}

// Furniture represents a parent class for all furnitures
type Furniture struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
}

//? A fake Database

var users = []*User{
	{
		Username: "john_doe",
		Email:    "john.doe@example.com",
		Password: "password123",
		Age:      30,
		Furnitures: []*Furniture{
			{ID: 1, Type: "bed"},
			{ID: 2, Type: "table"},
		},
	},
	{
		Username: "jane_doe",
		Email:    "jane.doe@example.com",
		Password: "password456",
		Age:      28,
		Furnitures: []*Furniture{
			{ID: 1, Type: "couch"},
		},
	},
	{
		Username: "bill_smith",
		Email:    "bill.smith@example.com",
		Password: "password789",
		Age:      35,
		Furnitures: []*Furniture{
			{ID: 1, Type: "chair"},
			{ID: 2, Type: "table"},
			{ID: 3, Type: "lamp"},
		},
	},
	{
		Username: "sarah_williams",
		Email:    "sarah.williams@example.com",
		Password: "password101",
		Age:      29,
		Furnitures: []*Furniture{
			{ID: 1, Type: "couch"},
			{ID: 2, Type: "table"},
			{ID: 3, Type: "lamp"},
			{ID: 4, Type: "sofa"},
		},
	},
	{
		Username: "emily_jones",
		Email:    "emily.jones@example.com",
		Password: "password111",
		Age:      32,
		Furnitures: []*Furniture{
			{ID: 1, Type: "couch"},
			{ID: 2, Type: "table"},
			{ID: 3, Type: "lamp"},
			{ID: 4, Type: "sofa"},
			{ID: 5, Type: "bed"},
		},
	},
}

//? Helpers

// IsEmpty checks if the user object is empty
func (user *User) IsEmpty() bool {
	if user.Username == "" && user.Email == "" {
		return true
	}
	return false
}

// IsValid checks if a user object is valid
func (user *User) IsValid() bool {
	if user.IsEmpty() || len(user.Password) < 8 || user.Age >= 18 {
		return false
	}
	return true
}

//? Request Handlers

// HomepageHandler handles the home page
func HomepageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("<h1>Yay! My first API Homepage</h1>"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func getUsers(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func getUserByIdentifier(w http.ResponseWriter, r *http.Request) *User {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, user := range users {
		if user.Email == params["id"] || user.Username == params["id"] {
			return user
		}
	}
	return nil
}

func main() {
	log.Fatal(http.ListenAndServe(":8080", nil))
	http.HandleFunc("/", HomepageHandler)

	// mux:=http.NewServeMux()
	// mux.HandleFunc("/", HomepageHandler)
	// mux.HandleFunc("/users", getUsers)

	// log.Fatal(http.ListenAndServe(":8080", mux))
}
