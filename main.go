package main

import (
	"flag"
	"log"
)

func main() {
	filePath := flag.String("file", "", "The path to the file")

	flag.Parse()

	if *filePath == "" {
		log.Fatal("You must specify a file path using the -file flag.")
	}

}
