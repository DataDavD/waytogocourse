package main

import (
	"fmt"
	"strings"

	"github.com/DataDavD/waytogocourse/ch9interfacesandreflection/mysort2"
)

type Persons []Person

type Person struct {
	firstName string
	LastName  string
}

func (p Persons) Len() int { return len(p) }

func (p Persons) Less(i, j int) bool {
	var res bool
	if val := strings.Compare(p[i].LastName, p[j].LastName); val == 0 {
		res = p[i].firstName < p[j].firstName
	} else {
		res = p[i].LastName < p[j].LastName
	}
	return res
}

func (p Persons) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	p1 := Person{"david", "dansby"}
	p2 := Person{"juliette", "dansby"}
	p3 := Person{"ryan", "dansby"}
	p4 := Person{"meg", "thai"}
	var ps Persons = []Person{p2, p4, p3, p1}
	mysort.Sort(ps)
	if !mysort.IsSorted(ps) {
		panic("ps is not sorted, you fucked up")
	}
	for _, p := range ps {
		fmt.Println(p)
	}
}
