package main

import (
	"fmt"

	"github.com/DataDavD/waytogocourse/ch9interfacesandreflection/min"
)

func ints() {
	data := []int{1, 2, 3, 4, 60, 99, -100}
	a := min.IntArray(data)
	m := min.Min(a)
	fmt.Println(m)
}

func strings() {
	data := []string{"david", "juju", "meg", "longassname"}
	a := min.StringArray(data)
	m := min.Min(a)
	fmt.Println(m)
}

func main() {
	ints()
	strings()
}
