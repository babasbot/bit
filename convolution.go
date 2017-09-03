package bit

import (
	"image"
	"image/color"
)

func Blur(img image.Image) image.Image {
	var factor, bias float64 = 1.0, 0.0

	kernel := [][]float64{
		{0.0, 0.2, 0.0},
		{0.2, 0.2, 0.2},
		{0.0, 0.2, 0.0},
	}

	return Convolution(kernel, factor, bias, img)
}

func MotionBlur(img image.Image) image.Image {
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

	return Convolution(kernel, factor, bias, img)
}

func Convolution(kernel [][]float64, factor, bias float64, src image.Image) image.Image {
	bounds := src.Bounds()
	img := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			point := Point{x: x, y: y}
			pixel := convolution(kernel, factor, bias, src, point)

			img.Set(x, y, pixel)
		}
	}

	return img
}

func convolution(kernel [][]float64, factor, bias float64, img image.Image, point Point) color.RGBA {
	padding := len(kernel) / 2

	x := point.x - padding
	y := point.y - padding

	var r, g, b float64

	for _, row := range kernel {
		for _, value := range row {
			_r, _g, _b, _ := img.At(x, y).RGBA()

			r += factor*float64(_r>>8)*value + bias
			g += factor*float64(_g>>8)*value + bias
			b += factor*float64(_b>>8)*value + bias

			x++
		}

		x = point.x - padding
		y++
	}

	r = truncate(r)
	g = truncate(g)
	b = truncate(b)

	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}
