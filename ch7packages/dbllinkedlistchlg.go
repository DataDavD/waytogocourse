package main

import (
	"container/list"
	"fmt"
)

func insertListElements(n int) *list.List { // add elements in list from 1 to n
	lst := list.New()
	for i := 1; i <= n; i++ {
		lst.PushBack(i) // insertion here
	}
	return lst
}

func main() {
	n := 5                          // total number of elements to be inserted
	myList := insertListElements(n) // function call
	for e := myList.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value) // printing values of list
	}
}
