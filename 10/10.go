package day10

func Solve(file string, part int) int {
	machines := ParseInput(file)

	switch part {
	case 1:
		return I(machines)
	case 2:
		return II(machines)
	}

	return 0
}
