package bit

import "math"

type Point struct {
	x int
	y int
}

func truncate(n float64) float64 {
	n = math.Max(n, 0)
	n = math.Min(n, 255)

	return n
}

func normalize(n int32) uint8 {
	if n > 255 {
		return 255
	} else if n < 0 {
		return 0
	}

	return uint8(n)
}
