package main

import (
	"fmt"
)

func main() {
	sla := []int{2, 6, 4, -10, 8, 89, 12, 68, -45, 37}
	fmt.Println("before sort: ", sla)
	// sla is passed via call by value, but since sla is a reference type
	// the underlying array is changed (sorted)
	bubbleSort(sla)
	fmt.Println("Solution implementation after sort:  ", sla)
	bubbleSort2(sla)
	fmt.Println("Personal dev implementation after sort:  ", sla)
}

func bubbleSort(sl []int) {
	// passes through the slice:
	for pass := 1; pass < len(sl); pass++ {
		// one pass:
		for i := 0; i < len(sl)-pass; i++ { // the bigger value 'bubbles up' to the last position
			if sl[i] > sl[i+1] {
				sl[i], sl[i+1] = sl[i+1], sl[i]
			}
		}
	}
}

func bubbleSort2(sl []int) []int {
	var N int = len(sl)
	var i int
	for i = 0; i < N; i++ {
		sweep(sl)
	}
	return sl
}

func sweep(numbers []int) {
	var N int = len(numbers)
	var firstIndex int = 0
	var secondIndex int = 1

	for secondIndex < N {
		var firstNumber int = numbers[firstIndex]
		var secondNumber int = numbers[secondIndex]

		if firstNumber > secondNumber {
			numbers[firstIndex] = secondNumber
			numbers[secondIndex] = firstNumber
		}

		firstIndex++
		secondIndex++
	}
}
