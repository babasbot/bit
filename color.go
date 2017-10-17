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
