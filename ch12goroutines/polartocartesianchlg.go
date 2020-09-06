package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"runtime"
	"strconv"
	"strings"
)

type polar struct {
	rad   float64
	angle float64
}

type cartesian struct {
	x float64
	y float64
}

const result = "Polar: radius=%.02f angle=%.02f degrees -- Cartesian: x=%.02f y=%.02f\n"

var prompt = "Enter a radius and an angle (in degrees), e.g., 12.5 90, " + "or %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { // Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+C")
	}
}

func main() {
	questions := make(chan polar)
	defer close(questions)
	answers := createSolver(questions)
	defer close(answers)
	interact(questions, answers)
}

func createSolver(questions chan polar) chan cartesian {
	answer := make(chan cartesian)
	go func() {
		for {
			polarCoor := <-questions
			degree := polarCoor.angle * math.Pi / 180.0
			x := polarCoor.rad * math.Cos(degree)
			y := polarCoor.rad * math.Sin(degree)
			answer <- cartesian{x, y}
		}
	}()
	return answer
}
func interact(questions chan polar, answers chan cartesian) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(prompt)
	for {
		fmt.Printf("Radius and angle: ")
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		line = line[:len(line)-1] // chop off newline character
		if numbers := strings.Fields(line); len(numbers) == 2 {
			polars, err := floatsToStrings(numbers)
			if err != nil {
				_, err := fmt.Fprintln(os.Stderr, "invalid number")
				if err != nil {
					log.Fatal("Program failed with error", err)
				}
				continue
			}
			questions <- polar{polars[0], polars[1]}
			coord := <-answers
			fmt.Printf(result, polars[0], polars[1], coord.x, coord.y)
		} else {
			_, err := fmt.Fprintln(os.Stderr, "invalid input")
			if err != nil {
				log.Fatal("Program failed with error", err)
			}
		}
	}
	fmt.Println()
}

func floatsToStrings(numbers []string) ([]float64, error) {
	var floats []float64
	for _, number := range numbers {
		if x, err := strconv.ParseFloat(number, 64); err != nil {
			return nil, err
		} else {
			floats = append(floats, x)
		}
	}
	return floats, nil
}
