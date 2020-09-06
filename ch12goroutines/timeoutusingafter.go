package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)
	go func() {
		time.Sleep(1e10)
		ch <- "success returned"
	}()
	select {
	case resp := <-ch:
		fmt.Println(resp)
	case <-time.After(1e9):
		fmt.Println("timeout")
		break
	}
}
