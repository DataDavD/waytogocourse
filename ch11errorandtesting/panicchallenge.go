package main

import (
	"fmt"
)

func badCall() {
	a, b := 5, 0
	n := a / b // this will panic
	fmt.Println(n)
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Panicking with error: ", e)
		}

	}()
	badCall()
}

func main() {
	fmt.Println("before calling test func")
	test()
	fmt.Println("after calling test func")
}
