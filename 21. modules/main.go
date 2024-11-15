//todo: (1) Check why it shows error when it is working

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Hello, world!")
	r := mux.NewRouter()
  r.HandleFunc("/", handleHomepage).Methods("GET")
	
	log.Fatal(http.ListenAndServe(":8080", r))				//? log basially logs out if there are any errors
}

func handleHomepage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>MY FIRST GO SERVER!!!</h1>"))
}