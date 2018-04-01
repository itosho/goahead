package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	lines = flag.Int("n", 10, "lines")
)

func main() {
	flag.Parse()
	args := flag.Args()

	fp, err := os.Open(args[0])
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

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
