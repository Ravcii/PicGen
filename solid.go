package picgen

import "image/color"

type SolidBackground struct {
	color color.Color
}

func WithSolid(c color.Color) PicOptions {
	return func(i *Image) {
		i.Background = &SolidBackground{c}
	}
}

func (s SolidBackground) Draw(image *Image) {

}
