package day06

import (
	"strconv"
	"strings"

	"github.com/Evokoo/AOC_2025_Go/utils"
)

// ========================
// COLUMN
// ========================

type Column struct {
	values   []int
	operator string
}

var INITIAL_VALUE = map[string]int{"+": 0, "*": 1}

func (c Column) Solve() int {
	total := INITIAL_VALUE[c.operator]

	for _, value := range c.values {
		switch c.operator {
		case "+":
			total += value
		case "*":
			total *= value
		default:
			panic("Invalid operator")
		}
	}

	return total
}

// ========================
// PART I
// ========================

func GetTotal(columns []Column) int {
	sum := 0
	for _, column := range columns {
		sum += column.Solve()
	}
	return sum
}

// ========================
// PARSER
// ========================

func ParseLTR(file string) []Column {
	data := utils.ReadFile(file)
	lines := strings.Split(data, "\n")
	rows := len(lines)

	var columns []Column

	for i, row := range lines[:rows-1] {
		values := utils.MatchInts(row)

		if i == 0 {
			columns = make([]Column, len(values))
		}
		for j, value := range values {
			columns[j].values = append(columns[j].values, value)
		}
	}

	for i, operator := range utils.QuickMatch(lines[rows-1], `\S+`) {
		columns[i].operator = operator
	}

	return columns
}

func ParseTTB(file string) []Column {
	data := utils.ReadFile(file)
	lines := strings.Split(data, "\n")

	rows := len(lines)
	cols := len(lines[0])

	var columns []Column
	var current Column
	var digits strings.Builder

	addColumn := func() {
		columns = append(columns, current)
		current = Column{}
	}

	for i := range cols {
		for j := range rows {
			char := lines[j][i]

			switch char {
			case '+', '*':
				current.operator = string(char)
			case ' ':
				continue
			default:
				digits.WriteByte(char)
			}
		}

		s := digits.String()

		switch s {
		case "":
			addColumn()
		default:
			n, _ := strconv.Atoi(s)
			current.values = append(current.values, n)
		}

		digits.Reset()
	}

	//Add final column
	addColumn()

	return columns
}
