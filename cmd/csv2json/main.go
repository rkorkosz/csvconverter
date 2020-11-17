package main

import (
	"log"
	"os"

	"github.com/rkorkosz/csvconverter/pkg/csvconverter"
)

func main() {
	err := csvconverter.ConvertJSON(os.Stdin, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}
