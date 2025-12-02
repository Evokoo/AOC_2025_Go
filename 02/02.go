package day02

func Solve(file string, part int) int {
	pairs := ParseInput(file)

	switch part {
	case 1:
		return I(pairs, part)
	case 2:
		return I(pairs, part)
	}
	return 0
}
