package main

import "fmt"

func sum(x, y int, c chan int) {
	fmt.Printf("summing %d and %d and placing into %T\n", x, y, c)
	c <- x + y
	fmt.Println("placed into channel successfully")

}

func main() {
	x, y := 1, 2
	c := make(chan int)
	go sum(x, y, c)
	fmt.Println(<-c)
	fmt.Println("done")
}
