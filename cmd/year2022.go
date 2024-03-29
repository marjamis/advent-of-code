package cmd

import (
	"fmt"

	"github.com/marjamis/advent-of-code/internal/pkg/advent2022"
	"github.com/marjamis/advent-of-code/pkg/helpers"
	"github.com/spf13/cobra"
)

const dataDirectory2022 = "./test/advent2022/"

var days2022 = map[string]day{
	"01": {
		Function: func() {
			data := helpers.ReadStringArray(dataDirectory2022 + "day1.txt")
			fmt.Println(advent2022.Day1Part1(data))
			fmt.Println(advent2022.Day1Part2(data))
		}},
	"02": {
		Function: func() {
			data := helpers.ReadStringArray(dataDirectory2022 + "day2.txt")
			fmt.Println(advent2022.Day2Part1(data))
			fmt.Println(advent2022.Day2Part2(data))
		}},
	"03": {
		Function: func() {
			data := helpers.ReadStringArray(dataDirectory2022 + "day3.txt")
			fmt.Println(advent2022.Day3Part1(data))
			fmt.Println(advent2022.Day3Part2(data))
		}},
	"04": {
		Function: func() {
			data := helpers.ReadStringArray(dataDirectory2022 + "day4.txt")
			fmt.Println(advent2022.Day4Part1(data))
			fmt.Println(advent2022.Day4Part2(data))
		}},
	"05": {
		Function: func() {
			data := helpers.ReadStringArray(dataDirectory2022 + "day5.txt")
			fmt.Println(advent2022.Day5Part1(data))
			fmt.Println(advent2022.Day5Part2(data))
		}},
	"06": {
		Function: func() {
			data := helpers.ReadString(dataDirectory2022 + "day6.txt")
			fmt.Println(advent2022.Day6Part1(data))
			fmt.Println(advent2022.Day6Part2(data))
		}},
	"07": {
		Function: func() {
			data := helpers.ReadStringArray(dataDirectory2022 + "day7.txt")
			fmt.Println(advent2022.Day7Part1(data))
			fmt.Println(advent2022.Day7Part2(data))
		}},
	"08": {
		Function: func() {
			data := helpers.ReadIntArray2d(dataDirectory2022 + "day8.txt")
			fmt.Println(advent2022.Day8Part1(data))
			fmt.Println(advent2022.Day8Part2(data))
		}},
	"09": {
		Function: func() {
			data := helpers.ReadStringArray(dataDirectory2022 + "day9.txt")
			fmt.Println(advent2022.Day9Part1(data))
			fmt.Println(advent2022.Day9Part2(data))
		}},
	"10": {
		Function: func() {
			data := helpers.ReadStringArray(dataDirectory2022 + "day10.txt")
			fmt.Println(advent2022.Day10Part1(data))
			advent2022.DisplayFrame(advent2022.Day10Part2(data))
		}},
	"11": {
		Function: func() {
			data := helpers.ReadStringArray(dataDirectory2022 + "day11.txt")
			fmt.Println(advent2022.Day11Part1(data))
			fmt.Println(advent2022.Day11Part2(data))
		}},
	"12": {
		Function: func() {
			data := helpers.ReadRuneArray2d(dataDirectory2022 + "day12.txt")
			fmt.Println(advent2022.Day12Part1(data, 'E'))
			fmt.Println(advent2022.Day12Part2(data, 'E'))
		}},
}

// year2022Cmd represents the year2022 command
var year2022Cmd = &cobra.Command{
	Use:   "year2022",
	Short: "Runs through the each advent day for the year",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("2022 called")
		printAllDaysOutput(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(year2022Cmd)
	addDaySubCommandToYearCommand(year2022Cmd, days2022)
}
