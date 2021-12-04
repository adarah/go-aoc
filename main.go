package main

//go:generate go run gen/gen_main.go

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
)

func init() {
	log.SetFlags(0)
}

func main() {
	var year uint
	var day uint
	var part uint

	currentYear := uint(time.Now().Year())
	app := &cli.App{
		Name:                 "go-aoc",
		Usage:                "my Advent of Code solutions",
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.UintFlag{
				Name:        "year",
				Aliases:     []string{"y"},
				Usage:       fmt.Sprintf("`YEAR` of the solution. Must be between 2015 and %d (inclusive).", currentYear),
				Value:       currentYear,
				DefaultText: fmt.Sprintf("%d", currentYear),
				Destination: &year,
			},
			&cli.UintFlag{
				Name:        "day",
				Aliases:     []string{"d"},
				Usage:       "`DAY` of the solution. Must be between 1 and 25 (inclusive).",
				Required:    true,
				Destination: &day,
			},
			&cli.UintFlag{
				Name:        "part",
				Aliases:     []string{"p"},
				Usage:       "`PART` of the solution. Must be either 1 or 2.",
				Required:    true,
				Destination: &part,
			},
		},
		Action: func(c *cli.Context) error {
			isValidYear := 2015 < year && year <= currentYear
			isValidDay := 1 <= day && day <= 25
			isValidPart := part == 1 || part == 2
			if !isValidYear {
				return fmt.Errorf("year must be in range: 2015 - %d", currentYear)
			}
			if !isValidDay {
				return fmt.Errorf("day must be in range: 1 - 25")
			}
			if !isValidPart {
				return fmt.Errorf("part must be either 1 or 2")
			}

			solution, err := getImplementation(year, day)
			if err != nil {
				return err
			}
			path := fmt.Sprintf("solutions/%d/day_%02d/input.txt", year, day)
			bytes, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			input := string(bytes)
			var result string

			if part == 1 {
				result, err = solution.PartOne(input)
			} else {
				result, err = solution.PartTwo(input)
			}
			if err != nil {
				return err
			}

			fmt.Printf("Result: %s\n", result)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
