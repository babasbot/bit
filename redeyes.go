package bit

import (
	"image"
	"image/color"
)

func RedEyes(src image.Image) image.Image {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			pixel := redEyes(src, Point{x: x, y: y})

			dst.Set(x, y, pixel)
		}
	}

	return dst
}

func redEyes(img image.Image, point Point) color.RGBA {
	r, g, b, _ := img.At(point.x, point.y).RGBA()

	intensity := float64(r) / (float64(g+b) / 2.0)

	if intensity > 1.5 {
		r = (g + b) / 2
	}

	return color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), 255}
}
