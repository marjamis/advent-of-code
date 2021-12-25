package cmd

import (
	"fmt"

	"github.com/marjamis/advent-of-code/internal/pkg/advent2019"
	"github.com/marjamis/advent-of-code/pkg/helpers"

	"github.com/spf13/cobra"
)

const dataDirectory2019 = "./test/advent2019/"

var days2019 = map[string]day{
	"01": {
		Function: func() {
			data := helpers.ReadIntArray(dataDirectory2019 + "day1.txt")
			fmt.Println(advent2019.Day1(false, data...))
			fmt.Println(advent2019.Day1(true, data...))
		}},
	"02": {
		Function: func() {
			// Only obtains the first line of ints as this is all the data that's required
			data := helpers.ReadIntCSV(dataDirectory2019 + "day2.csv")[0]
			fmt.Println(advent2019.Day2Part1(data...))
			fmt.Println(advent2019.Day2Part2(data...))
		}},
	"03": {
		Function: func() {
			data := helpers.ReadStringCSV(dataDirectory2019 + "day3.csv")
			fmt.Println(advent2019.Day3(data[0], data[1], advent2019.Day3Manhattan))
			// fmt.Println(advent2019.Day3(data[0], data[1], advent2019.Day3Steps))
		}},
	"04": {
		Function: func() {
			fmt.Println(advent2019.Day4("353096-843212", advent2019.Day4Rules1))
			fmt.Println(advent2019.Day4("353096-843212", advent2019.Day4Rules2))
		}},
	"05": {
		Function: func() {
			data := helpers.ReadIntCSV(dataDirectory2019 + "day5.csv")[0]
			fmt.Println(advent2019.Day5("1", data...))
			fmt.Println(advent2019.Day5("5", data...))
		}},
	"06": {
		Function: func() {
			data := helpers.ReadStringArray(dataDirectory2019 + "day6.txt")
			fmt.Println(advent2019.Day6Part1(data))
			fmt.Println(advent2019.Day6Part2(data))
		}},
}

// year2019Cmd represents the year2019 command
var year2019Cmd = &cobra.Command{
	Use:   "year2019",
	Short: "Runs through the each advent day for the year",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("2019 called")
		printAllDaysOutput(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(year2019Cmd)
	addDaySubCommandToYearCommand(year2019Cmd, days2019)
}
