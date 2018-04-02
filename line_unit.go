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
		return 1, err
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
		return 1, err
	}

	return 0, nil
}
