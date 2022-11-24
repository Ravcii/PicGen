package picgen

import (
	"image/color"
)

type Gradient struct {
	From, To color.Color
}

func WithGradientBackground(gradient *Gradient) PicOptions {
	return func(i *Image) {
		i.Background = gradient
	}
}

// GetSteps returns a single step for a value to change between
// first and last pixel of a gradient.
func (grad *Gradient) GetSteps(size int) (r, g, b, a float64) {
	fromR, fromG, fromB, fromA := grad.From.RGBA()
	toR, toG, toB, toA := grad.To.RGBA()

	// s := float64(size)

	r = (float64(toR) - float64(fromR))
	g = (float64(toG) - float64(fromG))
	b = (float64(toB) - float64(fromB))
	a = (float64(toA) - float64(fromA))

	return r, g, b, a
}

// Returns color at given pixel out of total pixels.
func (grad *Gradient) Draw(img *Image) {
	X, Y := img.h.Rect.Dx(), img.h.Rect.Dy()
	fr, fg, fb, fa := grad.From.RGBA()
	r, g, b, a := grad.GetSteps(Y)

	for x := 0; x <= X; x++ {
		for y := 0; y <= Y; y++ {
			at := float64(y) / float64(Y)

			c := color.NRGBA{
				R: uint8(float64(fr) + r*at),
				G: uint8(float64(fg) + g*at),
				B: uint8(float64(fb) + b*at),
				A: uint8(float64(fa) + a*at),
			}
			img.h.Set(x, y, c)
		}
	}
}
