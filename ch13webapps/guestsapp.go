package main

import (
	"fmt"
	"html/template"
	"net/http"
)

const addr = "localhost" // adjust ip address
const port = 8000        // adjust port
var guestList []string

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/add", addHandler)
	http.ListenAndServe(fmt.Sprintf("%s:%d", addr, port), nil)
}

// indexHandler serves the main page
func indexHandler(w http.ResponseWriter, req *http.Request) {
	t := template.New("index.html")
	t, err := t.Parse(indexHTML)
	if err != nil {
		message := fmt.Sprintf("bad template: %s", err)
		http.Error(w, message, http.StatusInternalServerError)
	}
	t.Execute(w, guestList)
}

// addHandler add a name to the names list
func addHandler(w http.ResponseWriter, req *http.Request) {
	guest := req.FormValue("name")
	if len(guest) > 0 {
		guestList = append(guestList, guest)
	}
	http.Redirect(w, req, "/", http.StatusFound)
}

var indexHTML = `
<!DOCTYPE html>
<html>
    <head>
		<title>Guest Book ::Web GUI</title>
    </head>
    <body>
		<h1>Guest Book :: Web GUI</h1>
		<form action="/add" method="post">
		Name: <input name="name" /><submit value="Sign Guest Book">
		</form>
		<hr />
		<h4>Previous Guests</h4>
		<ul>
			{{range .}}
			<li>{{.}}</li>
			{{end}}
		</ul>
	</body>
</html>
`
