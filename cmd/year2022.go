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
