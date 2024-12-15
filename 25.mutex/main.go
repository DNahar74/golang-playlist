package main

import (
	"fmt"
	"net/http"
	"sync"
)

//? An issue that comes up here is that when we are adding data to the signals slice, it is added based on the difference in time
//? But, we might want to store them in order, OR
//? Suppose we are making a database, and we get multiple goroutines that are writing to the same memory location.
//? This becomes an issue. So, to solve this locks were introduced.

var signals = []string{}
var weightGroup sync.WaitGroup // These are generally pointers
var mut sync.Mutex // These are generally pointers

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

	fmt.Println("signals :: ", signals)
}

func getStatusCode(endpoint string) {
	// This lets the weightGroup know that this goroutine has finished
	defer weightGroup.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Issue in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()
		fmt.Printf("%d status code for endpoint %s\n", res.StatusCode, endpoint)
	}
}
