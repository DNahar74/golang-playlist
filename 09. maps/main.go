package main

import "fmt"

func main() {
	lang := make(map[string]int)

	lang["Go"] = 2012
	lang["Java"] = 1995
	lang["C++"] = 1985
	lang["Python"] = 1991
	lang["Ruby"] = 1995
	lang["JavaScript"] = 1995
	lang["Swift"] = 2014
	lang["Kotlin"] = 2011
	lang["Rust"] = 2010

	// fmt.Println(lang)
	// fmt.Println("Go was introduced in:", lang["Go"])

	// delete(lang, "Go")
	// fmt.Println("\nAfter deleting Go, the map is:")

	// loops

	for key, value := range lang {
		// fmt.Println(key, value)
		fmt.Printf("%v was introduced in: %v\n", key, value)
	}

	for _, date := range lang {
		fmt.Printf("The years are: %d\n", date)
	}
}