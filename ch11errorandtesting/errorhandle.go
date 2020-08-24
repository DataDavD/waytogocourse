package main

import "fmt"

type MyError struct{}

func (myerror *MyError) Error() string {
	return "Something happened unexpectedly!!!"
}

func main() {

	// create error
	myErr := new(MyError)

	// print error
	fmt.Println(myErr)
}
