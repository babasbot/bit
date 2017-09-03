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
