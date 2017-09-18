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

			pixel = redEyes2(dst, Point{x: x, y: y})
			dst.Set(x, y, pixel)
		}
	}

	return dst
}

func redEyes(img image.Image, point Point) color.RGBA {
	r, g, b, _ := img.At(point.x, point.y).RGBA()

	intensity := float64(r) / (float64(g+b) / 2.0)

	if intensity > 1.5 {
		gray := grayGB(img, point).Y
		gray = normalize(int32(gray) + 8)

		return color.RGBA{gray, gray, gray, 255}

	}

	return color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), 255}
}

func redEyes2(img image.Image, point Point) color.RGBA {
	r, g, b, _ := img.At(point.x, point.y).RGBA()

	intensity := float64(r) / (float64(g+b) / 2.0)

	if intensity > 0.8 {
		return removeRed(img, point)
	}

	return color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), 255}
}

func removeRed(img image.Image, point Point) color.RGBA {
	r, g, b, _ := img.At(point.x, point.y).RGBA()

	averageGB := int32(g>>8+b>>8) / 2

	if abs(int32(r>>8)-averageGB) >= 10 {
		g := int32(grayRB(img, point).Y) + 16
		gray := uint8(normalize(g))

		return color.RGBA{gray, gray, gray, 255}
	}

	return color.RGBA{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8), 255}
}
