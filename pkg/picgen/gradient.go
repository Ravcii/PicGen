package picgen

import (
	"image/color"

	operations "github.com/Ravcii/PicGen/pkg/operations"
)

type Gradient struct {
	from, to color.Color
}

func NewGradient(from, to color.Color) *Gradient {
	return &Gradient{from, to}
}

// Returns color at given pixel out of total pixels.
func (grad *Gradient) At(at, total int) color.RGBA {
	fromR, fromG, fromB, _ := grad.from.RGBA()
	toR, toG, toB, _ := grad.to.RGBA()

	if at == 0 {
		at = 1
	}

	r := float64(operations.Diff(fromR, toR)/uint32(total)) / 256
	g := float64(operations.Diff(fromG, toG)/uint32(total)) / 256
	b := float64(operations.Diff(fromB, toB)/uint32(total)) / 256

	return color.RGBA{
		R: uint8(fromR) + uint8(r)*uint8(at),
		G: uint8(fromG) + uint8(g)*uint8(at),
		B: uint8(fromB) + uint8(b)*uint8(at),
		// A: uint8(fromA) + uint8(a)*uint8(at),
		A: 255,
	}
}
