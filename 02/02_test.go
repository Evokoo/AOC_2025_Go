package day02_test

import (
	"fmt"

	. "github.com/Evokoo/AOC_2025_Go/02"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type Test struct {
	file   string
	part   int
	target int
}

var tests = []Test{
	{part: 1, file: "inputs/example_I.txt", target: 1227775554},
	{part: 1, file: "inputs/input.txt", target: 26255179562},
	{part: 2, file: "inputs/example_II.txt", target: 4174379265},
	{part: 2, file: "inputs/input.txt", target: 31680313976},
}

var _ = Describe("AOC 2025 - Day 02", func() {
	for _, test := range tests {
		msg := fmt.Sprintf("Testing Part %d with %s", test.part, test.file)
		It(msg, func() {
			result := Solve(test.file, test.part)
			Expect(result).To(Equal((test.target)))
		})
	}
})
