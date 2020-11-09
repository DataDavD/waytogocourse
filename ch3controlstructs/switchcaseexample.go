package main

import "fmt"

func main() {
	var num1 int = 99
	// Adding switch on num1
	switch num1 {
	case 98, 99:
		fallthrough // first case: num1 = 98 or 99
		//fmt.Println("It's equal to 98")
	case 100:
		fallthrough // second case: num1 = 100
		//fmt.Println("It's equal to 100")
	case 101:
		fallthrough
		//fmt.Println("its 101")
	case 102:
		fallthrough
		//fmt.Println("its 102")
	default: // optional/ default case
		fmt.Println("It's not equal to 98 or 100")
	}
}
