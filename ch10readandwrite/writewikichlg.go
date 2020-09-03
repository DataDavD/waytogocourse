package main

import (
	"errors"
	"fmt"
	"io/ioutil"
)

// Page struct detailing title and body content
type Page struct {
	Title string
	Body  []byte
}

// Page save method, returns error
func (p *Page) save() error {
	filename := p.Title + ".txt"
	err := ioutil.WriteFile(filename, p.Body, 0666)
	if err != nil {
		return errors.New("error saving")
	}
	return nil
}

func load(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil

}

func main() {
	testPage := &Page{Title: "David_D_TestPage", Body: []byte("this is a simple Page.")}
	_ = testPage.save()
	p2, _ := load("David_D_TestPage")
	fmt.Println(string(p2.Body))
}
