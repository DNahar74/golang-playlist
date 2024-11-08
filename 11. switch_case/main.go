package main

import (
	"fmt"
	"math/rand"
	// "time"
)

func main() {
	fmt.Println("Dice Game")

	// rand.New(rand.NewSource(time.Now().UnixNano()))
	dice := rand.Intn(6) + 1

	// no break statements required
	// if you want to make the next case run, put fallthrough

	switch dice {
	case 1:
		fmt.Println("You rolled a one!")
		fallthrough
	case 2:
		fmt.Println("You rolled a two!")
	case 3:
		fmt.Println("You rolled a three!")
	case 4:
		fmt.Println("You rolled a four!")
	case 5:
		fmt.Println("You rolled a five!")
	case 6:
		fmt.Println("You rolled a six!")
	default:
		fmt.Println("Invalid roll!")
	}
}