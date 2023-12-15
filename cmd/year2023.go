package cmd

import (
	"fmt"

	"github.com/marjamis/advent-of-code/internal/pkg/advent2023"
	"github.com/marjamis/advent-of-code/pkg/helpers"
	"github.com/spf13/cobra"
)

const dataDirectory2023 = "./test/advent2023/"

var days2023 = map[string]day{
	"01": {
		Function: func() {
			data := helpers.ReadRuneArray2d(dataDirectory2023 + "day1.txt")
			fmt.Println(advent2023.Day1Part1(data))
			fmt.Println(advent2023.Day1Part2(data))
		}},
	"02": {
		Function: func() {
			data := helpers.ReadStringArray(dataDirectory2023 + "day2.txt")
			fmt.Println(advent2023.Day2Part1(data))
			fmt.Println(advent2023.Day2Part2(data))
		}},
	"03": {
		Function: func() {
			data := helpers.ReadRuneArray2d(dataDirectory2023 + "day3.txt")
			fmt.Println(advent2023.Day3Part1(data))
			fmt.Println(advent2023.Day3Part2(data))
		}},
	"04": {
		Function: func() {
			data := helpers.ReadStringArray(dataDirectory2023 + "day4.txt")
			fmt.Println(advent2023.Day4Part1(data))
			fmt.Println(advent2023.Day4Part2(data))
		}},
	"05": {
		Function: func() {
			data := helpers.ReadString(dataDirectory2023 + "day5.txt")
			fmt.Println(advent2023.Day5Part1(data))
			fmt.Println(advent2023.Day5Part2(data))
		}},
	"06": {
		Function: func() {
			data := helpers.ReadStringArray(dataDirectory2023 + "day6.txt")
			fmt.Println(advent2023.Day6Part1(data))
			fmt.Println(advent2023.Day6Part2(data))
		}},
	"07": {
		Function: func() {
			data := helpers.ReadStringArray(dataDirectory2023 + "day7.txt")
			fmt.Println(advent2023.Day7Part1(data))
			fmt.Println(advent2023.Day7Part2(data))
		}},
	"08": {
		Function: func() {
			data := helpers.ReadStringArray(dataDirectory2023 + "day8.txt")
			fmt.Println(advent2023.Day8Part1(data))
			fmt.Println(advent2023.Day8Part2(data))
		}},
}

// year2022Cmd represents the year2022 command
var year2023Cmd = &cobra.Command{
	Use:   "year2023",
	Short: "Runs through the each advent day for the year",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("2023 called")
		printAllDaysOutput(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(year2023Cmd)
	addDaySubCommandToYearCommand(year2023Cmd, days2023)
}
