package bit

import (
	"image/color"
)

type Color struct {
	rgba color.RGBA
}

func (c *Color) RGBA() (r, g, b, a uint32) {
	r, g, b, a = c.rgba.RGBA()

	return
}

func NewRGBA(c color.Color) (new Color) {
	r, g, b, a := c.RGBA()

	new = Color{}
	new.rgba.R = uint8(r >> 8)
	new.rgba.G = uint8(g >> 8)
	new.rgba.B = uint8(b >> 8)
	new.rgba.A = uint8(a >> 8)

	return
}
