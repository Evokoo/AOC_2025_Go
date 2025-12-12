package day12

func Solve(file string, part int) int {
	shapes, regions := ParseInput(file)
	return I(shapes, regions)
}
