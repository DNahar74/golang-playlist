package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	// "math/rand"
)

func main() {
	//? Typecasting shuld generally be to higher precision
	// var num1 int = 2
	// var num2 float64 = 4

	// fmt.Println("Sum is ::", num1+int(num2))
	// fmt.Println("Sum is ::", float64(num1) + num2)

	//? Random number from math/rand
	// fmt.Println("Random number ::", rand.Int())
	// rand.Seed(23)	// No need for this now for go versions

	// fmt.Println("Random number ::", rand.Intn(5))

	//? Random number from math/crypto (more reliable)
	num, err := rand.Int(rand.Reader, big.NewInt(200))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Random number ::", num)
}
