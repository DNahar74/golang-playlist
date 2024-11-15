package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"Course"`
	Price    int      `json:"Price"`
	Platform string   `json:"Website"`
	Password string   `json:"-"`
	Tags     []string `json:"Tags,omitempty"`
}

func main() {
	fmt.Println("Hello, world!")
	// encodeJSON()
	decodeJSON()
}

func encodeJSON() []byte {
	courses := []course{}
	courses = append(courses, course{Name: "Go Programming", Price: 100, Platform: "Udemy", Password: "1secret1", Tags: []string{"programming", "golang"}})
	courses = append(courses, course{Name: "Python Programming", Price: 150, Platform: "Coursera", Password: "2secret2", Tags: []string{"programming", "python"}})
	courses = append(courses, course{Name: "Java Programming", Price: 120, Platform: "EdX", Password: "3secret3", Tags: nil})

	//? Package this Data

	finalJSON, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling JSON:", err)
		panic(err)
	}

	// fmt.Println(string(finalJSON))
	return finalJSON
}

func decodeJSON() {
	dataJSON := encodeJSON()

	//? For normal usecases

	// var courses []course

	// checkValid := json.Valid(dataJSON)
	// if !checkValid {
	// 	fmt.Println("Invalid JSON")
	// 	return
	// }
	// fmt.Println("Is valid ::", checkValid)
	// json.Unmarshal(dataJSON, &courses)

	// fmt.Printf("%#v\n", courses)
	// for _, course := range courses {
	// 	fmt.Printf("%#v\n", course)
	// }

	//? For key value pairs

	var courseMap []map[string]interface{}
  json.Unmarshal(dataJSON, &courseMap)
  for _, body:= range courseMap {
    fmt.Printf("%+v\n", body)
  }
}
