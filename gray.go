package bit

import (
	"image"
	"image/color"
)

func Gray(src image.Image) image.Image {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			pixel := gray(src, Point{x: x, y: y})

			dst.Set(x, y, pixel)
		}
	}

	return dst
}

func gray(img image.Image, point Point) color.Gray {
	r, g, b, _ := img.At(point.x, point.y).RGBA()

	col := 0.3*float32(r) + 0.59*float32(g) + 0.11*float32(b)

	return color.Gray{uint8(uint32(col) >> 8)}
}
