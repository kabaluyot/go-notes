package main

import (
	"fmt"
	"sync"
)

var numbers = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

// Objectives:
// Write some possible ways to print the int array numbers
func main() {}

// directGoRoutines randomly prints value without synchorinzation
func directGoRoutines() {
	for _, number := range numbers {
		go func(number int) {
			fmt.Println("Counter:", number)
		}(number)
	}
}

// simpleWaitGroup makes use of sync package (synchronization)
// This function adds a worker fpr every iteration then waits until all workers are finished.
func simpleWaitGroup() {
	var wg = sync.WaitGroup{}

	for _, number := range numbers {
		wg.Add(1)

		go func(number int) {
			defer wg.Done()

			fmt.Println("Counter:", number)
		}(number)
	}

	wg.Wait()
}

// waitGroupAndRWMutex use mutex exclusion to ensure 1 goroutine is reading the state at a time.
// With this synchornization, the output is expected ti be printed in order
func waitGroupAndRWMutex() {
	var wg = sync.WaitGroup{}
	var rw = sync.RWMutex{}

	for _, number := range numbers {
		wg.Add(1)
		rw.Lock()
		go func(number int) {
			defer wg.Done()
			fmt.Println("Counter:", number)
			rw.Unlock()
		}(number)
	}

	wg.Wait()
}
