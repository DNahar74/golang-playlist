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
	Password   string       `json:"password"`
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
	if user.IsEmpty() || len(user.Password) < 8 || user.Age < 18 {
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

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// func getUserByIdentifier(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)

// 	json.NewEncoder(w).Encode(params)

// 	for _, user := range users {
// 		if user.Email == params["id"] || user.Username == params["id"] {
// 			json.NewEncoder(w).Encode(user)
// 		}
// 	}
// }

func getUserByIdentifier(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	identifier := params["id"]

	for _, user := range users {
		if user.Email == identifier || user.Username == identifier {
			json.NewEncoder(w).Encode(user)
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

// func appendUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	// params := mux.Vars(r)

// 	var newUser User
// 	err := json.NewDecoder(r.Body).Decode(&newUser)
// 	if err!= nil {
//     http.Error(w, err.Error(), http.StatusBadRequest)
//     return
//   }

// 	if !newUser.IsValid() {
//     http.Error(w, "Invalid user data", http.StatusBadRequest)
//     return
//   }

// 	if newUser.IsEmpty() {
//     http.Error(w, "Empty user data", http.StatusBadRequest)
// 		return
// 	}

// 	users = append(users, &newUser)
// }

func appendUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// The issue is that if we put password = `json: '-'`, it ignores the field during encoding as well as decoding
	if !newUser.IsValid() {
		http.Error(w, "Invalid user data", http.StatusBadRequest)
		return
	}

	users = append(users, &newUser)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newUser)
}

// func deleteUser(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")

// 	params := mux.Vars(r)

// 	for index, user := range users {
// 		if user.Email == params["id"] || user.Username == params["id"] {
// 			users = append(users[:index], users[index+1:]...)
// 			return
// 		}
// 	}
// }

func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	identifier := params["id"]

	for index, user := range users {
		if user.Email == identifier || user.Username == identifier {
			users = append(users[:index], users[index+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"})
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

func main() {
	// log.Fatal(http.ListenAndServe(":8080", nil))
	// http.HandleFunc("/", HomepageHandler)

	// mux := http.NewServeMux() // this does not support dynamic routes
	// mux.HandleFunc("/", HomepageHandler)
	// mux.HandleFunc("/users", getUsers)
	// mux.HandleFunc("/users/{id}", getUserByIdentifier)

	// log.Fatal(http.ListenAndServe(":8080", mux))

	router := mux.NewRouter()

	router.HandleFunc("/", HomepageHandler)
	router.HandleFunc("/users", getUsers).Methods("GET")
	router.HandleFunc("/users/{id}", getUserByIdentifier).Methods("GET")
	router.HandleFunc("/users", appendUser).Methods("POST")
	router.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}
