package day12

import (
	"strings"

	"github.com/Evokoo/AOC_2025_Go/utils"
)

// ========================
// SHAPES
// ========================
type Shapes map[int]int

// ========================
// REGIONS
// ========================

type Region struct {
	width, height int
	shapes        []int
}

// ========================
// PART I
// ========================

func I(shapes Shapes, regions []Region) int {

	count := 0

	for _, region := range regions {
		available := region.height * region.width
		required := 0

		for i, c := range region.shapes {
			required += shapes[i] * c
		}

		if available >= required {
			count++
		}
	}

	return count
}

// ========================
// PARSER
// ========================

func ParseInput(file string) (Shapes, []Region) {
	data := utils.ReadFile(file)
	sections := strings.Split(data, "\n\n")

	shapes := make(Shapes)
	for _, shape := range sections[:len(sections)-1] {
		var id int
		var area int

		for i, line := range strings.Split(shape, "\n") {
			if i == 0 {
				id = utils.MatchInts(line)[0]
			} else {
				for _, r := range line {
					if r == '#' {
						area++
					}
				}
			}
		}
		shapes[id] = area
	}

	regions := make([]Region, 0)
	for region := range strings.SplitSeq(sections[len(sections)-1], "\n") {
		values := utils.MatchInts(region)

		regions = append(regions, Region{
			width:  values[0],
			height: values[1],
			shapes: values[2:],
		})
	}

	return shapes, regions
}
