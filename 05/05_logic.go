package day05

import (
	"sort"
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2025_Go/utils"
)

// ========================
// RANGE
// ========================

type Range [2]int

func (r Range) InRange(n int) bool {
	return n >= r[0] && n <= r[1]
}
func (r Range) Size() int {
	return r[1] - r[0] + 1
}

// ========================
// PART I
// ========================

func I(ranges []Range, ingridents []int) int {
	count := 0

	for _, ingrident := range ingridents {
		for _, r := range ranges {
			if r.InRange(ingrident) {
				count++
				break
			}
		}
	}
	return count
}

// ========================
// PART II
// ========================

func II(ranges []Range) int {
	count := 0

	for _, r := range ranges {
		count += r.Size()
	}
	return count
}

// ========================
// PARSER
// ========================

func ParseInput(file string) ([]Range, []int) {
	data := utils.ReadFile(file)

	var ranges []Range
	var ingridents []int

	for i, section := range strings.Split(data, "\n\n") {
		if i == 0 {
			for _, pair := range strings.Split(section, "\n") {
				values := strings.Split(pair, "-")
				a, _ := strconv.Atoi(values[0])
				b, _ := strconv.Atoi(values[1])
				ranges = append(ranges, Range{a, b})
			}
		} else {
			for _, ingrident := range strings.Split(section, "\n") {
				n, _ := strconv.Atoi(ingrident)
				ingridents = append(ingridents, n)
			}
		}
	}

	return MergeRanges(ranges), ingridents
}
func MergeRanges(ranges []Range) []Range {
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	merged := []Range{ranges[0]}
	for _, cur := range ranges[1:] {
		last := &merged[len(merged)-1]

		if cur[0] <= (*last)[1] {
			if cur[1] > (*last)[1] {
				(*last)[1] = cur[1]
			}
		} else {
			merged = append(merged, cur)
		}
	}

	return merged
}
