package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")

	// days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

	// for i := 0; i < len(days); i++ {
	// 	fmt.Printf("%s%s\n", days[i], "day")
	// }

	// for i := range days {
	// 	fmt.Printf("%s\n", days[i])
	// }

	// for i, day := range days {
	// 	fmt.Printf("%d. %s\n", i+1, day)
	// }

	// for _, day := range days {
	// 	fmt.Printf("%s\n", day)
	// }

	val := 1

	for val < 10 {
		if val%2 == 0 {
			val++
			continue
		} else if val == 5 {
			goto jump
		}
		fmt.Println(val)
		val++
	}

	jump:
	fmt.Println("GOTO...", val)
}
