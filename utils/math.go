package utils

import "math"

// ========================
// ABS
// ========================
func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

// ========================
// CEIL
// ========================
func Ceil(a, b int) int {
	return (a + b - 1) / b
}

// ========================
// POWER
// ========================
func Pow(base, exp int) int {
	result := 1
	for exp > 0 {
		if exp&1 == 1 {
			result *= base
		}
		base *= base
		exp >>= 1
	}
	return result
}

// ========================
// DISTANCE
// ========================

func EuclideanDistance3D(a, b [3]int) float64 {
	dx := float64(b[0] - a[0])
	dy := float64(b[1] - a[1])
	dz := float64(b[2] - a[2])

	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}
