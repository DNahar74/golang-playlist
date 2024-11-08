package main

import "fmt"

func main() {
	res1, _ := add(1, 2, 3, 4, 5)
	res2, message := add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println("Result 1:", res1)
	fmt.Println(message, "Result 2:", res2)
}

func add(val ...int) (int, string) {
	sum := 0
	for _, v := range val {
		sum += v
	}
	return sum, "HEHEHE"
}