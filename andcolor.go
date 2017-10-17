package bit

import (
	"image/color"
)

func (c *Color) And(p color.Color) color.Color {
	r, g, b, _ := p.RGBA()

	and := color.RGBA{}
	and.R = c.rgba.R & uint8(r>>8)
	and.G = c.rgba.G & uint8(g>>8)
	and.B = c.rgba.B & uint8(b>>8)
	and.A = 255

	return and
}
