package main

import (
	"fmt"
)

func main() {
	for i := uint64(0); i < uint64(22); i++ { // calc for first 21 ints
		fmt.Printf("Factorial of %d is %d\n", i, Factorial(i))
	}
}

// named return variables:
func Factorial(n uint64) (fac uint64) {
	if n <= 1 { //base case
		return 1
	}
	fac = n * Factorial(n-1) // recursive case
	return
}
