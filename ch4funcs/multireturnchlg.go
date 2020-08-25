package main

import (
	"fmt"
)

func SumProductDiff(i, j int) (s int, p int, d int) {
	s, p, d = i+j, i*j, i-j
	return
}

func main() {
	sum, prod, diff := SumProductDiff(3, 4)
	fmt.Println("Sum:", sum, "| Product:", prod, "| Diff:", diff)
}
