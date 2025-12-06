package day06

func Solve(file string, part int) int {
	switch part {
	case 1:
		return GetTotal(ParseLTR(file))
	case 2:
		return GetTotal(ParseTTB(file))
	}
	return 0
}
