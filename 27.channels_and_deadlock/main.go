//* An issue that comes with using weight groups and mutex is that they are just there to manage sync on a global scale
//* Weight group handles everything without letting goroutine know about other goroutines
//* Sometimes it is not possible to declare everything globally, like a local variable that only 2 goroutines need to access
//* In such cases, weight groups cannot do anything (We will need to declare the variable globally and make a seperate weight group maybe?)
//* Channels are a way for goroutines to communicate with each other
//* They are basically pipelines for communication, and can pass on any type of data

//* The arrows always point inwards (<-)

package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan int)

	//? This will cause deadlock
	// channel <- 5
	// fmt.Println(<-channel)

	//? This shows communication between main and goroutine
	// wg := &sync.WaitGroup{}

	// wg.Add(1)

	// go func(wg *sync.WaitGroup, channel chan int) {
	// 	channel <- 5
	// 	wg.Done()
	// }(wg, channel)

	// fmt.Println(<-channel)

	// wg.Wait()


	//? Sending messages from one channel to another
	// wg := &sync.WaitGroup{}

	// wg.Add(2)

	// go func(wg *sync.WaitGroup, channel chan int) {
	// 	fmt.Println("Start, Goroutine-1")
	// 	channel <- 5
	// 	fmt.Println("Sent channel, Goroutine-1")
	// 	wg.Done()
	// }(wg, channel)

	// go func(wg *sync.WaitGroup, channel chan int) {
	// 	fmt.Println("Start, Goroutine-2")
	// 	value := <-channel
	// 	fmt.Println("Recieved channel, Goroutine-2")
	// 	fmt.Println(value) 
	// 	wg.Done()
	// }(wg, channel)

	// wg.Wait()

	//? Using 2 channels, 3 goroutines, and one reciever in main
	//? One possible output is ::

	//* Start, Goroutine-3
	//* Start, Goroutine-2
	//* Start, Goroutine-1
	//* Sent channel, Goroutine-1
	//* 2
	//* Sent channel, Goroutine-3
	//* Recieved channel, Goroutine-2
	//* 5

	//? My understanding of the flow ::
	//? GR-3 starts, tries to send signal, reciever not ready, blocked
	//? GR-2 starts, tries to recieve value, no value sent, blocked
	//? GR-1 starts, sent value, printed, wg.Done()
	//? main recieved and printed
	//? GR-3, value sent successfully, wg.Done()
	//? GR-2, value recieved, printed, wg.Done()

	//! Issue :: Why Doesn't 2 Print Last?

	//* (1) Channels Are Blocking: The main function immediately reads from channel2 (fmt.Println(<-channel2)) as soon as "Goroutine-3" sends 2.
	//* (2) The main function's execution is not synchronized with the WaitGroup or the other goroutines—it processes channel2 independently.
	
	//? On adding close(channel) it allows the channel to close, and continue with the goroutine execution
	//? Also, the channel2 in main func is always ready to recieve
	wg := &sync.WaitGroup{}
	channel2 := make(chan int)

	wg.Add(3)

	go func(wg *sync.WaitGroup, channel chan int) {
		fmt.Println("Start, Goroutine-1")
		channel <- 5
		close(channel)
		fmt.Println("Sent channel, Goroutine-1")
		wg.Done()
	}(wg, channel)

	go func(wg *sync.WaitGroup, channel chan int) {
		fmt.Println("Start, Goroutine-2")
		value := <-channel
		fmt.Println("Recieved channel, Goroutine-2")
		fmt.Println(value) 
		wg.Done()
	}(wg, channel)

	go func(wg *sync.WaitGroup, channel2 chan int) {
		fmt.Println("Start, Goroutine-3")
		channel2 <- 2
		close(channel2)
		fmt.Println("Sent channel, Goroutine-3")
		wg.Done()
	}(wg, channel2)

	fmt.Println(<-channel2)

	wg.Wait()

	//? What if we want to make a buffered channel
	// channel := make(chan int, 10)
	// wg := &sync.WaitGroup{}
	// wg.Add(2)

	// go func(wg *sync.WaitGroup, channel chan int) {
	// 	for i := 0; i < 10; i++ {
	// 		channel <- i
	// 	}
	// 	// If you don't close then there is a deadlock
	// 	close(channel)
	// 	wg.Done()
	// }(wg, channel)

	// // go func(wg *sync.WaitGroup, channel chan int) {
	// // 	for val, ok := <-channel; ok == true; val, ok = <-channel {
	// // 		fmt.Println(val)
	// // 	}
	// // 	wg.Done()
	// // }(wg, channel)
	
	// //? Usecase of buffered channel
	// go func(wg *sync.WaitGroup, channel chan int) {
	// 	val := <-channel
	// 	// if ok == false {
	// 	// 	wg.Done()
	// 	// }
	// 	fmt.Println(val)
	// 	wg.Done()
	// }(wg, channel)

	// wg.Wait()

	//? Also, you can make channels Recieve ONLY or Send ONLY

	//? recieve Only channel
	// go func(wg *sync.WaitGroup, channel <-chan int) {
	// 	val := <-channel
	// 	// if ok == false {
	// 	// 	wg.Done()
	// 	// }
	// 	fmt.Println(val)
	// 	wg.Done()
	// }(wg, channel)

	//? send Only channel
	// go func(wg *sync.WaitGroup, channel chan<- int) {
	// 	for i := 0; i < 10; i++ {
	// 		channel <- i
	// 	}
	// 	// If you don't close then there is a deadlock
	// 	close(channel)
	// 	wg.Done()
	// }(wg, channel)
}
