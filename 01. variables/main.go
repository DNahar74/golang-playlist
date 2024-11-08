package main

import "fmt"

// This is not allowed
// JWT := "jwt"

// JWT is a json web token
var JWT string= "jwt"

func main() {
	var username string = "username"
	fmt.Println("username: ", username)
	fmt.Printf("type: %T \n", username)					//? %T shows the type

	var isLoggedIn bool = true
	fmt.Println("isLoggedIn: ", isLoggedIn)
	fmt.Printf("type: %T \n", isLoggedIn)
	
	var smallValue uint8 = 255
	fmt.Println("smallValue: ", smallValue)
	fmt.Printf("type: %T \n", smallValue)
	
	var dec32 float32 = 8764.98765432123456789
	fmt.Println("32-decimal: ", dec32)					//? This cuts off the decimal point after 3 points (will change on the basis of the integer part)
	fmt.Printf("type: %T \n", dec32)
	var dec64 float64 = 8764.98765432123456789
	fmt.Println("64-decimal: ", dec64)					//? This cuts off the decimal point after 12 points (will change on the basis of the integer part)
	fmt.Printf("type: %T \n", dec64)
	
	//? Default values
	//? All data types have an initial value in Go, unlike garbage values used in C
	var newVar int
	fmt.Println("newVar: ", newVar)
	fmt.Printf("type: %T \n", newVar)

	//? Implicit declarations
	var website = "shobu.io"
	fmt.Println("website:", website)

	//? Walrus operator (Remember, it can only be used inside function blocks-not globally)
	authors := 7000000
	fmt.Println("Walrus operator:", authors)

	//? Constants
	const pi = 3.141592653589793
  fmt.Println("pi:", pi)

  //? iota (used for incrementing a value inside a const declaration)
  const (
    zero = iota      													// 0
    one = iota       													// 1
    two = iota       													// 2
  )
  fmt.Println("iota:", zero, one, two)

  //? Type conversions
  var myInt int = 123
  var myFloat64 float64 = float64(myInt)
  fmt.Println("myFloat64:", myFloat64)
}