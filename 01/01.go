package day01

func Solve(file string, part int) int {
	turns := ParseInput(file)

	switch part {
	case 1:
		return I(turns)
	case 2:
		return II(turns)

	}

	return 0
}
