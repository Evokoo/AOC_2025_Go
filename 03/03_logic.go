package day03

import (
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2025_Go/utils"
)

// ========================
// PART I && PART II
// ========================
func GetOutputJoltage(batteries [][]rune, part int) int {
	cellCount := 2

	if part == 2 {
		cellCount = 12
	}

	total := 0
	for _, battery := range batteries {
		total += GetJoltage(battery, cellCount)
	}

	return total
}

func GetJoltage(battery []rune, cellCount int) int {
	digits := make([]rune, 0, cellCount)
	index := 0

	for len(digits) < cellCount {
		remaining := cellCount - len(digits)
		searchEnd := len(battery) - remaining

		best := '0'
		bestIndex := index

		for j := index; j <= searchEnd; j++ {
			if battery[j] > best {
				best = battery[j]
				bestIndex = j
			}
		}

		digits = append(digits, best)
		index = bestIndex + 1
	}

	n, _ := strconv.Atoi(string(digits))
	return n
}

// ========================
// PARSER
// ========================

func ParseInput(file string) [][]rune {
	data := utils.ReadFile(file)
	batteries := make([][]rune, 0)

	for battery := range strings.SplitSeq(data, "\n") {
		batteries = append(batteries, []rune(battery))
	}

	return batteries
}
