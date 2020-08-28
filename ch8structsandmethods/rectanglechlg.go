package main

import "fmt"

// struct of type Rectangle
type Rectangle struct {
	length, width int
}

// method calculating area of rectangle
func (r *Rectangle) Area() int {
	return r.length * r.width
}

// method calculating perimeter of rectangle
func (r *Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func main() {
	r1 := Rectangle{4, 3}
	fmt.Println("Rectangle is: ", r1)
	fmt.Println("Rectangle area is: ", r1.Area())           // calling method of area
	fmt.Println("Rectangle perimeter is: ", r1.Perimeter()) // calling method of perimeter
}
