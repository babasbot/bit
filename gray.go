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

func GrayR(src image.Image) image.Image {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			pixel := grayR(src, Point{x: x, y: y})

			dst.Set(x, y, pixel)
		}
	}

	return dst
}

func grayR(img image.Image, point Point) color.Gray {
	r, _, _, _ := img.At(point.x, point.y).RGBA()

	return color.Gray{uint8(r)}
}

func GrayG(src image.Image) image.Image {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			pixel := grayG(src, Point{x: x, y: y})

			dst.Set(x, y, pixel)
		}
	}

	return dst
}

func grayG(img image.Image, point Point) color.Gray {
	_, g, _, _ := img.At(point.x, point.y).RGBA()

	return color.Gray{uint8(g >> 8)}
}

func GrayB(src image.Image) image.Image {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			pixel := grayB(src, Point{x: x, y: y})

			dst.Set(x, y, pixel)
		}
	}

	return dst
}

func grayB(img image.Image, point Point) color.Gray {
	_, _, b, _ := img.At(point.x, point.y).RGBA()

	return color.Gray{uint8(b >> 8)}
}

func GrayRG(src image.Image) image.Image {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			pixel := grayRG(src, Point{x: x, y: y})

			dst.Set(x, y, pixel)
		}
	}

	return dst
}

func grayRG(img image.Image, point Point) color.Gray {
	r, g, _, _ := img.At(point.x, point.y).RGBA()

	return color.Gray{uint8((r + g) / 2 >> 8)}
}

func GrayRB(src image.Image) image.Image {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			pixel := grayRB(src, Point{x: x, y: y})

			dst.Set(x, y, pixel)
		}
	}

	return dst
}

func grayRB(img image.Image, point Point) color.Gray {
	r, _, b, _ := img.At(point.x, point.y).RGBA()

	return color.Gray{uint8((r + b) / 2 >> 8)}
}

func GrayGB(src image.Image) image.Image {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			pixel := grayGB(src, Point{x: x, y: y})

			dst.Set(x, y, pixel)
		}
	}

	return dst
}

func grayGB(img image.Image, point Point) color.Gray {
	_, g, b, _ := img.At(point.x, point.y).RGBA()

	return color.Gray{uint8((g + b) / 2 >> 8)}
}

func GrayRGB(src image.Image) image.Image {
	bounds := src.Bounds()
	dst := image.NewRGBA(bounds)

	for x := bounds.Min.X; x < bounds.Max.X; x++ {
		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			pixel := grayRGB(src, Point{x: x, y: y})

			dst.Set(x, y, pixel)
		}
	}

	return dst
}

func grayRGB(img image.Image, point Point) color.Gray {
	r, g, b, _ := img.At(point.x, point.y).RGBA()

	return color.Gray{uint8((r + g + b) / 3 >> 8)}
}
