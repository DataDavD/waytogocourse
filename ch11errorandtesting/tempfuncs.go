package main

import (
	"os"
	"text/template"
)

func main() {
	t := template.New("test")
	t = template.Must(t.Parse("{{with $x := `hello`}} {{printf `%s %s` $x `Mary`}} {{end}}!\n"))
	t.Execute(os.Stdout, nil)
	t = template.Must(t.Parse(`{{with $x := "david"}} The length of the word {{.}} is {{len .}} {{end}}`))
	t.Execute(os.Stdout, nil)
}
