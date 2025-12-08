package day08

import (
	"sort"
	"strings"

	"github.com/Evokoo/AOC_2025_Go/utils"
)

// ========================
// PART I
// ========================

func ConnectBoxes(pairs []Link, count int) {

}

func I(junctions Junctions, part int, isExample bool) int {
	pairs := junctions.PairAndSortByDistance()
	size := len(pairs)
	lastConnection := 0

	if part == 1 {
		if isExample {
			size = 10
		} else {
			size = 1000
		}
	}

	//Connect Boxes
	for _, link := range pairs[:size] {
		aRoot := link.a.FindRoot()
		bRoot := link.b.FindRoot()

		if aRoot != bRoot {
			bRoot.parent = aRoot
			lastConnection = link.a.position[0] * link.b.position[0]
		}
	}

	switch part {
	case 1:
		return junctions.GetCircuitValue()
	case 2:
		return lastConnection
	}

	return 0
}

// ========================
// JUNCTIONS
// ========================

type Box struct {
	id       int
	position [3]int
	parent   *Box
}

func NewBox(id int) *Box {
	return &Box{
		id:       id,
		position: [3]int{},
		parent:   nil,
	}
}

func (b *Box) SharedRoot(other *Box) bool {
	return b.FindRoot() == other.FindRoot()
}
func (b *Box) FindRoot() *Box {
	if b.parent == nil {
		return b
	}
	b.parent = b.parent.FindRoot()
	return b.parent
}

type Junctions map[int]*Box
type Link struct {
	distance float64
	a, b     *Box
}

func (j Junctions) PairAndSortByDistance() []Link {
	links := make([]Link, 0)

	for i := range j {
		a := j[i]

		for k := i + 1; k < len(j); k++ {
			b := j[k]
			dist := utils.EuclideanDistance3D(a.position, b.position)
			links = append(links, Link{dist, a, b})

		}
	}

	sort.Slice(links, func(a, b int) bool {
		return links[a].distance < links[b].distance
	})

	return links
}
func (j Junctions) GetCircuitValue() int {
	circuitSizes := make([]int, len(j))
	for _, box := range j {
		rootID := box.FindRoot().id
		circuitSizes[rootID]++
	}

	sort.Slice(circuitSizes, func(i, j int) bool {
		return circuitSizes[i] > circuitSizes[j]
	})

	total := 1
	for _, size := range circuitSizes[:3] {
		total *= size
	}

	return total
}

// ========================
// PARSER
// ========================

func ParseInput(file string) Junctions {
	data := utils.ReadFile(file)
	lines := strings.Split(data, "\n")

	junctions := make(Junctions, len(lines))

	for id, xyz := range lines {
		box := NewBox(id)

		for index, value := range utils.MatchInts(xyz) {
			box.position[index] = value
		}

		junctions[id] = box
	}

	return junctions
}
