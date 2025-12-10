package day10

func Solve(file string, part int) int {
	machines := ParseInput(file)

	switch part {
	case 1:
		return I(machines)
	}

	return 0
}
