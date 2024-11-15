//todo: (1) Explore string buffer in strings package

package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const myURL = "http://localhost:3000"

func main() {
	getRequest(myURL)
	postRequestJSON(myURL)
	postRequestEncodedForm(myURL)
}

func getRequest(URL string) {
	response, err := http.Get(URL)
	checkError(err)
	defer response.Body.Close()

	fmt.Println("Status code ::", response.StatusCode)
	fmt.Println("Content length ::", response.ContentLength)

	var responseString strings.Builder

	content, err := io.ReadAll(response.Body)
	checkError(err)
	// fmt.Println("Content ::", string(content))

	byteCount, _ := responseString.Write(content)
	fmt.Println("Byte count ::", byteCount)

	fmt.Println("Response string ::", responseString.String())
}

func postRequestJSON(URL string) {
	//? Making a fake JSON payload
	requestBody := strings.NewReader(`
		{
			"name": "John Doe",
			"age": 30,
			"city": "New York"
		}
	`)
	response, err := http.Post(URL, "application/json", requestBody)
	checkError(err)
	defer response.Body.Close()

	var responseString strings.Builder

	content, err := io.ReadAll(response.Body)
	checkError(err)

	responseString.Write(content)
	fmt.Println("Response string ::", responseString.String())
}

func postRequestEncodedForm(URL string) {
	//? Form data

	data := url.Values{}

	data.Add("name", "John Doe")
	data.Add("email", "john@example.com")

	response, err := http.PostForm(URL, data)
	checkError(err)
	defer response.Body.Close()

	var responseString strings.Builder

	content, err := io.ReadAll(response.Body)
	checkError(err)

	responseString.Write(content)
	fmt.Println("Response string ::", responseString.String())
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
