package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello to Pizza Palace!")
	fmt.Println("Enter your name: ")

	reader := bufio.NewReader(os.Stdin)

	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println(name, "please rate our pizza from 1 to 5")

	rating, _ := reader.ReadString('\n')
	numRating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err!= nil {
    fmt.Println("Invalid rating, please enter a number between 1 and 5.")
		fmt.Println(err)
    os.Exit(1)
  } else if numRating < 1 || numRating > 5 {
		fmt.Println("Invalid rating, please enter a number between 1 and 5.")
		os.Exit(1)
	} else {
		fmt.Printf("%s Thank you for your rating of %.2f.\n", name, numRating)					//? This will round it up to 2 digits after .
	}
}