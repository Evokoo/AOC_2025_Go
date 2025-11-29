package utils

import (
	"regexp"
	"strconv"
)

// ========================
// STRINGS
// ========================
func QuickMatch(str, pattern string) []string {
	re := regexp.MustCompile(pattern)
	return re.FindAllString(str, -1)
}

func MatchInts(str string) []int {
	re := regexp.MustCompile(`-?\d+`)
	matches := re.FindAllString(str, -1)

	numbers := make([]int, len(matches))
	for i, str := range matches {
		numbers[i], _ = strconv.Atoi(str)
	}

	return numbers
}
