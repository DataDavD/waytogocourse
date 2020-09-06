package main

import "fmt"

// Send the sequence 2, 3, 4, ... to channel ch.
func generate(ch chan int) {
	for i := 2; ; i++ {
		ch <- i // Send i to channel ch.
	}
}

// Copy the values from channel in to channel out,
// removing those divisible by prime.
func filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.

		if i%prime != 0 {
			fmt.Printf("Prime filter (%d): passing %d\n", prime, i)
			out <- i // Send 'i' to 'out'.
		} else {
			fmt.Printf("Prime filter (%d): filtered %d\n", prime, i)
		}
	}
}

// The prime sieve
func main() {
	ch := make(chan int) // Create a new channel.
	go generate(ch)      // Start generate() as a goroutine.
	for i := 0; i < 3; i++ {
		prime := <-ch
		fmt.Println(prime)
		ch1 := make(chan int)
		go filter(ch, ch1, prime)
		ch = ch1
	}
}
