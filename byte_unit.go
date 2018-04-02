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
		return ExitError, err
	}

	if *bytes >= len(body) {
		fmt.Print(string(body))
	} else {
		fp, err := os.Open(file)
		if err != nil {
			return ExitError, err
		}
		defer fp.Close()

		reader := bufio.NewReader(fp)
		body, err := reader.Peek(*bytes)
		if err != nil {
			return ExitError, err
		}
		fmt.Println(string(body))
	}

	return ExitSuccess, nil
}
