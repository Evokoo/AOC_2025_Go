package day11

func Solve(file string, part int) int {
	server := ParseInput(file)

	switch part {
	case 1:
		return I(server)
	case 2:
		return II(server)
	}
	return 0
}
