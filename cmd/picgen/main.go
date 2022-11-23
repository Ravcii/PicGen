package main

import (
	"os"

	"github.com/Ravcii/PicGen/pkg/picgen"
)

func main() {
	img, err := picgen.NewImage(1000, 1000)
	if err != nil {
		panic(err)
	}

	f, err := os.Create("test.jpg")
	if err != nil {
		panic(err)
	}

	err = img.Generate(f)
	if err != nil {
		panic(err)
	}
}
