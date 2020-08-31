package main

import "fmt"

type Square struct {
	side float32
}

// implement this struct
type Triangle struct {
	base   float32
	height float32
}

type AreaInterface interface {
	Area() float32
}

// implement this interface
type PeriInterface interface {
	Perimeter() float32
}

func main() {
	var ai AreaInterface
	var pi PeriInterface
	trgl := &Triangle{5, 5}
	ai = trgl
	fmt.Println(ai.Area())

	sqr := &Square{5}
	ai = sqr
	fmt.Println(ai.Area())
	pi = sqr
	fmt.Println(pi.Perimeter())

}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

// implement method called on square to calculate its perimeter
func (sq *Square) Perimeter() float32 {
	return (sq.side * 4)
}

// implement method called on triangle to calculate its area
func (tr *Triangle) Area() float32 {
	return (0.5 * (tr.height * tr.base))
}
