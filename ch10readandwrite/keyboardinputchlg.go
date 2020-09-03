package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var nrchars, nrwords, nrlines int

func main() {
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter some inputs; end by entering 'S'")
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Printf("An error occurred: %s\n", err)
			return
		}

		if input == "S\n" {
			fmt.Println("Here are the counts from the input:")
			fmt.Printf("Number of characters: %d\n", nrchars)
			fmt.Printf("Number of words: %d\n", nrwords)
			fmt.Printf("Number of lines: %d\n", nrlines)
			os.Exit(0)
		}
		Counters(input)
	}

}

func Counters(input string) {
	nrchars += len(input) - 1                // -1 for \n
	nrwords += strings.Count(input, " ") + 1 // num spaces + 1
	nrlines++
}
