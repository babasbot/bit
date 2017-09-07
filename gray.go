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

func gray(img image.Image, point Point) color.Gray {
	r, g, b, _ := img.At(point.x, point.y).RGBA()

	col := 0.3*float32(r) + 0.59*float32(g) + 0.11*float32(b)

	return color.Gray{uint8(uint32(col) >> 8)}
}

func highContrast(img image.Image, point Point) color.Gray {
	gray := gray(img, point)

	if gray.Y > 128 {
		return color.Gray{255}
	}

	return color.Gray{0}
}

func inverseHighContrast(img image.Image, point Point) color.Gray {
	gray := gray(img, point)

	if gray.Y < 128 {
		return color.Gray{255}
	}

	return color.Gray{0}
}
