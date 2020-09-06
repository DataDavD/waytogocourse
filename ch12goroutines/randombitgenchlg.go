package main

import "fmt"

func main() {
	ch := make(chan int)
	// consumer:
	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()
	// producer:
	for i := 0; i < 100000; i++ {
		select {
		case ch <- 0:
		case ch <- 1:
		}
	}
}
