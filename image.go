package bit

import (
	"image"
	"image/color"
)

type Image struct {
	image image.Image
}

func (i *Image) ColorModel() color.Model {
	return i.image.ColorModel()
}

func (i *Image) Bounds() image.Rectangle {
	return i.image.Bounds()
}

func (i *Image) At(x, y int) color.Color {
	return i.image.At(x, y)
}
