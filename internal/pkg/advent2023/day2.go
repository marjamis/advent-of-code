package advent2023

import (
	"strconv"
	"strings"
)

type CubeGame struct {
	gameNumber int
	rounds     []map[string]int
}

type CubeGames []CubeGame

func generateGameDetails(gameData []string) (games CubeGames) {
	for gameIndex, gameDetails := range gameData {
		gameRounds := strings.Split(gameDetails, ":")
		game := CubeGame{gameNumber: gameIndex + 1}

		rounds := strings.Split(gameRounds[1], ";")
		for _, round := range rounds {
			cubes := strings.Split(round, ", ")
			for _, cube := range cubes {
				cubeDetails := strings.Split(strings.Trim(cube, " "), " ")

				colour := cubeDetails[1]
				count, err := strconv.Atoi(cubeDetails[0])
				if err != nil {
					return CubeGames{}
				}

				c := make(map[string]int)
				c[colour] = count
				game.rounds = append(game.rounds, c)
			}
		}

		games = append(games, game)
	}

	return
}

// Day2Part1 returns the sum of all valid games
func Day2Part1(gameData []string) (sumOfValidGameIds int) {
	for _, game := range generateGameDetails(gameData) {
		validGame := true
		for _, round := range game.rounds {
			if round["red"] > 12 || round["green"] > 13 || round["blue"] > 14 {
				validGame = false
				break
			}
		}

		if validGame {
			sumOfValidGameIds += game.gameNumber

		}
	}

	return
}

// Day2Part2 returns the sum of the power of the minimum number of cubes per colour per game
func Day2Part2(gameData []string) (sumOfPowerOfCubes int) {
	for _, game := range generateGameDetails(gameData) {
		minRedCubes := 0
		minGreenCubes := 0
		minBlueCubes := 0

		for _, round := range game.rounds {
			if round["red"] > minRedCubes {
				minRedCubes = round["red"]
			}
			if round["green"] > minGreenCubes {
				minGreenCubes = round["green"]
			}

			if round["blue"] > minBlueCubes {
				minBlueCubes = round["blue"]
			}
		}

		sumOfPowerOfCubes += (minRedCubes * minGreenCubes * minBlueCubes)
	}

	return
}
