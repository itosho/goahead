package main

type unit interface {
	display(string) int
}

func head(u unit, file string) int {
	return u.display(file)
}

func getUnit(lines int, bytes int) unit {
	if bytes > 0 {
		return &byteUnit{bytes}
	} else {
		return &lineUnit{lines}
	}
}
