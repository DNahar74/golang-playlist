package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const url = "https://go.dev/solutions/use-cases"

func main() {
	response, err := http.Get(url)
	checkError(err)

	fmt.Printf("Response is of type :: %T\n", response)
	fmt.Println("response :: ", response)

	defer response.Body.Close()		// close the connection

	data, err := io.ReadAll(response.Body)
	checkError(err)

	writeToFile("./use_cases.html", string(data))

	h,_,_ := strings.Cut(string(data), "<body")
	head := strings.Split(h, "\n")
	fmt.Println("Data received from the server :: ")
	for _, line := range head {
    fmt.Println(line)
  }
}

func checkError(err error) error {
	if err != nil {
		panic(err)
	}
	return nil
}

func writeToFile(filePath string, content string)  {
	file, err := os.Create(filePath)
	checkError(err)
	defer file.Close()

	_, err = io.WriteString(file, content)
	checkError(err)
}