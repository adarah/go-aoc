package day_03

import (
	"github.com/yourbasic/radix"
	"strconv"
	"strings"
)

type Solution struct{}

func (s *Solution) PartOne(input string) (string, error) {
	var seqLen int
	for i, c := range input {
		if c != '\n' {
			continue
		}
		seqLen = i
		break
	}
	counts := make([]int, seqLen)
	idx := 0

	for _, c := range input {
		if c == '\n' {
			idx = 0
			continue
		}
		if c == '1' {
			counts[idx]++
		} else {
			counts[idx]--
		}
		idx++
	}

	gamma := 0
	epsilon := 0
	for _, c := range counts {
		gamma <<= 1
		epsilon <<= 1
		if c > 0 {
			gamma |= 1
		} else {
			epsilon |= 1
		}
	}
	return strconv.Itoa(gamma * epsilon), nil
}

func (s *Solution) PartTwo(input string) (string, error) {
	lines := strings.Split(input, "\n")
	radix.Sort(lines)
	oxygen := findOxygen(0, len(lines), 0, lines)
	o2, err := strconv.ParseInt(oxygen, 2, 64)
	if err != nil {
		return "", err
	}
	carbonDioxide := findCO2(0, len(lines), 0, lines)
	co2, err := strconv.ParseInt(carbonDioxide, 2, 64)
	if err != nil {
		return "", err
	}

	return strconv.Itoa(int(o2 * co2)), nil
}

func findMetrics(start, end, bitIdx int, lines []string, checkFunc func(edge, middle int) bool) string {
	if bitIdx == len(lines[0]) || end-start <= 1 {
		return lines[start]
	}
	i := start
	for lines[i][bitIdx] == '0' {
		i++
	}
	edge := i

	middle := start + (end-start)/2
	if checkFunc(edge, middle) {
		return findMetrics(start, edge, bitIdx+1, lines, checkFunc)
	} else {
		return findMetrics(edge, end, bitIdx+1, lines, checkFunc)
	}
}

func findOxygen(start, end, bitIdx int, lines []string) string {
	return findMetrics(start, end, bitIdx, lines, func(edge, middle int) bool { return edge > middle })
}

func findCO2(start, end, bitIdx int, lines []string) string {
	return findMetrics(start, end, bitIdx, lines, func(edge, middle int) bool { return edge <= middle })
}
