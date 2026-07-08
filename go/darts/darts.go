package darts

import "math"

func Score(x, y float64) int {
	distance := math.Sqrt(x*x + y*y)
	switch {
	case distance >= 0.0 && distance <= 1.0:
		return 10
	case distance > 1.0 && distance <= 5.0:
		return 5
	case distance > 5.0 && distance <= 10.0:
		return 1
	default:
		return 0

	}
}
