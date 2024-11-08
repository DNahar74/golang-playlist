package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("File Handling")
	content := "This is the content to be written to the file, part 3\nHello World\nWe the best"

	writeFile("./myfile.txt", content)
	r := readFile("./myfile.txt")
	// fmt.Println(byte('\n'))
	// fmt.Println(string(65))
	fmt.Println(r)
}

func writeFile(fileName string, content string) int {
	file, err := os.Create(fileName)
	checkNilError(err)

	length, err := io.WriteString(file, content)
	checkNilError(err)

	defer file.Close()
	return length
}

func readFile(fileName string) []byte {
	content, err := os.ReadFile(fileName)
	checkNilError(err)
	for _, v := range strings.Split(string(content), "\n") {
		fmt.Println(v)
	}
	return content
}

func checkNilError(err error) error {
	if err!= nil {
		// fmt.Println("Error creating file:", err)
		// os.Exit(1)
    panic(err)
  }
  return nil  // return nil if no error
}