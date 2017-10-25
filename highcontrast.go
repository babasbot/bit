package bit

import (
	"image"
	"image/color"
)

func (img *Image) HighContrast() Image {
	bounds := img.Bounds()
	newimg := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			col := NewRGBA(img.At(x, y))

			newimg.Set(x, y, col.HighContrast())
		}
	}

	return Image{newimg}
}

func (col *Color) HighContrast() color.Gray {
	graycol := col.Gray()

	if graycol.Y > 128 {
		return color.Gray{255}
	}

	return color.Gray{0}
}

func (img *Image) InverseHighContrast() Image {
	bounds := img.Bounds()
	newimg := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			col := NewRGBA(img.At(x, y))

			newimg.Set(x, y, col.InverseHighContrast())
		}
	}

	return Image{newimg}
}

func (col *Color) InverseHighContrast() color.Gray {
	graycol := col.Gray()

	if graycol.Y > 128 {
		return color.Gray{0}
	}

	return color.Gray{255}
}
