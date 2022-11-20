package main

import "os"

func main() {
	img, err := NewImage(50, 50)
	if err != nil {
		panic(err)
	}

	err = img.Generate(os.Stdout)
	if err != nil {
		panic(err)
	}
}
