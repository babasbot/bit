package bit

import(
	"image"
	"image/color"
)

func Average(img image.Image, center image.Point, n int) color.RGBA {
	kernel := make([][]float64, n)

	var factor float64 = 1.0 
	var bias   float64 = 0.0

	for i := 0; i < n; i++ {
		kernel[i] = make([]float64, n)

		for j := 0; j < n; j++ {
			kernel[i][j] = 1.0
		}
	}

	// TODO: Replace Point with image.Point
	// see: https://github.com/babasbot/bit/issues/2
	_center := Point{x: center.X, y: center.Y}

	return convolution(kernel, factor, bias, img, _center)
}
