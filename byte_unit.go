package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

type byteUnit struct{}

func (bu *byteUnit) display(file string) int {
	body, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "File Read Error. The following are the details.")
		fmt.Fprintln(os.Stderr, err)
		return ExitError
	}

	if *bytes >= len(body) {
		fmt.Print(string(body))
	} else {
		fp, err := os.Open(file)
		if err != nil {
			fmt.Fprintln(os.Stderr, "File Open Error. The following are the details.")
			fmt.Fprintln(os.Stderr, err)
			return ExitError
		}
		defer fp.Close()

		reader := bufio.NewReader(fp)
		body, err := reader.Peek(*bytes)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Byte Peek Error. The following are the details.")
			fmt.Fprintln(os.Stderr, err)
			return ExitError
		}
		fmt.Println(string(body))
	}

	return ExitSuccess
}
