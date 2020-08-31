package main

import (
	"fmt"
	"sync"
	"time"
)

func HeavyFunction1(wg *sync.WaitGroup) {
	defer wg.Done()
	// Do a lot of stuff
	fmt.Println("heavy func 1 is starting")
	time.Sleep(3 * 1e9) // sleep for 2 seconds
	fmt.Println("heavy func 1 is done")
}

func HeavyFunction2(wg *sync.WaitGroup) {
	defer wg.Done()
	// Do a lot of stuff
	fmt.Println("heavy func 2 is starting")
	time.Sleep(2 * 1e9) // sleep for 2 seconds
	fmt.Println("heavy func 2 is done")
}

func main() {
	wg := new(sync.WaitGroup)
	wg.Add(2)
	fmt.Println("starting heavy func1")
	go HeavyFunction1(wg)
	fmt.Println("starting heavy func2")
	go HeavyFunction2(wg)

	wg.Wait()
	fmt.Printf("All Finished!")
}
