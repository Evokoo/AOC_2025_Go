package day04

import (
	"strings"

	"github.com/Evokoo/AOC_2025_Go/utils"
)

// ========================
// PART I
// ========================
func I(rolls Rolls) int {
	count := 0
	for roll := range rolls {
		if roll.IsAccessible(rolls) {
			count++
		}
	}
	return count
}

// ========================
// PART II
// ========================

func II(rolls Rolls) int {
	removed := 0

	for {
		toRemove := make(utils.Set[Roll])
		for roll := range rolls {
			if roll.IsAccessible(rolls) {
				toRemove.Add(roll)
			}
		}

		if len(toRemove) == 0 {
			return removed
		} else {
			removed += len(toRemove)
		}

		for roll := range toRemove {
			rolls.Remove(roll)
		}
	}
}

// ========================
// ROLLS
// ========================
type Roll [2]int
type Rolls = utils.Set[Roll]

var DIRECTIONS = [][2]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}

func (r Roll) IsAccessible(rolls Rolls) bool {
	neighbours := 0

	for _, dir := range DIRECTIONS {
		neighbour := [2]int{r[0] + dir[0], r[1] + dir[1]}
		if rolls.Has(neighbour) {
			neighbours++
		}

		if neighbours == 4 {
			return false
		}
	}

	return true
}

// ========================
// PARSER
// ========================

func ParseInput(file string) Rolls {
	data := utils.ReadFile(file)
	rolls := make(Rolls)

	for y, row := range strings.Split(data, "\n") {
		for x, col := range row {
			if col == '@' {
				rolls.Add(Roll{x, y})
			}
		}
	}

	return rolls
}
