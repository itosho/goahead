package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type line struct{}

func (l *line) display(file string) {
	fp, err := os.Open(file)
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
		log.Fatal(err)
	}
}
