package day09

func Solve(file string, part int) int {
	tiles := ParseInput(file)

	switch part {
	case 1:
		return I(tiles)
	}

	return 0
}
