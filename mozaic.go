package bit

import (
	"image"
)

func (img *Image) Mozaic(n int) Image {
	grayimg := img.Gray()
	bounds := grayimg.Bounds()

	ymax := (bounds.Max.Y / n) * bounds.Max.Y
	xmax := (bounds.Max.X / n) * bounds.Max.X

	rectangle := image.Rect(0, 0, xmax, ymax)
	newimg := image.NewRGBA(rectangle)

	point := image.Point{0, 0}

	for y := 0; y < bounds.Max.Y; y += n {
		for x := 0; x < bounds.Max.X; x += n {
			avgcol := img.AverageColor(image.Point{x, y}, n)
			avgimg := grayimg.And(avgcol)

			imagefill(&avgimg, newimg, point)

			point.X += bounds.Max.X
		}
		point.X = 0
		point.Y += bounds.Max.Y
	}

	return Image{newimg}
}

func imagefill(src *Image, dst *image.RGBA, point image.Point) {
	bounds := src.Bounds()

	ymax := point.Y + bounds.Max.Y
	xmax := point.X + bounds.Max.X

	for y := point.Y; y < ymax; y++ {
		for x := point.X; x < xmax; x++ {
			_x := x % bounds.Max.X
			_y := y % bounds.Max.Y

			col := src.At(_x, _y)

			dst.Set(x, y, col)
		}
	}
}
