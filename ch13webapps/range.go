package main

import (
	"os"
	"text/template"
)

func main() {
	s := []int{1, 2, 4, 5, 6}
	t := template.New("range test")
	t = template.Must(t.Parse("{{range .}} {{.}} {{end}}"))
	t.Execute(os.Stdout, s)
}
