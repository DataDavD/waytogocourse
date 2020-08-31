package main

import "fmt"

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

type RSimple struct {
	i int
}

func (p *RSimple) Get() int {
	return p.i
}

func (p *RSimple) Set(u int) {
	p.i = u

}

func fI(it Simpler) int {
	switch it.(type) {
	case *Simple:
		it.Set(10)
		return it.Get()
	case *RSimple:
		it.Set(100)
		return it.Get()
	default:
		return -1
	}
	return -1
}

func main() {
	s := new(Simple)
	rs := new(RSimple)
	fmt.Println(fI(s))
	fmt.Println(fI(rs))

}
