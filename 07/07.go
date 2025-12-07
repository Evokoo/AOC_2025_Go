package day07

func Solve(file string, part int) int {
	grid := ParseInput(file)
	splitters, paths := TraceBeam(grid)

	switch part {
	case 1:
		return splitters
	case 2:
		return CountTimelines(grid.start, paths, make(map[[2]int]int))
	}
	return 0
}
