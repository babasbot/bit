package bit

import (
	"image"
	"image/color"
)

func (i *Image) And(col color.Color) Image {
	bounds := i.Bounds()
	newimg := image.NewRGBA(bounds)

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			c := Color{}
			c = NewRGBA(i.At(x, y))

			newimg.Set(x, y, c.And(col))
		}
	}

	return Image{newimg}
}

func (c *Color) And(p color.Color) color.Color {
	r, g, b, _ := p.RGBA()

	and := color.RGBA{}
	and.R = c.rgba.R & uint8(r>>8)
	and.G = c.rgba.G & uint8(g>>8)
	and.B = c.rgba.B & uint8(b>>8)
	and.A = 255

	return and
}
