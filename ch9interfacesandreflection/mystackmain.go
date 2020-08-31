package main

import (
	"fmt"

	"github.com/DataDavD/waytogocourse/ch9interfacesandreflection/mystack"
)

type Dd struct {
	name string
}

func main() {
	var sl mystack.Stack
	sl.Push("a")
	sl.Push("david")
	sl.Push(`juju`)
	sl.Push([]int{1, 2, 3})
	fmt.Println(sl.Len())
	sl.Push(Dd{"david"})
	fmt.Println(sl.Len())
	sl.Pop()
	fmt.Println(sl.Len())
	fmt.Println(sl)
	fmt.Println(sl.IsEmpty())
	sl.Pop()
	sl.Pop()
	sl.Pop()
	sl.Pop()
	fmt.Println(sl.IsEmpty())
	fmt.Println(sl.Len())
	fmt.Println(sl)
}
