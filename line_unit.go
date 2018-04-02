package main

import (
	"bufio"
	"fmt"
	"os"
)

type lineUnit struct {
	lines int
}

func (lu *lineUnit) display(file string) int {
	fp, err := os.Open(file)
	if err != nil {
		fmt.Fprintln(os.Stderr, "File Open Error. The following are the details.")
		fmt.Fprintln(os.Stderr, err)
		return ExitError
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	i := 0
	for scanner.Scan() {
		i++
		if i > lu.lines {
			break
		}
		fmt.Println(scanner.Text())
	}

	if err = scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "File Scan Error. The following are the details.")
		fmt.Fprintln(os.Stderr, err)
		return ExitError
	}

	return ExitSuccess
}
