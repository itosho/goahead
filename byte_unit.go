package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type byteUnit struct{}

func (bu *byteUnit) display(file string) {
	body, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	if *bytes >= len(body) {
		fmt.Print(string(body))
	} else {
		fp, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}
		defer fp.Close()

		reader := bufio.NewReader(fp)
		body, err := reader.Peek(*bytes)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(body))
	}
}
