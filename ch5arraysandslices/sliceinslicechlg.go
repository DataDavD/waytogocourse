package main

import (
	"fmt"
)

func main() {
	s := []string{"M", "N", "O", "P", "Q", "R"}
	in := []string{"A", "B", "C"}
	res := insertSlice(s, in, 0) // at the front
	fmt.Println(res)             // [A B C M N O P Q R]
	res = insertSlice(s, in, 3)  // [M N O A B C P Q R]
	fmt.Println(res)
}

func insertSlice(slice, insertion []string, index int) []string {
	result := make([]string, len(slice)+len(insertion))
	at := copy(result, slice[:index])
	at += copy(result[at:], insertion)
	copy(result[at:], slice[index:])
	return result
}
