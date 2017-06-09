package main

import (
	"flag"
	"log"
	"fmt"
	"os"
)

var (
	algoType = ""
	filepath = ""
)

func init() {
	generateUsage()
	flag.StringVar(&algoType, "type", "", "Type of the algorythm")
	flag.StringVar(&filepath, "path", "", "Path to a file with values")
	flag.Parse()
}

func main() {
	if filepath == "" {
		log.Fatal("You should pass me a file with values")
	}
	switch algoType {
	case "selection-sort":
		//
	default:
		log.Fatalf("I don't have algorythm `%s`", algoType)
	}
}

func generateUsage() {
	flagSet := flag.CommandLine
	flagSet.Usage = func() {
		fmt.Fprint(os.Stderr, "This app starts the chosen algorythm against data from a file.\n")
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
}