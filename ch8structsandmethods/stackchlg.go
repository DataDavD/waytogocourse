package main

import (
	"fmt"
	"strconv"
)

const LIMIT = 4 // stack size limit

type Stack struct {
	ix   int // first free position, so data[ix] == 0
	data [LIMIT]int
}

func main() {
	st1 := new(Stack)
	fmt.Printf("%v\n", st1)
	st1.Push(3) // function call to Push
	fmt.Printf("%v\n", st1)
	st1.Push(7) // function call to Push
	fmt.Printf("%v\n", st1)
	st1.Push(10) // function call to Push
	fmt.Printf("%v\n", st1)
	st1.Push(99) // function call to Push
	fmt.Printf("%v\n", st1)
	p := st1.Pop() // function call to Pop
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
	p = st1.Pop() // function call to Pop
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
	p = st1.Pop() // function call to Pop
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
	p = st1.Pop() // function call to Pop
	fmt.Printf("Popped %d\n", p)
	fmt.Printf("%v\n", st1)
}

func (st *Stack) Push(n int) {
	if st.ix == LIMIT {
		return // stack is full!
	}
	st.data[st.ix] = n
	st.ix++
}

func (st *Stack) Pop() int {
	if st.ix > 0 {
		st.ix--
		element := st.data[st.ix]
		st.data[st.ix] = 0
		return element
	}
	return -1
}

func (st Stack) String() string {
	str := ""
	for ix := 0; ix < st.ix; ix++ {
		str += "[" + strconv.Itoa(ix) + ":" + strconv.Itoa(st.data[ix]) + "] "
	}
	return str
}
