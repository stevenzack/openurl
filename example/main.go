package main

import (
	"log"

	"github.com/stevenzack/openurl"
)

func main() {
	if err := openurl.Open("google.com"); err != nil {
		log.Fatal(err)
	}
}
