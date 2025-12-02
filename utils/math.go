package utils

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
