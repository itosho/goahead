package main

type unit interface {
	display(string)
}

func head(u unit, file string) {
	u.display(file)
}

func getUnit() unit {
	if *bytes > 0 {
		return new(byteUnit)
	} else {
		return new(lineUnit)
	}
}
