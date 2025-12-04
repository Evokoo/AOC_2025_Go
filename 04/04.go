package day04

func Solve(file string, part int) int {
	rolls := ParseInput(file)

	switch part {
	case 1:
		return I(rolls)
	case 2:
		return II(rolls)
	}
	return 0
}
