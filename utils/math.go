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
