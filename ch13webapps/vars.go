package main

import (
	"os"
	"text/template"
)

func main() {
	t := template.New("test")
	t = template.Must(t.Parse("{{with $3 := `hello`}}{{$3}}{{end}}!\n"))
	t.Execute(os.Stdout, nil)
	t = template.Must(t.Parse("{{with $x3 := `bonjour`}}{{$x3}}{{end}}!\n"))
	t.Execute(os.Stdout, nil)
	t = template.Must(t.Parse("{{with $x_1 := `salud`}}{{$x_1}} {{.}} {{$x_1}}{{end}}!\n"))
	t.Execute(os.Stdout, nil)
}
