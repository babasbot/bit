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

func (img *Image) Pixel(n int) Image {
	bounds := img.Bounds()
	newimg := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y += n {
		for x := bounds.Min.X; x < bounds.Max.X; x += n {
			colorfill(img, newimg, image.Point{X: x, Y: y}, n)
		}
	}

	return Image{newimg}
}

func colorfill(src *Image, dst *image.RGBA, point image.Point, n int) {
	col := src.AverageColor(point, n)

	for x := point.X; x < point.X+n; x++ {
		for y := point.Y; y < point.Y+n; y++ {
			dst.Set(x, y, col)
		}
	}
}
