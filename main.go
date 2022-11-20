package main

import "os"

func main() {
	img, err := NewImage(200, 200)
	if err != nil {
		panic(err)
	}

	err = img.Generate(os.Stdout)
	if err != nil {
		panic(err)
	}
}
