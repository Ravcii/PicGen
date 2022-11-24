package picgen

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"io"
)

type PicOptions func(*Image)

type Image struct {
	h          *image.RGBA
	Background Background
}

func NewImage(x, y int, options ...PicOptions) *Image {
	image := &Image{h: image.NewRGBA(image.Rect(0, 0, x, y))}

	for _, fn := range options {
		fn(image)
	}

	return image
}

func (img *Image) Generate() {
	img.Background.Draw(img)
}

// AsJpeg writes final image to dst as .jpg with given quality.
// Quality is an int in a [0, 100] range.
func (img *Image) AsJpeg(dst io.Writer, quality int) error {
	err := jpeg.Encode(dst, img.h, &jpeg.Options{Quality: quality})
	if err != nil {
		return fmt.Errorf("error while encoding the image: %w", err)
	}

	return nil
}

// AsPng writes final image to dst as .png.
func (img *Image) AsPng(dst io.Writer) error {
	err := png.Encode(dst, img.h)
	if err != nil {
		return fmt.Errorf("error while encoding the image: %w", err)
	}

	return nil
}
