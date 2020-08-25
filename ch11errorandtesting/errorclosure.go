package main

import (
	"errors"
	"fmt"
)

type fType1 func(a, b int)

func errorHandler(fn fType1) fType1 {
	return func(a, b int) {
		defer func() {
			if err, ok := recover().(error); ok {
				fmt.Println("defer/recovered error: ", err)
			}
		}()
		fn(a, b)
	}
}

func main() {
	var try fType1 = func(a, b int) {
		if b == 1 {
			panic(errors.New("this is an error"))
		} else {
			fmt.Println(a + b)
		}
	}
	er := errorHandler(try)
	er(1, 1)
}
