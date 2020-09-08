package main

import (
	"fmt"
)

var resume chan int

func integers() chan int {
	yield := make(chan int)
	count := 0
	go func() {
		for {
			yield <- count
			count++
		}
	}()
	return yield
}

func generateInteger(rs <-chan int) int {
	return <-rs
}

func main() {
	resume = integers()
	fmt.Println(generateInteger(resume)) //=> 0
	fmt.Println(generateInteger(resume)) //=> 1
	fmt.Println(generateInteger(resume)) //=> 2
}
