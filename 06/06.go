package day06

func Solve(file string, part int) int {
	columns := ParseInput(file)

	switch part {
	case 1:
		return I(columns)
	case 2:
		return II(columns)
	}
	return 0
}
