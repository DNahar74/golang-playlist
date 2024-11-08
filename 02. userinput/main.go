//todo: (1) Read up on bufio at pkg.go.dev
//todo: (2) Read up on bufio at pkg.go.dev

//? There are 2 packages that are used (1) bufio and (2) os

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	hello := "hello, this is the user input section"
	fmt.Println(hello)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter your name: ")
	
	//? This is the ,ok syntax. 
	//? If everything is good, then the input is stored in name otherwise it is stored in err

	// name, err := reader.ReadString('\n')

	//? This is the shorter version of the above syntax. But, here we do not care about err so we replace it with _
	name, _ := reader.ReadString('\n')
	fmt.Printf("Hello, %s\n", name)
}