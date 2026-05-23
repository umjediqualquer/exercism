package darts

import "math"

func Score(x, y float64) int {
	radius := math.Hypot(x, y)

	score := 0
	if radius <= 1 {
		score = 10
	} else if radius <= 5 {
		score = 5
	} else if radius <= 10 {
		score = 1
	}

	return score
}
