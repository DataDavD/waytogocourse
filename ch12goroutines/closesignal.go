package main

import "fmt"

func main() {
	ch := make(chan string)
	go sendData(ch)
	getData(ch)
}

func sendData(ch chan<- string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
	close(ch)
}

// getData gets data from channel
func getData(ch <-chan string) {
	for {
		input, open := <-ch
		if !open {
			fmt.Println("breaking out of loop")
			break
		}
		fmt.Printf("%s \n", input)
	}
	fmt.Println("out of loop")
}
