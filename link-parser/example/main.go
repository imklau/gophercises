package main

import (
	"bytes"
	"fmt"
	"log"
	"os"

	"gophercises/linkparser"
)

func main() {
	file, err := os.ReadFile("example.html")

	if err != nil {
		log.Fatal("Unable to read the file", err)
	}

	links, err := linkparser.Parse(bytes.NewReader(file))

	if err != nil {
		fmt.Println("Parse error:", err)
	}

	fmt.Printf("%+v\n", links)
}