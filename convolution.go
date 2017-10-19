package bit

import (
	"image"
	"image/color"
)

func (i *Image) Convolve(point image.Point, kernel [][]float64, factor, bias float64) color.Color {
	padding := len(kernel) / 2

	x := point.X - padding
	y := point.Y - padding

	rgba := RGBAf64{R: 0.0, G: 0.0, B: 0.0, A: 255.0}

	for _, row := range kernel {
		for _, value := range row {
			r, g, b, _ := i.At(x, y).RGBA()

			rgba.R += factor*float64(r>>8)*value + bias
			rgba.G += factor*float64(g>>8)*value + bias
			rgba.B += factor*float64(b>>8)*value + bias

			x++
		}

		x = point.X - padding
		y++
	}

	return rgba
}
