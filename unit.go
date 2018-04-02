package main

type unit interface {
	display(string) int
}

func head(u unit, file string) int {
	return u.display(file)
}

func getUnit() unit {
	if *bytes > 0 {
		return new(byteUnit)
	} else {
		return new(lineUnit)
	}
}
