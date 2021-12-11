package day_03

import (
	"github.com/adarah/go-aoc/lib"
	"sort"
	"strings"
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

func TestPartTwo(t *testing.T) {
	input = lib.ReadFile("example.txt")
	lines := strings.Split(input, "\n")
	sort.Strings(lines)
	t.Run("case=finds-oxygen", func(t *testing.T) {
		actual := findOxygen(0, len(lines), 0, lines)
		expected := "10111"
		if actual != expected {
			t.Errorf("expected: %s, actual: %s\n", expected, actual)
		}
	})

	t.Run("case=finds-co2", func(t *testing.T) {
		actual := findCO2(0, len(lines), 0, lines)
		expected := "01010"
		if actual != expected {
			t.Errorf("expected: %s, actual: %s\n", expected, actual)
		}
	})
}
