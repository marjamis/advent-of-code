package cmd

import (
	"fmt"

	"github.com/marjamis/advent-of-code/internal/pkg/advent2021"
	"github.com/marjamis/advent-of-code/pkg/helpers"

	"github.com/spf13/cobra"
)

const dataDirectory2021 = "./test/advent2021/"

var days2021 = map[string]day{
	"01": {
		Function: func() {
			input := helpers.ReadIntArray(dataDirectory2021 + "day1.txt")
			fmt.Println(advent2021.Day1Part1(input...))
			fmt.Println(advent2021.Day1Part2(input...))
		}},
	"02": {
		Function: func() {
			input := helpers.ReadStringArray(dataDirectory2021 + "day2.txt")
			fmt.Println(advent2021.Day2Part1(input...))
			fmt.Println(advent2021.Day2Part2(input...))
		}},
	"03": {
		Function: func() {
			input := helpers.ReadStringArray(dataDirectory2021 + "day3.txt")
			fmt.Println(advent2021.Day3Part1(input))
			fmt.Println(advent2021.Day3Part2(input))
		}},
	"04": {
		Function: func() {
			input := helpers.ReadStringArray(dataDirectory2021 + "day4.txt")
			fmt.Println(advent2021.Day4Part1(input))
			fmt.Println(advent2021.Day4Part2(input))
		}},
	"05": {
		Function: func() {
			input := helpers.ReadStringArray(dataDirectory2021 + "day5.txt")
			fmt.Println(advent2021.Day5Part1(input))
			fmt.Println(advent2021.Day5Part2(input))
		}},
	"06": {
		Function: func() {
			// Only the first line from the input is required
			input := helpers.ReadIntCSV(dataDirectory2021 + "day6.csv")[0]
			fmt.Println(advent2021.Day6Part1(input, 80))
			// fmt.Println(advent2021.Day6Part2(input[0], 256))
		}},
	"07": {
		Function: func() {
			// Only the first line from the input is required
			input := helpers.ReadIntCSV(dataDirectory2021 + "day7.csv")[0]
			fmt.Println(advent2021.Day7Part1(input...))
			fmt.Println(advent2021.Day7Part2(input...))
		}},
	"08": {
		Function: func() {
			input := helpers.ReadStringArray(dataDirectory2021 + "day8.txt")
			fmt.Println(advent2021.Day8Part1(input))
			fmt.Println(advent2021.Day8Part2(input))
		}},
	"09": {
		Function: func() {
			input := helpers.ReadStringArray(dataDirectory2021 + "day9.txt")
			fmt.Println(advent2021.Day9Part1(input))
			// fmt.Println(advent2021.Day8Part2(input))
		}},
	"10": {
		Function: func() {
			input := helpers.ReadStringArray(dataDirectory2021 + "day10.txt")
			fmt.Println(advent2021.Day10Part1(input))
			fmt.Println(advent2021.Day10Part2(input))
		}},
	"11": {
		Function: func() {
			input := helpers.ReadIntArray2d(dataDirectory2021 + "day11.txt")
			fmt.Println(advent2021.Day11Part1(input))
			fmt.Println(advent2021.Day11Part2(input))
		}},
}

// year2021Cmd represents the year2021 command
var year2021Cmd = &cobra.Command{
	Use:   "year2021",
	Short: "Runs through the each advent day for the year",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("year2021 called")
		printAllDaysOutput(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(year2021Cmd)
	addDaySubCommandToYearCommand(year2021Cmd, days2021)
}
