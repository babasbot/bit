package bit

import (
	"image"
	"math"
)

func (img *Image) Scale(width, height int) Image {
	bounds := img.Bounds()

	rect := image.Rect(0, 0, width, height)
	newimg := image.NewRGBA(rect)

	xratio := float64(bounds.Max.X) / float64(width)
	yratio := float64(bounds.Max.Y) / float64(height)

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			px := int(math.Floor(float64(x) * xratio))
			py := int(math.Floor(float64(y) * yratio))

			newimg.Set(x, y, img.At(px, py))
		}
	}

	return Image{newimg}
}
