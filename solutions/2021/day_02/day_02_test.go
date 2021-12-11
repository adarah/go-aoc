package day_02

import (
	"github.com/adarah/go-aoc/lib"
	"testing"
)

var solution Solution
var input string

func init() {
	input = lib.ReadFile("input.txt")
}

func BenchmarkPartOne(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = solution.PartOne(input)
	}
}

func BenchmarkPartTwo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_, _ = solution.PartTwo(input)
	}
}
