package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
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
	args := flag.Args()

	if flag.NFlag() > 1 {
		log.Fatal("Too many optional arguments.")
	}

	if flag.NArg() != 1 {
		log.Fatal("invalid file paths.")
	}

	filePath := args[0]
	fp, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	if *bytes > 0 { // use -c option
		body, e := ioutil.ReadFile(filePath)
		if e != nil {
			log.Fatal(err)
		}

		if *bytes >= len(body) {
			fmt.Print(string(body))
		} else {
			reader := bufio.NewReader(fp)
			b, e := reader.Peek(*bytes)
			if e != nil {
				log.Fatal(e)
			}
			fmt.Println(string(b))
		}

	} else { // use -n option
		scanner := bufio.NewScanner(fp)

		i := 0
		for scanner.Scan() {
			i++
			if i > *lines {
				break
			}
			fmt.Println(scanner.Text())
		}

		if err = scanner.Err(); err != nil {
			fmt.Println(err)
		}
	}
}

func usage() {
	fmt.Println("usage: goahead [-n lines | -c bytes] [file]")
	flag.PrintDefaults()
	os.Exit(0)
}
