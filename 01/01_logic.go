package day01

import (
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2025_Go/utils"
)

const DIAL_MAX = 100
const DIAL_MIN = 0

// ========================
// PART I
// ========================
func I(turns []int) int {
	current := 50
	count := 0

	for _, turn := range turns {
		current = (current + turn + DIAL_MAX) % DIAL_MAX
		if current == DIAL_MIN {
			count++
		}
	}

	return count
}

// ========================
// PART II
// ========================
func II(turns []int) int {
	current := 50
	count := 0

	for _, turn := range turns {
		count += utils.Abs(turn) / DIAL_MAX
		step := 1

		if turn < 0 {
			step = -1
		}

		remaining := utils.Abs(turn) % DIAL_MAX
		for range remaining {
			current = (current + step + DIAL_MAX) % DIAL_MAX
			if current == DIAL_MIN {
				count++
			}
		}
	}

	return count
}

// ========================
// PARSER
// ========================

func ParseInput(file string) []int {
	data := utils.ReadFile(file)
	turns := make([]int, 0)

	for line := range strings.SplitSeq(data, "\n") {
		n, _ := strconv.Atoi(line[1:])
		if line[0] == 'L' {
			n *= -1
		}
		turns = append(turns, n)
	}

	return turns
}
