package equation

import "math"

// SolveLinear solves linear equation in the format of ax + b = 0
func SolveLinear(a, b float64) float64 {
	return -b / a
}

// SolveQuadratic solves quadratic equation in the format of ax^2 + bx + c = 0
func SolveQuadratic(a, b, c float64) (float64, float64) {
	s := math.Sqrt(math.Pow(b, 2) - 4*a*c)
	solve1 := (-b + s) / (2 * a)
	solve2 := (-b - s) / (2 * a)
	return solve1, solve2
}
