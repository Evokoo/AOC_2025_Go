package day08

import (
	"strings"
)

func Solve(file string, part int) int {
	junctions := ParseInput(file)
	isExample := false

	if strings.Contains(file, "example") {
		isExample = true
	}

	return I(junctions, part, isExample)
}
