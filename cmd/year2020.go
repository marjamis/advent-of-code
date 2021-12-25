package cmd

import (
	"fmt"

	"github.com/marjamis/advent-of-code/internal/pkg/advent2020"
	"github.com/marjamis/advent-of-code/pkg/helpers"

	"github.com/spf13/cobra"
)

const dataDirectory2020 = "./test/advent2020/"

var days2020 = map[string]day{
	"01": {
		Function: func() {
			data := helpers.ReadIntArray(dataDirectory2020 + "day1.txt")
			fmt.Println(advent2020.Day1(data))
			fmt.Println(advent2020.Day1Part2(data))
		}},
	"02": {
		Function: func() {
			data := helpers.ReadStringArray(dataDirectory2020 + "day2.txt")
			fmt.Println(advent2020.Day2(data, advent2020.Day2CheckOptionGeneral))
			fmt.Println(advent2020.Day2(data, advent2020.Day2CheckOptionSpecial))
		}},
	"03": {
		Function: func() {
			data := helpers.ReadStringArray2d(dataDirectory2020 + "day3.txt")
			fmt.Println(advent2020.Day3(data, []advent2020.ToboganMovement{{X: 3, Y: 1}}))
			fmt.Println(advent2020.Day3(data, []advent2020.ToboganMovement{
				{X: 1, Y: 1},
				{X: 3, Y: 1},
				{X: 5, Y: 1},
				{X: 7, Y: 1},
				{X: 1, Y: 2},
			}))
		},
	},
	"04": {
		Function: func() {
			data := helpers.ReadString(dataDirectory2020 + "day4.txt")
			fmt.Println(advent2020.Day4(data, false))
			fmt.Println(advent2020.Day4(data, true))
		},
	},
	"05": {
		Function: func() {
			data := helpers.ReadStringArray(dataDirectory2020 + "day5.txt")
			fmt.Println(advent2020.Day5(data))
			fmt.Println(advent2020.Day5Part2(data))
		},
	},
	"06": {
		Function: func() {
			data := helpers.ReadString(dataDirectory2020 + "day6.txt")
			fmt.Println(advent2020.Day6(data, false))
			fmt.Println(advent2020.Day6(data, true))
		},
	},
	"07": {
		Function: func() {
			data := helpers.ReadString(dataDirectory2020 + "day7.txt")
			fmt.Println(advent2020.Day7(data, "shiny gold", advent2020.Day7SearchOptionIsIn))
			fmt.Println(advent2020.Day7(data, "shiny gold", advent2020.Day7SearchOptionContains))
		},
	},
	"08": {
		Function: func() {
			data := helpers.ReadStringArray(dataDirectory2020 + "day8.txt")
			fmt.Println(advent2020.Day8(data, false))
			fmt.Println(advent2020.Day8(data, true))
		},
	},
	"09": {
		Function: func() {
			data := helpers.ReadIntArray(dataDirectory2020 + "day9.txt")
			fmt.Println(advent2020.Day9(data, 25))
			fmt.Println(advent2020.Day9Part2(data, 25))
		},
	},
	"10": {
		Function: func() {
			data := helpers.ReadIntArray(dataDirectory2020 + "day10.txt")
			fmt.Println(advent2020.Day10(data))
			fmt.Println(advent2020.Day10Part2(data))
		},
	},
	"11": {
		Function: func() {
			data := helpers.ReadRuneArray2d(dataDirectory2020 + "day11.txt")
			fmt.Println(advent2020.Day11(data, advent2020.Day11RuleOptionBasic))
			fmt.Println(advent2020.Day11(data, advent2020.Day11RuleOptionAdvanced))
		},
	},
	"12": {
		Function: func() {
			data := helpers.ReadStringArray(dataDirectory2020 + "day12.txt")
			fmt.Println(advent2020.Day12(data, advent2020.Day12MovementTypeShip))
			fmt.Println(advent2020.Day12(data, advent2020.Day12MovementTypeWaypoint))
		},
	},
	"13": {
		Function: func() {
			data := helpers.ReadStringArray(dataDirectory2020 + "day13.txt")
			fmt.Println(advent2020.Day13(data))
			fmt.Println(advent2020.Day13Part2(data, 100000000000000))
		},
	},
	"14": {
		Function: func() {
			data := helpers.ReadString(dataDirectory2020 + "day14.txt")
			fmt.Println(advent2020.Day14(data))
			fmt.Println(advent2020.Day14Part2(data))
		},
	},
	"15": {
		Function: func() {
			fmt.Println(advent2020.Day15([]int{2, 0, 1, 7, 4, 14, 18}, advent2020.Day15PositionOption2020))
			fmt.Println(advent2020.Day15([]int{2, 0, 1, 7, 4, 14, 18}, advent2020.Day15PositionOption30mil))
		},
	},
	"16": {
		Function: func() {
			data := helpers.ReadString(dataDirectory2020 + "day16.txt")
			fmt.Println(advent2020.Day16(data))
			fmt.Println(advent2020.Day16Part2(data))
		},
	},
	"18": {
		Function: func() {
			data := helpers.ReadStringArray(dataDirectory2020 + "day18.txt")
			fmt.Println(advent2020.Day18Wrapper(data))
			fmt.Println(advent2020.Day18WrapperPart2(data))
		},
	},
	"19": {
		Function: func() {
			data := helpers.ReadString(dataDirectory2020 + "day19.txt")
			fmt.Println(advent2020.Day19(data, false))
			fmt.Println(advent2020.Day19(data, true))
		},
	},
	"20": {
		Function: func() {
			data := helpers.ReadString(dataDirectory2020 + "day20.txt")
			fmt.Println(advent2020.Day20(data))
		},
	},
	"22": {
		Function: func() {
			data := helpers.ReadString(dataDirectory2020 + "day22.txt")
			fmt.Println(advent2020.Day22(data, false))
			fmt.Println(advent2020.Day22(data, true))
		},
	},
	"23": {
		Function: func() {
			data := `538914762`
			fmt.Println(advent2020.Day23(data, 100))
		},
	},
}

// year2020Cmd represents the year2020 command
var year2020Cmd = &cobra.Command{
	Use:   "year2020",
	Short: "Runs through the each advent day for the year",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("year2020 called")
		printAllDaysOutput(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(year2020Cmd)
	addDaySubCommandToYearCommand(year2020Cmd, days2020)
}
