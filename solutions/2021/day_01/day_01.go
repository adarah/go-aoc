package day_01

import (
	"math"
	"strconv"
)

type Solution struct{}

func (s *Solution) PartOne(input string) (string, error) {
	depths := make(chan int)
	go func() {
		boundary := 0
		for i, c := range input {
			if c != '\n' {
				continue
			}
			depth, err := strconv.Atoi(input[boundary:i])
			if err != nil {
				panic(err)
			}
			depths <- depth
			boundary = i + 1
		}
		close(depths)
	}()

	increased := 0
	lastDepth := math.MaxInt64
	for depth := range depths {
		if depth > lastDepth {
			increased++
		}
		lastDepth = depth
	}
	return strconv.Itoa(increased), nil
}

func (s *Solution) PartTwo(input string) (string, error) {
	return doPartTwo(input, 3)
}

func doPartTwo(input string, windowSize int) (string, error) {
	measurements, err := parse(input)
	if err != nil {
		return "", err
	}

	increased := 0
	for i := windowSize; i < len(measurements); i++ {
		if measurements[i-windowSize] < measurements[i] {
			increased++
		}
	}
	return strconv.Itoa(increased), nil
}

func parse(text string) ([]int, error) {
	var measurements []int
	boundary := 0
	for i, c := range text {
		if c != '\n' {
			continue
		}
		depth, err := strconv.Atoi(text[boundary:i])
		if err != nil {
			return nil, err
		}
		measurements = append(measurements, depth)
		boundary = i + 1
	}
	return measurements, nil
}
