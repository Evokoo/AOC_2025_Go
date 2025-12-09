package day09

import (
	"strings"

	"github.com/Evokoo/AOC_2025_Go/utils"
)

// ========================
// TILE
// ========================
type Tile [2]int

func (t Tile) Area(corner *Tile) int {
	w := utils.Abs(t[0]-(*corner)[0]) + 1
	h := utils.Abs(t[1]-(*corner)[1]) + 1
	return w * h
}

// ========================
// PART I
// ========================

func I(tiles []*Tile) int {
	length := len(tiles)
	largest := 0

	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			largest = max(largest, tiles[i].Area(tiles[j]))
		}
	}

	return largest
}

// ========================
// PARSER
// ========================

func ParseInput(file string) []*Tile {
	data := utils.ReadFile(file)
	lines := strings.Split(data, "\n")

	tiles := make([]*Tile, len(lines))
	for i, line := range lines {
		xy := utils.MatchInts(line)
		tiles[i] = &Tile{xy[0], xy[1]}
	}

	return tiles
}
