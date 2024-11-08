// Whatever is in defer basically gets pushed into a stack that gets unloaded at the end of a function call

package main

import "fmt"

func add(val ...int) (int) {
	sum := 0
	for _, v := range val {
		sum += v
	}
	defer fmt.Println("Defer print in add, sum: ", sum)
	fmt.Println("Normal print in add")
	return sum
}

func main() {
	fmt.Println("Start")

	for i := 0; i < 3; i++ {
		fmt.Println("In the loop")
		defer fmt.Println("defer in the loop")
	}

	defer func() {
		fmt.Println("defer function called")
	}()

	add(12, 13, 14)

	defer fmt.Println("defer print")

	fmt.Println("End")
}