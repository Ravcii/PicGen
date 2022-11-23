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
	fromR, fromG, fromB, fromA := grad.from.RGBA()
	toR, toG, toB, toA := grad.to.RGBA()

	multiplier := float32(at) / float32(total)
	stepR := (operations.Diff(fromR>>8, toR>>8))
	stepG := (operations.Diff(fromG>>8, toG>>8))
	stepB := (operations.Diff(fromB>>8, toB>>8))
	stepA := (operations.Diff(fromA>>8, toA>>8))

	return color.RGBA{
		R: uint8(float64(fromR) + float64(multiplier)*float64(stepR)),
		G: uint8(float64(fromG) + float64(multiplier)*float64(stepG)),
		B: uint8(float64(fromB) + float64(multiplier)*float64(stepB)),
		A: uint8(float64(fromA) + float64(multiplier)*float64(stepA)),
	}
}
