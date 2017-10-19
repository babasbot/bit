package bit

import (
	"image"
	"image/color"
)

func (i *Image) Convolution(kernel [][]float64, factor, bias float64) Image {
	bounds := i.Bounds()
	newimg := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			point := image.Point{X: x, Y: y}
			col := i.Convolve(point, kernel, factor, bias)

			newimg.Set(x, y, col)
		}
	}

	return Image{newimg}
}

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

func (i *Image) Blur() Image {
	var factor, bias float64 = 1.0, 0.0

	kernel := [][]float64{
		{0.0, 0.2, 0.0},
		{0.2, 0.2, 0.2},
		{0.0, 0.2, 0.0},
	}

	return i.Convolution(kernel, factor, bias)
}

func (i *Image) MotionBlur() Image {
	var factor, bias float64 = 1.0 / 9.0, 0.0

	kernel := [][]float64{
		{1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1},
	}

	return i.Convolution(kernel, factor, bias)
}

func (i *Image) FindEdges() Image {
	var factor, bias float64 = 1.0, 0.0

	kernel := [][]float64{
		{0, 0, -1, 0, 0},
		{0, 0, -1, 0, 0},
		{0, 0, 2, 0, 0},
		{0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0},
	}

	return i.Convolution(kernel, factor, bias)
}
