package day11_test

import (
	"fmt"

	. "github.com/Evokoo/AOC_2025_Go/11"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Test struct {
	file   string
	part   int
	target int
}

var tests = []Test{
	{part: 1, file: "inputs/example_I.txt", target: 5},
	{part: 1, file: "inputs/input.txt", target: 670},
	{part: 2, file: "inputs/example_II.txt", target: 2},
	{part: 2, file: "inputs/input.txt", target: 332052564714990},
}

var _ = Describe("AOC 2025 - Day 11", func() {
	for _, test := range tests {
		msg := fmt.Sprintf("Testing Part %d with %s", test.part, test.file)
		It(msg, func() {
			result := Solve(test.file, test.part)
			Expect(result).To(Equal((test.target)))
		})
	}
})
