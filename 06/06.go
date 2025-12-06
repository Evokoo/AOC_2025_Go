package day06

func Solve(file string, part int) int {
	switch part {
	case 1:
		return I(ParseInput(file))
	case 2:
		return I(ParseInputII(file))
	}
	return 0
}
