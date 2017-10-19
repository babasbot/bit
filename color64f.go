package bit

import (
	"math"
)

type RGBAf64 struct {
	R, G, B, A float64
}

func (c RGBAf64) RGBA() (r, g, b, a uint32) {
	r = Normalizef64(c.R) << 8
	g = Normalizef64(c.G) << 8
	b = Normalizef64(c.B) << 8
	a = Normalizef64(c.A) << 8

	return
}

func Normalizef64(x float64) uint32 {
	return uint32(math.Min(math.Max(x, 0), 255))
}
