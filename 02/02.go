package day02

func Solve(file string, part int) int {
	pairs := ParseInput(file)

	switch part {
	case 1:
		return I(pairs)
	case 2:
		return II(pairs)
	}
	return 0
}
