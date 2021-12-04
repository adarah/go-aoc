package lib

import "os"

func ReadFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return string(data)
}

type Solution interface {
	PartOne(input string) (string, error)
	PartTwo(input string) (string, error)
}
