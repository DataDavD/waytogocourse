package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"os"
)

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func main() {
	// write your implementation here
	// didn't use bufio since i know file size is small
	vc := new(VCard)
	file, _ := os.Open("vcard.gob")
	defer file.Close()

	dec := gob.NewDecoder(file)
	err := dec.Decode(vc)
	if err != nil {
		log.Fatalf("Error in decoding: %s", err)
	}
	fmt.Printf("%+v", *vc)
	for _, v := range vc.Addresses {
		fmt.Printf("%+v", *v)
	}
}
