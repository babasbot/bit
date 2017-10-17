package bit

import (
	"image/color"
)

func (c *Color) Gray() color.Gray {
	luma := uint8(c.Luma())

	return color.Gray{luma}
}

func (c *Color) Luma() float32 {
	r, g, b, _ := c.rgba.RGBA()

	return 0.299*float32(r>>8) + 0.587*float32(g>>8) + 0.114*float32(b>>8)
}
