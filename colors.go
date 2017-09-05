package bit

import (
	"image"
	"image/color"
)

func And(src image.Image, col color.RGBA) image.Image {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			point := Point{x: x, y: y}
			pixel := and(src, point, col)

			dst.Set(x, y, pixel)
		}
	}

	return dst
}

func and(src image.Image, point Point, col color.RGBA) color.RGBA {
	r, g, b, _ := src.At(point.x, point.y).RGBA()

	pixel := color.RGBA{}

	pixel.R = uint8(r>>8) & col.R
	pixel.G = uint8(g>>8) & col.G
	pixel.B = uint8(b>>8) & col.B
	pixel.A = 255

	return pixel
}
