package main

import (
	"os"
	"text/template"
)

type example struct {
	Name, LastName string
}

func main() {
	tEmpty := template.New("template test")
	var tempExp example = example{Name: "David", LastName: "David"}
	//empty pipeline following if
	tEmpty = template.Must(tEmpty.Parse("Empty pipeline if demo: {{if ``}} Will not print. {{end}}\n"))
	tEmpty.Execute(os.Stdout, tempExp)

	tWithValue := template.New("template test")
	//non empty pipeline following if condition
	tWithValue = template.Must(tWithValue.Parse("Non empty pipeline if demo: {{if .Name}} {{.Name}} {{end}}\n"))
	tWithValue.Execute(os.Stdout, tempExp)

	tIfElse := template.New("template test")
	//non empty pipeline following if condition
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if .LastName}} {{.LastName}} {{else}} {{.Name}} {{end}}\n"))
	tIfElse.Execute(os.Stdout, tempExp)
}
