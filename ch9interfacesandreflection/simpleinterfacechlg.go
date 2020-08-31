package main

import "fmt"

// interface implementing functions called on Simple struct
type Simpler interface {
	Get() int
	Set(int)
}

type Simple struct {
	i int
}

func (p *Simple) Get() int {
	return p.i
}

func (p *Simple) Set(u int) {
	p.i = u
}

// function calling both methods through interface
func fI(it Simpler) int {
	it.Set(10)
	return it.Get()
}

func main() {
	var s Simple
	fmt.Println(fI(&s)) // &s is required because Get() is defined with a receiver type pointer
}
