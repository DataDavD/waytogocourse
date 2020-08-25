package main

import "fmt"

var fib [10]int64 // global array containing Fibonacci values

func fibs() [10]int64 {
	fib[0] = 1 // base cases for 0
	fib[1] = 1 // base case for 1

	for i := 2; i < 10; i++ {
		fib[i] = fib[i-1] + fib[i-2] // recursive case without recursion using array
	}
	return fib
}

func main() {

	arr := fibs()
	for i := 0; i < 10; i++ {
		fmt.Printf("The %d-th Fibonacci number is: %d\n", i, arr[i])
	}
}
