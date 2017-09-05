package bit

import (
	"image"
	"image/color"
)

func AndyWarhol(src image.Image) image.Image {
	col1 := color.RGBA{R: 93, G: 211, B: 201, A: 255}
	col2 := color.RGBA{R: 241, G: 178, B: 220, A: 255}
	col3 := color.RGBA{R: 225, G: 206, B: 44, A: 255}
	col4 := color.RGBA{R: 215, G: 63, B: 50, A: 255}

	bounds := src.Bounds()
	rectangle := image.Rect(0, 0, bounds.Max.X*2, bounds.Max.Y*2)

	dst := image.NewRGBA(rectangle)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			point := Point{x: x, y: y}

			p1 := and(src, point, col1)
			p2 := and(src, point, col2)
			p3 := and(src, point, col3)
			p4 := and(src, point, col4)

			dst.Set(x, y, p1)
			dst.Set(x, y+bounds.Max.Y, p2)
			dst.Set(x+bounds.Max.X, y, p3)
			dst.Set(x+bounds.Max.X, y+bounds.Max.Y, p4)
		}
	}

	return dst
}

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
