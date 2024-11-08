// For main.exe files you have to run ./main.exe

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time!")

	present := time.Now()
	fmt.Println("Present time:", present.Format("02-01-2006 15:04:05 Monday"))				
	//? for these formattings, 02=date, 01=month, 2006=year, 15=hour, 04=minute, 05=second and Monday=day

	createdDate := time.Date(2005, time.August, 17, 12, 30, 0, 0, time.UTC)
	fmt.Println("Created date:", createdDate.Format("02-01-2006 15:04:05 Monday"))
}