package bit

import (
	"image"
	"image/color"
)

func Bright(src image.Image, n int32) image.Image {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			pixel := bright(src, Point{x: x, y: y}, n)

			dst.Set(x, y, pixel)
		}
	}

	return dst
}

func bright(img image.Image, point Point, n int32) color.RGBA {
	normalize := func(n int32) uint8 {
		if n > 255 {
			return 255
		} else if n < 0 {
			return 0
		}

		return uint8(n)
	}

	r, g, b, _ := img.At(point.x, point.y).RGBA()
	pixel := color.RGBA{}

	pixel.R = normalize(int32(r>>8) + n)
	pixel.G = normalize(int32(g>>8) + n)
	pixel.B = normalize(int32(b>>8) + n)
	pixel.A = 255

	return pixel
}
