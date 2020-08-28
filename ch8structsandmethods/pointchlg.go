package main

import (
	"fmt"
	"math"
)

type Point struct { /// struct of type Point
	X, Y float64
}

func (p *Point) Abs() float64 { // method calculating absolute value
	return math.Sqrt(float64(p.X*p.X + p.Y*p.Y))
}

func (p *Point) Scale(s float64) { // method to scale a point
	p.X = p.X * s
	p.Y = p.Y * s
	return
}

func main() {
	p1 := new(Point)
	p1.X = 3
	p1.Y = 4
	fmt.Printf("The length of the vector p1 is: %f\n", p1.Abs()) // calling Abs() func

	p2 := &Point{4, 5}
	fmt.Printf("The length of the vector p2 is: %f\n", p2.Abs()) // calling Abs() func

	p1.Scale(5)                                                  // calling Scale() func
	fmt.Printf("The length of the vector p1 is: %f\n", p1.Abs()) // calling Abs() func
	fmt.Printf("Point p1 scaled by 5 has the following coordinates: X %f - Y %f", p1.X, p1.Y)
}
