package main

import (
	"fmt"
	"net/url"
)

const myURL = "http://darsh.example.com:8080/resume?languages=english,hindi,spanish&name=Darsh"

func main() {
	fmt.Println(myURL)

	//? Parsing the URL

	result, _ := url.Parse(myURL)
	fmt.Println("RESULT :: ", result)
	fmt.Println("SCHEME :: ", result.Scheme)
	fmt.Println("HOST :: ", result.Host)
	fmt.Println("HOSTNAME :: ", result.Hostname())
	fmt.Println("PORT :: ", result.Port())
	fmt.Println("PATH :: ", result.Path)
	// fmt.Println("RAW QUERY :: ", result.RawQuery)
	fmt.Println("QUERY PARAMETERS :: ", result.Query())

	for i, v := range result.Query() {
		fmt.Println(i, "::", v)
	}

	//? Making a URL

	parts := &url.URL{
		Scheme: "https",
		Host: "WhySoSerious.dev:3000",
		Path: "/History/Origin",
		RawQuery: "user=Darsh&city=Gotham",
	}

	madeURL := parts.String()
	fmt.Println("MADE URL :: ", madeURL)
}
