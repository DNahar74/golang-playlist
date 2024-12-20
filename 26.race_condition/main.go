package main

import (
	"fmt"
	"sync"
)

//? Because these are three independent goroutines, there is no guarantee which will complete earlier i.e, you cannot determine the order of append
func main() {
	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{}
	score := []int{0}

	// Instead of adding one everytime, add 3 for the number of goroutines (No difference in performance)
	wg.Add(4)

	// This is basically like an arrow function or a lambda function
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		defer wg.Done()
		fmt.Println("Routine-1 ::")
		mut.Lock()
		fmt.Println("score ::", score)
		score = append(score, 1)
		mut.Unlock()
	}(wg, mut)

	// If we don't unlock mut, then no other function can read or write
	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		defer wg.Done()
		fmt.Println("Routine-2 ::")
		mut.Lock()
		fmt.Println("score ::", score)
		score = append(score, 2)
		mut.Unlock()
	}(wg, mut)

	go func(wg *sync.WaitGroup, mut *sync.RWMutex) {
		defer wg.Done()
		fmt.Println("Routine-3 ::")
		mut.Lock()
		fmt.Println("score ::", score)
		score = append(score, 3)
		mut.Unlock()
	}(wg, mut)

	// If you RLock, it means that resources trying to read can access it but no func can write
	go func(wg *sync.WaitGroup, mut *sync.RWMutex)  {
		defer wg.Done()
		fmt.Println("Routine-4 ::")
		mut.RLock()
		fmt.Println("score ::", score)
		mut.RUnlock()
	}(wg, mut)

	wg.Wait()
	fmt.Println("score ::", score)
}
