package day11

func Solve(file string, part int) int {
	server := ParseInput(file)

	switch part {
	case 1:
		return I(server)
	}
	return 0
}
