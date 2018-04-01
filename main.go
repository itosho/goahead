package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	lines = flag.Int("n", 10, "lines")
	bytes = flag.Int("c", 0, "bytes")
)

func main() {
	flag.Usage = usage

	flag.Parse()

	if flag.NFlag() > 1 {
		log.Fatal("Too many optional arguments.")
	}

	if flag.NArg() != 1 {
		log.Fatal("Invalid file paths.")
	}

	args := flag.Args()
	file := args[0]

	unit := getUnit()
	head(unit, file)
}

func usage() {
	fmt.Println("usage: goahead [-n lines | -c bytes] [file]")
	flag.PrintDefaults()
	os.Exit(0)
}
