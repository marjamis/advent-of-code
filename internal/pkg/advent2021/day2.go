package advent2021

import (
	"errors"
	"log"
	"strconv"
	"strings"
)

// splitCommand takes a command (such as "forward 3") and splits based on direction and number of units to move
func splitCommand(command string) (string, int, error) {
	splitCommand := strings.Split(command, " ")
	action := splitCommand[0]
	stringNum := splitCommand[1]

	num, err := strconv.Atoi(stringNum)
	if err != nil {
		return "", 0, errors.New("Failed to convert string to int")
	}

	return action, num, nil
}

// Day2Part1 returns the position of the submarine based off of the provided command explanations
func Day2Part1(commands ...string) int {
	var horizontalPosition int
	var depth int

	for _, command := range commands {
		action, num, err := splitCommand(command)
		if err != nil {
			log.Fatal(err)
		}

		switch action {
		case "forward":
			horizontalPosition += num
		case "down":
			depth += num
		case "up":
			depth -= num
		}
	}

	return horizontalPosition * depth
}

// Day2Part2 returns the position of the submarine based off of the updated command explanations
func Day2Part2(commands ...string) int {
	var horizontalPosition int
	var depth int
	var aim int

	for _, command := range commands {
		action, num, err := splitCommand(command)
		if err != nil {
			log.Fatal(err)
		}

		switch action {
		case "forward":
			horizontalPosition += num
			depth += aim * num
		case "down":
			aim += num
		case "up":
			aim -= num
		}
	}

	return horizontalPosition * depth
}
