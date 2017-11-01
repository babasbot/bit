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

func (img *Image) Pixel(psize, margin int) Image {
	bounds := img.Bounds()
	newimg := Image{image.NewRGBA(bounds)}

	for y := bounds.Min.Y; y < bounds.Max.Y; y += psize {
		for x := bounds.Min.X; x < bounds.Max.X; x += psize {
			point := image.Point{x, y}
			avgcol := img.AverageColor(point, psize)

			newimg.fill(avgcol, point, psize, margin)
		}
	}

	return newimg
}

func (img *Image) fill(col color.Color, point image.Point, psize, margin int) {
	rgba := img.image.(*image.RGBA)

	for x := point.X; x < point.X+psize; x++ {
		for y := point.Y; y < point.Y+psize; y++ {
			if y < point.Y+margin || x < point.X+margin {
				rgba.Set(x, y, color.Gray{0})
			} else {
				rgba.Set(x, y, col)
			}
		}
	}
}
