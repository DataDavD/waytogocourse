package main

import "fmt"

// basic data structure upon which we'll define methods
type employee struct {
	salary float32
}

// a method which will add a specified percent to an
// employees salary
func (this *employee) giveRaise(pct float32) {
	this.salary = this.salary * (1 + pct)
}

func main() {
	var e = new(employee)
	e.salary = 60000
	fmt.Println("employee currenlty makes: ", e.salary)
	e.giveRaise(0.6)
	fmt.Println("employee now makes: ", e.salary)
}
