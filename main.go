package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

const (
	ExitSuccess int = iota
	ExitError
	ExitFileError
)

var (
	lines = flag.Int("n", 10, "lines")
	bytes = flag.Int("c", 0, "bytes")
)

func main() {
	flag.Usage = usage

	flag.Parse()

	if flag.NFlag() > 1 {
		log.Fatal("Too many options. Please specify `n` or `c` for option.")
	}

	if flag.NArg() != 1 {
		log.Fatal("Invalid Args. Please specify only one file.")
	}

	args := flag.Args()
	file := args[0]

	unit := getUnit()
	code := head(unit, file)

	os.Exit(code)
}

func usage() {
	fmt.Println("usage: goahead [-n lines | -c bytes] [file]")
	flag.PrintDefaults()
	os.Exit(ExitSuccess)
}
