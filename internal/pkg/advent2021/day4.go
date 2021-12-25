package advent2021

import (
	"log"
	"strconv"
	"strings"
)

type position struct {
	number int
	marked bool
}

// Board layout
type Board [][]position

// Boards is a list of Board's
type Boards []Board

// parseInput takes the input and extracts the bingo numbers and the boards that will be used
func parseInput(input []string) (bingoNumbers []int, boards Boards) {
	numbers := strings.Split(input[0], ",")
	bingoNumbers = make([]int, len(numbers))
	for index, number := range numbers {
		tmp, err := strconv.Atoi(number)
		if err != nil {
			log.Fatal(err)
		}
		bingoNumbers[index] = tmp
	}

	currentBoard := Board{}
	for i := 2; i < len(input); i++ {
		if input[i] != "" {
			boardInputLine := strings.Split(input[i], " ")

			boardLine := []position{}
			for j := 0; j < len(boardInputLine); j++ {
				if boardInputLine[j] != "" {
					num, err := strconv.Atoi(boardInputLine[j])
					if err != nil {
						log.Fatal(err)
					}

					boardLine = append(boardLine, position{
						number: num,
						marked: false,
					})
				}
			}
			currentBoard = append(currentBoard, boardLine)
		} else {
			boards = append(boards, currentBoard)
			currentBoard = Board{}
		}
	}
	boards = append(boards, currentBoard)

	return bingoNumbers, boards
}

// callNumber takes the bingo number and marks this number appropriately for each board in use
func callNumber(number int, boards Boards) {
	for _, board := range boards {
		for row := 0; row < len(board); row++ {
			for col := 0; col < len(board[row]); col++ {
				if board[row][col].number == number {
					board[row][col].marked = true
				}
			}
		}
	}
}

// checkBoardsForWinner iterates through each board to see if there are any boards that have won the game
func checkBoardsForWinner(boards Boards) (winningBoards int, winnerFound bool) {
outerLoop:
	for boardNumber, board := range boards {
		// Rows check
		for row := 0; row < len(board); row++ {
			winner := true
			for col := 0; col < len(board[row]); col++ {
				if !board[row][col].marked {
					winner = false
					break
				}
			}
			if winner {
				winningBoards = boardNumber
				winnerFound = true
				break outerLoop
			}
		}

		// Columns check
		for col := 0; col < len(board[0]); col++ {
			winner := true
			for row := 0; row < len(board); row++ {
				if !board[row][col].marked {
					winner = false
					break
				}
			}
			if winner {
				winningBoards = boardNumber
				winnerFound = true
				break outerLoop
			}
		}
	}

	return
}

// calculateScore takes the given board and returns the number of all non-marked numbers on that board
func calculateScore(board Board) (score int) {
	for row := 0; row < len(board); row++ {
		for col := 0; col < len(board[row]); col++ {
			if !board[row][col].marked {
				score += board[row][col].number
			}
		}
	}

	return
}

// Day4Part1 takes the provided input, the bingo numbers and list of boards, and returns the score of the winning board
func Day4Part1(input []string) (finalScore int) {
	bingoNumbers, boards := parseInput(input)

	for _, number := range bingoNumbers {
		callNumber(number, boards)
		winner, winnerFound := checkBoardsForWinner(boards)
		if winnerFound {
			finalScore = calculateScore(boards[winner]) * number
			break
		}
	}

	return
}

// Day4Part2 takes the provided input, the bingo numbers and list of boards, and returns the board that will win
// last (assuming the game kept playing) and returns this boards score
func Day4Part2(input []string) (finalScore int) {
	bingoNumbers, boards := parseInput(input)

	latestWinningBoard := Board{}
	var latestNumber int

	for _, number := range bingoNumbers {
		callNumber(number, boards)

		// loop this for winners
		test := true
		for test {
			winner, winnerFound := checkBoardsForWinner(boards)
			if winnerFound {
				latestWinningBoard = boards[winner]
				latestNumber = number
				boards = append(boards[:winner], boards[winner+1:]...)
				if len(boards) == 0 {
					break
				}
			} else {
				test = false
			}

		}
	}

	return calculateScore(latestWinningBoard) * latestNumber
}
