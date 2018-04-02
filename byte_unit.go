package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

type byteUnit struct{}

func (bu *byteUnit) display(file string) (int, error) {
	body, err := ioutil.ReadFile(file)
	if err != nil {
		return 1, err
	}

	if *bytes >= len(body) {
		fmt.Print(string(body))
	} else {
		fp, err := os.Open(file)
		if err != nil {
			return 1, err
		}
		defer fp.Close()

		reader := bufio.NewReader(fp)
		body, err := reader.Peek(*bytes)
		if err != nil {
			return 1, err
		}
		fmt.Println(string(body))
	}

	return 0, nil
}
