package bit

import (
	"image"
	"image/color"
)

func HighContrast(src image.Image) image.Image {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			pixel := highContrast(src, Point{x: x, y: y})

			dst.Set(x, y, pixel)
		}
	}

	return dst
}

func highContrast(img image.Image, point Point) color.Gray {
	gray := gray(img, point)

	if gray.Y > 128 {
		return color.Gray{255}
	}

	return color.Gray{0}
}

func InverseHighContrast(src image.Image) image.Image {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			pixel := inverseHighContrast(src, Point{x: x, y: y})

			dst.Set(x, y, pixel)
		}
	}

	return dst
}

func inverseHighContrast(img image.Image, point Point) color.Gray {
	gray := gray(img, point)

	if gray.Y < 128 {
		return color.Gray{255}
	}

	return color.Gray{0}
}
