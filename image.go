package picgen

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"log"
	"math"
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
	rect := image.Rect(0, 0, img.x+1, img.y+1)
	paletted := image.NewPaletted(rect, img.Palette)

	startingColor := color.RGBA{4, 59, 92, 255}
	endingColor := color.RGBA{11, 127, 171, 127}
	r2, g2, b2, _ := startingColor.RGBA()
	r1, g1, b1, _ := endingColor.RGBA()

	sr := (r1 - r2) / uint32(img.x)
	sg := (g1 - g2) / uint32(img.x)
	sb := (b1 - b2) / uint32(img.x)
	sa := 127 / float64((img.x))

	// sr, sg, sb, sa := 7/float64(img.x)*255, (127-59)/float64(img.x)*255, (171-92)/float64(img.x)*255, (127-255)/float64(img.x)*255

	for x := 0; x <= img.x; x++ {
		for y := 0; y <= img.y; y++ {
			multiplier := math.Pow(2, 11)
			r := uint16(uint32(sr)*uint32(y)) * uint16(multiplier)
			g := uint16(uint32(sg)*uint32(y)) * uint16(multiplier)
			b := uint16(uint32(sb)*uint32(y)) * uint16(multiplier)
			a := uint16(uint32(sa)*uint32(y)) * uint16(multiplier)
			// paletted.SetColorIndex(x, y, img.RandomColorIndex())
			log.Println(r, g, b, a)
			paletted.SetRGBA64(x, y, color.RGBA64{r, g, b, 1})
		}
	}

	log.Println(sr, sg, sb, sa)

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
