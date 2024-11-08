//todo: (1) len vs cap

package main

import (
	"fmt"
)

func main() {
	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[2] = "grape"
	fmt.Println(fruitList)
	fmt.Println(fruitList[2])
	fmt.Println("Capacity of the array:", cap(fruitList))
	fmt.Println("Length of the array:", len(fruitList))
}