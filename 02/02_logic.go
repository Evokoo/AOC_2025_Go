package day02

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2025_Go/utils"
)

// ========================
// PART I
// ========================
func I(pairs [][2]int) int {
	sum := 0

	for _, pair := range pairs {
		start, end := pair[0], pair[1]

		for i := start; i <= end; i++ {
			if Repeats(i) {
				sum += i
			}
		}
	}

	return sum
}

func Repeats(n int) bool {
	s := strconv.Itoa(n)
	l := len(s)

	if l%2 != 0 {
		return false
	}

	for i := 0; i < l/2; i++ {
		if s[i] != s[i+l/2] {
			return false
		}
	}

	return true
}

// ========================
// PART II
// ========================
func II(pairs [][2]int) int {
	sum := 0

	for _, pair := range pairs {
		start, end := pair[0], pair[1]

		for i := start; i <= end; i++ {
			if ContainsPattern(i) {
				sum += i
			}
		}
	}

	return sum
}

func ContainsPattern(n int) bool {
	s := strconv.Itoa(n)
	l := len(s)

	for i := 0; i <= l/2; i++ {
		re := regexp.MustCompile(fmt.Sprintf("^(%s)+$", s[0:i]))
		if re.MatchString(s) {
			return true
		}
	}
	return false
}

// ========================
// PARSER
// ========================

func ParseInput(file string) [][2]int {
	data := utils.ReadFile(file)
	pairs := make([][2]int, 0)

	for line := range strings.SplitSeq(data, ",") {
		numbers := strings.Split(line, "-")
		a, _ := strconv.Atoi(numbers[0])
		b, _ := strconv.Atoi(numbers[1])

		pairs = append(pairs, [2]int{a, b})
	}
	return pairs
}
