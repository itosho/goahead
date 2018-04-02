package main

import (
	"bufio"
	"fmt"
	"os"
)

type lineUnit struct{}

func (lu *lineUnit) display(file string) (int, error) {
	fp, err := os.Open(file)
	if err != nil {
		return ExitError, err
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
		return ExitError, err
	}

	return ExitSuccess, nil
}
