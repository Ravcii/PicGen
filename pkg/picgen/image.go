package picgen

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"math/rand"
	"time"
)

var basePallete = color.Palette{
	color.RGBA{0, 0, 0, 0},
	color.RGBA{255, 0, 0, 1},
	color.RGBA{0, 255, 0, 1},
	color.RGBA{0, 0, 255, 1},
	color.RGBA{255, 255, 255, 255},
}

type param func(*Image) error

type Image struct {
	x, y    int
	Palette color.Palette
}

func NewImage(x, y int, params ...param) (*Image, error) {
	rand.Seed(time.Now().UTC().UnixNano())
	image := &Image{x: x, y: y, Palette: nil}

	for _, param := range params {
		err := param(image)
		if err != nil {
			return nil, fmt.Errorf("error while applying params: %w", err)
		}
	}

	if image.Palette == nil {
		image.Palette = basePallete
	}

	return image, nil
}

func (img *Image) Generate(out io.Writer) error {
	paletted := image.NewRGBA(image.Rect(0, 0, img.x+1, img.y+1))

	gradient := NewGradient(color.RGBA{255, 0, 0, 255}, color.RGBA{0, 0, 255, 255})

	for x := 0; x <= img.x; x++ {
		for y := 0; y <= img.y; y++ {
			c := gradient.At(y, img.y)
			fmt.Printf("x %d y %d c %v\n", x, y, c)
			paletted.Set(x, y, c)
		}
	}

	err := jpeg.Encode(out, paletted, &jpeg.Options{Quality: 100})
	if err != nil {
		return fmt.Errorf("error while eccoding the image: %w", err)
	}

	return nil
}

func (img *Image) RandomColorIndex() uint8 {
	return uint8(rand.Intn(len(img.Palette)))
}

func (img *Image) RandomColor() color.Color {
	return img.Palette[img.RandomColorIndex()]
}
