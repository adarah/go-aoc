package day_02

import (
	"strconv"
	"strings"
)

type Solution struct{}

func (s *Solution) PartOne(input string) (string, error) {
	lines := strings.Split(input, "\n")
	h := 0
	v := 0
	for _, l := range lines {
		words := strings.Split(l, " ")
		dist, err := strconv.Atoi(words[1])
		if err != nil {
			return "", err
		}
		if l[0] == 'f' {
			h += dist
		} else if l[0] == 'd' {
			v += dist
		} else {
			v -= dist
		}
	}

	return strconv.Itoa(h * v), nil
}

func (s *Solution) PartTwo(input string) (string, error) {
	var left, right int
	aim := 0
	h := 0
	v := 0
	lineId := input[0]
	for i := 0; i < len(input); i++ {
		c := input[i]
		if c != ' ' {
			continue
		}
		i++
		left = i
		for i < len(input) && input[i] != '\n' {
			i++
		}
		right = i

		dist, err := strconv.Atoi(input[left:right])
		if err != nil {
			return "", err
		}
		if lineId == 'f' {
			h += dist
			v += aim * dist
		} else if lineId == 'u' {
			aim -= dist
		} else {
			aim += dist
		}
		if i+1 < len(input) {
			lineId = input[i+1]
		}
	}
	return strconv.Itoa(h * v), nil
}
