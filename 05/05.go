package day05

func Solve(file string, part int) int {
	ranges, ingridents := ParseInput(file)

	switch part {
	case 1:
		return I(ranges, ingridents)
	case 2:
		return II(ranges)
	}
	return 0
}
