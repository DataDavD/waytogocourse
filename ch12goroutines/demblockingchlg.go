package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	go func() {
		time.Sleep(5 * 1e9)
		x := <-c // receive value from channel c
		fmt.Println("received", x)
	}()
	fmt.Println("sending", 10)
	c <- 10 // sending 10 to the channel
	fmt.Println("sent", 10)
}
