package main

type unit interface {
	display(string) (int, error)
}

func head(u unit, file string) (int, error) {
	return u.display(file)
}

func getUnit() unit {
	if *bytes > 0 {
		return new(byteUnit)
	} else {
		return new(lineUnit)
	}
}
