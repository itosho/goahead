package main

type delimiter interface {
	display(string)
}

func head(d delimiter, file string) {
	d.display(file)
}

func getDelimiter() delimiter {
	if *bytes > 0 {
		return new(byte)
	} else {
		return new(line)
	}
}
