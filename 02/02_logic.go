package day02

import (
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2025_Go/utils"
)

// ========================
// PART I && PART II
// ========================
func I(pairs [][2]int, part int) int {
	sum := 0

	for _, pair := range pairs {
		start, end := pair[0], pair[1]

		for i := start; i <= end; i++ {
			s := strconv.Itoa(i)
			l := len(s)

			if part == 1 && s[:l/2] == s[l/2:] {
				sum += i
			}
			if part == 2 && HasPattern(s, l) {
				sum += i
			}
		}
	}

	return sum
}

func HasPattern(s string, l int) bool {
blocks:
	for size := 1; size <= l/2; size++ {
		if l%size != 0 {
			continue
		}

		block := s[:size]
		for i := size; i <= l-size; i += size {
			nextBlock := s[i : i+size]

			if nextBlock != block {
				continue blocks
			}
		}
		return true
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
