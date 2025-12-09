package day09

import (
	"sort"
	"strings"

	"github.com/Evokoo/AOC_2025_Go/utils"
)

// ========================
// CONSTANTS
// ========================
const OFFSET = 1

var DIRECTIONS = [][2]int{
	{1, 0}, {-1, 0}, {0, 1}, {0, -1},
}

// ========================
// TILE
// ========================
type Tile [2]int

func (t Tile) Area(corner *Tile) int {
	w := utils.Abs(t[0]-(*corner)[0]) + 1
	h := utils.Abs(t[1]-(*corner)[1]) + 1
	return w * h
}
func (t Tile) CompressTile(lookup Lookup, offset int) (int, int) {
	cx := lookup.CompressX(t[0]) + offset
	cy := lookup.CompressY(t[1]) + offset
	return cx, cy
}

// ========================
// LOOKUP
// ========================

type Lookup struct {
	cx, cy map[int]int
}

func GenerateLookup(tiles []*Tile) Lookup {
	// Gather unique, remove duplicates and sort
	xSet := make(utils.Set[int])
	ySet := make(utils.Set[int])

	for _, tile := range tiles {
		xSet.Add(tile[0])
		ySet.Add(tile[1])
	}

	xUnique := make([]int, 0, len(xSet))
	yUnique := make([]int, 0, len(ySet))

	for x := range xSet {
		xUnique = append(xUnique, x)
	}
	for y := range ySet {
		yUnique = append(yUnique, y)
	}

	sort.Ints(xUnique)
	sort.Ints(yUnique)

	// Lookups
	cx := make(map[int]int, len(xUnique))
	cy := make(map[int]int, len(yUnique))

	for i, x := range xUnique {
		cx[x] = i
	}
	for i, y := range yUnique {
		cy[y] = i
	}

	return Lookup{
		cx: cx,
		cy: cy,
	}
}

func (l Lookup) CompressX(x int) int {
	return l.cx[x]
}
func (l Lookup) CompressY(y int) int {
	return l.cy[y]
}
func (l Lookup) BuildGrid(points []*Tile, offset int) [][]string {
	//Compress and offset points
	cTiles := make([]Tile, len(points))
	rows, cols := 0, 0

	for i, tile := range points {
		cx, cy := tile.CompressTile(l, 0)
		rows = max(rows, cy+3)
		cols = max(cols, cx+3)

		cTiles[i] = Tile{cx + 1, cy + 1}
	}

	// Add first tile on end to complete loop
	cTiles = append(cTiles, cTiles[0])

	// Empty Grid
	grid := make([][]string, rows)
	for y := range grid {
		grid[y] = make([]string, cols)
		for x := range cols {
			grid[y][x] = "#"
		}
	}

	//Trace Outline
	outline := make(utils.Set[[2]int])

	for i := 1; i < len(cTiles); i++ {
		a := cTiles[i-1]
		b := cTiles[i]

		if a[1] == b[1] {
			y := a[1]
			x1, x2 := a[0], b[0]
			if x1 > x2 {
				x1, x2 = x2, x1
			}
			for x := x1; x <= x2; x++ {
				outline.Add([2]int{x, y})
			}
		} else if a[0] == b[0] {
			x := a[0]
			y1, y2 := a[1], b[1]
			if y1 > y2 {
				y1, y2 = y2, y1
			}
			for y := y1; y <= y2; y++ {
				outline.Add([2]int{x, y})
			}
		} else {
			panic("non-axis-aligned segment not supported")
		}
	}

	//Flood fill

	queue := make(utils.Queue[[2]int], 0)
	queue.Push([2]int{0, 0})

	seen := make(utils.Set[[2]int])
	seen.Add([2]int{0, 0})

	grid[0][0] = "."

	for !queue.IsEmpty() {
		c := queue.Pop()

		for _, d := range DIRECTIONS {
			nx, ny := c[0]+d[0], c[1]+d[1]

			if nx >= 0 && nx < cols && ny >= 0 && ny < rows {
				next := [2]int{nx, ny}

				if !outline.Has(next) && !seen.Has(next) {
					grid[ny][nx] = "."
					seen.Add(next)
					queue.Push(next)
				}
			}
		}
	}

	return grid
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
// PART II
// ========================

func II(tiles []*Tile) int {
	lookup := GenerateLookup(tiles)
	grid := lookup.BuildGrid(tiles, OFFSET)
	length := len(tiles)

	largest := 0

	for i := 0; i < length; i++ {
		cx1, cy1 := tiles[i].CompressTile(lookup, OFFSET)

		for j := i + 1; j < length; j++ {
			cx2, cy2 := tiles[j].CompressTile(lookup, OFFSET)

			// Determine bounding rectangle
			minX, maxX := min(cx1, cx2), max(cx1, cx2)
			minY, maxY := min(cy1, cy2), max(cy1, cy2)

			// Check all cells inside the rectangle
			inside := true
			for y := minY; y <= maxY && inside; y++ {
				for x := minX; x <= maxX; x++ {
					if grid[y][x] != "#" {
						inside = false
						break
					}
				}
			}

			if inside {
				largest = max(largest, tiles[i].Area(tiles[j]))
			}
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
