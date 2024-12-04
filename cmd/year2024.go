package cmd

import (
	"fmt"

	"github.com/marjamis/advent-of-code/internal/pkg/advent2024"
	"github.com/marjamis/advent-of-code/pkg/helpers"
	"github.com/spf13/cobra"
)

const dataDirectory2024 = "./test/advent2024/"

var days2024 = map[string]day{
	"01": {
		Function: func() {
			data := helpers.ReadStringArray(dataDirectory2024 + "day1.txt")
			fmt.Println(advent2024.Day1Part1(data))
			fmt.Println(advent2024.Day1Part2(data))
		}},
	"03": {
		Function: func() {
			data := helpers.ReadString(dataDirectory2024 + "day3.txt")
			fmt.Println(advent2024.Day3Part1(data))
			fmt.Println(advent2024.Day3Part2(data))
		}},
	"04": {
		Function: func() {
			data := helpers.ReadRuneArray2d(dataDirectory2024 + "day4.txt")
			fmt.Println(advent2024.Day4Part1(data))
			fmt.Println(advent2024.Day4Part2(data))
		}},
}

// year2022Cmd represents the year2022 command
var year2024Cmd = &cobra.Command{
	Use:   "year2024",
	Short: "Runs through the each advent day for the year",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("2024 called")
		printAllDaysOutput(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(year2024Cmd)
	addDaySubCommandToYearCommand(year2024Cmd, days2024)
}
