package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)
	go sendData(ch) // calling goroutine to send the data
	go getData(ch)  // calling goroutine to receive the data
	time.Sleep(1e9)
	close(ch) // closed the channel
}

// sending data to ch channel
func sendData(ch chan string) {
	for _, v := range [5]string{"Washington", "Tripoli", "London", "Beijing", "Tokyo"} {
		fmt.Println("sending city")
		ch <- v
	}

}

func getData(ch chan string) {
	var input string
	for {
		input = <-ch // receiving data sent to ch channel
		fmt.Println("receiving city")
		fmt.Printf("%s ", input)
	}
}
