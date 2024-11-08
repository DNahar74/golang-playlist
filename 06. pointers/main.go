package main

import "fmt"

func main() {
	sum := 50
	sumPtr := &sum
	a := 10
	b := 20
	*sumPtr = a + b
	fmt.Println("Sum:", sum)
}