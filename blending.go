package bit

import (
	"image"
	"image/color"
)

func (img *Image) Blending(blend Image, alpha float32) Image {
	bounds := img.Bounds()
	newimg := image.NewRGBA(bounds)

	for y := 0; y < bounds.Max.Y; y++ {
		for x := 0; x < bounds.Max.X; x++ {
			col1 := img.At(x, y)
			col2 := blend.At(x, y)

			newimg.Set(x, y, Blending(col1, col2, alpha))
		}
	}

	return Image{newimg}
}

func Blending(col1, col2 color.Color, alpha float32) color.Color {
	r1, g1, b1, a1 := col1.RGBA()
	r2, g2, b2, a2 := col2.RGBA()

	result := color.RGBA{}
	result.R = uint8(alphablending(r1, r2, alpha) >> 8)
	result.G = uint8(alphablending(g1, g2, alpha) >> 8)
	result.B = uint8(alphablending(b1, b2, alpha) >> 8)
	result.A = uint8(alphablending(a1, a2, alpha) >> 8)

	return result
}

func alphablending(a, b uint32, alpha float32) uint32 {
	result := float32(a)*alpha + float32(b)*(1.0-alpha)

	return uint32(result)
}
