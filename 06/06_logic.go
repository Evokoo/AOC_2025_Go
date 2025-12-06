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

func (c Column) Solve() int {
	total := 0

	if c.operator == "*" {
		total = 1
	}

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

func I(columns []Column) int {
	sum := 0

	for _, column := range columns {
		sum += column.Solve()
	}

	return sum
}

// ========================
// PART II
// ========================
func II(columns []Column) int {
	sum := 0

	return sum
}

// ========================
// PARSER
// ========================

func ParseInput(file string) []Column {
	data := utils.ReadFile(file)
	rows := strings.Split(data, "\n")
	length := len(rows)

	var columns []Column

	for i, row := range rows[:length-1] {
		values := utils.MatchInts(row)

		if i == 0 {
			columns = make([]Column, len(values))
		}
		for j, value := range values {
			columns[j].values = append(columns[j].values, value)
		}
	}

	for i, operator := range utils.QuickMatch(rows[length-1], `\S+`) {
		columns[i].operator = operator
	}

	return columns
}

func ParseInputII(file string) []Column {
	data := utils.ReadFile(file)
	rows := strings.Split(data, "\n")
	length := len(rows)

	var columns []Column
	var current Column
	var digits strings.Builder

	for i := range rows[0] {
		for j := range length - 1 {
			digits.WriteByte(rows[j][i])
		}

		s := strings.TrimSpace(digits.String())
		if s == "" {
			columns = append(columns, current)
			current = Column{}
		} else {
			n, _ := strconv.Atoi(s)
			current.values = append(current.values, n)
		}

		digits.Reset()

		if i == len(rows[0])-1 {
			columns = append(columns, current)
		}
	}

	for i, operator := range utils.QuickMatch(rows[length-1], `\S+`) {
		columns[i].operator = operator
	}

	return columns
}
