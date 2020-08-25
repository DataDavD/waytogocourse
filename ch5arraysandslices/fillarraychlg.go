package main

import "fmt"

func main() {
	var arr [15]int
	for i := 0; i < 15; i++ {
		arr[i] = i
	}
	fmt.Println(arr)
}
