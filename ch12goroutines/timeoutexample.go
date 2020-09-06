package main

import (
	"fmt"
	"time"
)

func main() {
	timeout := make(chan bool, 1)
	go func() {
		time.Sleep(1e9) // one second
		timeout <- true
	}()

	select {
	case <-timeout:
		fmt.Println("received from ch")
	case <-timeout:
		fmt.Println("timeout")
		break
	}
}
