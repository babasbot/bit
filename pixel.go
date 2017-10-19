package bit

import (
	"image"
	"image/color"
)

func (img *Image) AverageColor(point image.Point, n int) color.Color {
	var factor, bias float64 = 1.0 / float64(int64(n*n)), 0.0

	kernel := make([][]float64, n)
	for i := 0; i < n; i++ {
		kernel[i] = make([]float64, n)

		for j := 0; j < n; j++ {
			kernel[i][j] = 1.0
		}
	}

	center := image.Point{X: point.X + n/2, Y: point.Y + n/2}

	return img.Convolve(center, kernel, factor, bias)
}
