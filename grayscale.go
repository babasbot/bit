package bit

func (c *Color) Luma() float32 {
	r, g, b, _ := c.rgba.RGBA()

	return 0.299*float32(r>>8) + 0.587*float32(g>>8) + 0.114*float32(b>>8)
}
