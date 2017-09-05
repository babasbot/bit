package bit

import (
	"image"
	"image/color"
)

func Mozaic(src image.Image, n int) image.Image {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x += n {
		for y := bounds.Min.Y; y < bounds.Max.Y; y += n {
			point := Point{x: x, y: y}
			mozaic(src, dst, point, n)
		}
	}

	return dst
}

func mozaic(src image.Image, dst *image.RGBA, point Point, n int) {
	center := Point{x: point.x + n/2, y: point.y + n/2}
	pixel := averageColor(src, center, n)

	for x := point.x; x < point.x+n; x++ {
		for y := point.y; y < point.y+n; y++ {
			dst.Set(x, y, pixel)
		}
	}
}

func averageColor(img image.Image, point Point, n int) color.RGBA {
	var factor, bias float64 = 1.0 / float64(int64(n*n)), 0.0
	kernel := make([][]float64, n)

	for i := 0; i < n; i++ {
		kernel[i] = make([]float64, n)

		for j := 0; j < n; j++ {
			kernel[i][j] = 1.0
		}
	}

	return convolution(kernel, factor, bias, img, point)
}
