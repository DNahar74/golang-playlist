//todo: (1) Check the runtime package

package main

import (
	"fmt"
	"runtime"
)

func main() {
	num := runtime.NumCPU()
	fmt.Println(num)
}