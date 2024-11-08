package main

import (
	"fmt"
	"sort"
)

func main() {
	var s = []string{"hello", "world"}
	s = append(s, "wow!")

	fmt.Println(s)

	scores := make([]int, 4)
	scores[0] = 90
	scores[1] = 85
	scores[2] = 95
	scores[3] = 80

	scores = append(scores, 60, 70, 65, 75)

	fmt.Println(scores)
	fmt.Println(len(scores))

	scores[7] = 45
	fmt.Println(scores)

	// You can only add indexes using append
	// scores[8] = 100

	sort.Ints(scores)
	fmt.Println(scores)

	courses := []string{"java", "swift", "golang", "rust", "c#", "python"}
	fmt.Println(courses)

	indexToRemove := 1

	courses = append(courses[:indexToRemove], courses[(indexToRemove+1):]...)		// spread operator
	fmt.Println(courses)
}