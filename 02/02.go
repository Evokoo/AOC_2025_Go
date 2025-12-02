package day02

func Solve(file string, part int) int {
	pairs := ParseInput(file)

	switch part {
	case 1:
		return I(pairs)
	}
	return 0
}
