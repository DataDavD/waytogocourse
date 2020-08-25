package main

import "fmt"

func main() {
	fmt.Println(Season(3))
	fmt.Println(Season(12))
	fmt.Println(Season(300))
}

func Season(month int) string {

	if (month == 12) || (month == 1) || (month == 2) { // Jan, Feb and Dec have winter
		return "Winter"
	} else if (month == 3) || (month == 4) || (month == 5) { // March, Apr and May have spring
		return "Spring"
	} else if (month == 6) || (month == 7) || (month == 8) { // June, July and Aug have summer
		return "Summer"
	} else if (month == 9) || (month == 10) || (month == 11) { // Sept, Oct and Nov have autumn
		return "Autumn"
	} else { //value outside [1,12], then season is unkown
		return "Season unknown"
	}

}
