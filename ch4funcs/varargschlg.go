package main

import "fmt"

func main() {
	fmt.Println(sumInts())
	fmt.Println(sumInts(2, 3))
	fmt.Println(sumInts(0, 1, 2, 3, 4, 5, 6, 7, 8, 9))
}

func sumInts(list ...int) (sum int) {
	for _, v := range list {
		sum = sum + v
	}
	return
}
