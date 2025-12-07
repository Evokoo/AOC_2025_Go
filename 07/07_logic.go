package day07

import (
	"strings"

	"github.com/Evokoo/AOC_2025_Go/utils"
)

// ========================
// GRID
// ========================
type Grid struct {
	start      [2]int
	splits     utils.Set[[2]int]
	rows, cols int
}

func (g Grid) IsSplitter(x, y int) bool {
	return g.splits.Has([2]int{x, y})
}

// ========================
// PART I
// ========================

type BeamPaths map[[2]int][][2]int

func TraceBeam(grid Grid) (int, BeamPaths) {
	queue := make(utils.Queue[[2]int], 0)
	queue.Push(grid.start)

	seen := make(utils.Set[[2]int])
	count := make(utils.Set[[2]int])

	paths := make(BeamPaths)

	for !queue.IsEmpty() {
		cur := queue.Pop()
		source := cur

		if seen.Has(cur) {
			continue
		} else {
			seen.Add(cur)
		}

		for cur[1] < grid.rows {
			if grid.IsSplitter(cur[0], cur[1]) {

				left := [2]int{cur[0] - 1, cur[1]}
				right := [2]int{cur[0] + 1, cur[1]}

				queue.Push(left)
				queue.Push(right)

				paths[source] = [][2]int{left, right}
				count.Add(cur)

				break
			}
			cur[1]++
		}

		if cur[1] >= grid.rows {
			paths[source] = nil
		}
	}

	return len(count), paths
}

// ========================
// PART II
// ========================

func CountTimelines(source [2]int, paths BeamPaths, memo map[[2]int]int) int {
	if val, found := memo[source]; found {
		return val
	}

	children := paths[source]

	if len(children) == 0 {
		memo[source] = 1
		return 1
	} else {
		total := 0

		for _, child := range children {
			total += CountTimelines(child, paths, memo)
		}

		memo[source] = total
		return total
	}
}

// ========================
// PARSER
// ========================
func ParseInput(file string) Grid {
	data := utils.ReadFile(file)
	lines := strings.Split(data, "\n")

	grid := Grid{
		start:  [2]int{-1, -1},
		splits: make(utils.Set[[2]int]),
		rows:   len(lines),
		cols:   len(lines[0]),
	}

	for y := range grid.rows {
		for x := range grid.cols {
			switch lines[y][x] {
			case 'S':
				grid.start = [2]int{x, y}
			case '^':
				grid.splits.Add([2]int{x, y})
			}
		}
	}

	return grid
}
