package main

import (
	"fmt"
	"net/http"
	"sync"
)

//? This only prints the world part, because the goroutine started but the main function exited before it could be completed
// func main() {
// 	go greeter("hello")
// 	greeter("world")
// }

// func greeter(word string) {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println(word, i)
// 	}
// }

//? One way of overcoming this is by making the threads sleep for a certain time, letting both functions execute
// func main() {
// 	go greeter("hello")
// 	greeter("world")
// }

// func greeter(word string) {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(500 * time.Millisecond)
// 		fmt.Println(word, i)
// 	}
// }

//? In a wait group, if you say that you need to wait for 5 executions to return, It will not let the function end before they return
var weightGroup sync.WaitGroup // These are generally pointers

func main() {
	websiteList := []string{
		"https://lco.dev",
		"https://go.dev",
		"https://google.com",
		"https://github.com",
		"https://fb.com",
		"https://socialwinterofcode.com",
		"https://youtube.com",
	}

	for _, site := range websiteList {
		go getStatusCode(site)
		weightGroup.Add(1)
	}

	// This helps to keep the main function from ending until weightGroup = 0
	weightGroup.Wait()
}

func getStatusCode(endpoint string) {
	// This lets the weightGroup know that this goroutine has finished
	defer weightGroup.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Issue in endpoint")
	} else {
		fmt.Printf("%d status code for endpoint %s\n", res.StatusCode, endpoint)
	}
}
