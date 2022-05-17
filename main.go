package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	var idea string

	flag.StringVar(&idea, "i", "Idea", "Name of Idea")

	flag.Usage = func() {
		fmt.Printf("Usage of %s\n", os.Args[0])
		fmt.Printf("main -i='Idea'")
		flag.PrintDefaults()
	}

	flag.Parse()

	fmt.Printf("Idea = %s\n", idea)
}
