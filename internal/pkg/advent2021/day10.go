package advent2021

import (
	"sort"
)

// Stack used as our Stack construct
type Stack struct {
	stack []rune
}

// Push adds a new value to the Stack
func (s *Stack) Push(value rune) {
	s.stack = append(s.stack, value)
}

// Pop returns the last value on the Stack and updates to reomve the value from the Stack
func (s *Stack) Pop() (value rune) {
	// Returns a blank rune if there are no rune's on the Stack
	if len(s.stack) == 0 {
		return ' '
	}

	// Gets the popped value
	value = s.stack[len(s.stack)-1]
	// Removes the value from the Stack
	s.stack = s.stack[:len(s.stack)-1]

	return
}

// isOpenning returns if the bracket is an openning
func isOpenning(r rune) bool {
	return r == '{' || r == '(' || r == '<' || r == '['
}

// isClosing returns if the bracket is a closing
func isClosing(ending, popped rune) (closing bool) {
	switch ending {
	case ')':
		if popped == '(' {
			closing = true
		}
	case '}':
		if popped == '{' {
			closing = true
		}
	case ']':
		if popped == '[' {
			closing = true
		}
	case '>':
		if popped == '<' {
			closing = true
		}
	}

	return
}

// isLineCorrupted returns if the line is corrupted or not and the braket that starts the corruption
func (s Stack) isLineCorrupted(input string) (rune, bool) {
	for _, char := range input {
		charRune := rune(char)

		if isOpenning(charRune) {
			s.Push(char)
		} else {
			if popped := s.Pop(); !isClosing(charRune, popped) {
				return charRune, true
			}
		}
	}

	return ' ', false
}

// calculateCharScore returns the score based on the provided rune
func calculateCharScore(r rune) (s int) {
	switch r {
	case ')':
		s = 3
	case ']':
		s = 57
	case '}':
		s = 1197
	case '>':
		s = 25137
	}

	return
}

// completeLine returns the remainder of a partial line based on the remaining unmatched portions in the Stack
func (s Stack) completeLine(input string) string {
	for _, char := range input {
		charRune := rune(char)
		if isOpenning(charRune) {
			s.Push(char)
		} else {
			s.Pop()
		}
	}

	remainingChars := []rune{}
	for _, charRune := range s.stack {
		switch charRune {
		case '(':
			remainingChars = append(remainingChars, ')')
		case '[':
			remainingChars = append(remainingChars, ']')
		case '{':
			remainingChars = append(remainingChars, '}')
		case '<':
			remainingChars = append(remainingChars, '>')
		}
	}

	orderedChars := []rune{}
	for i := len(remainingChars) - 1; i >= 0; i-- {
		orderedChars = append(orderedChars, remainingChars[i])
	}

	return string(orderedChars)
}

// calculateLineScore returns the score based on the provided remainder of the line
func calculateLineScore(line string) (score int) {
	for _, char := range line {
		score *= 5

		switch char {
		case ')':
			score++
		case ']':
			score += 2
		case '}':
			score += 3
		case '>':
			score += 4
		}
	}

	return
}

// Day10Part1 returns the cumalative score of all the corrupted lines based on the individual char scores
func Day10Part1(inputs []string) (score int) {
	stack := Stack{}

	for _, input := range inputs {
		if char, corrupted := stack.isLineCorrupted(input); corrupted {
			score += calculateCharScore(char)
		}
	}

	return
}

// Day10Part2 returns the middle score from all the incomplete scores after sorting
func Day10Part2(inputs []string) (middleScore int) {
	stack := Stack{}

	scores := []int{}
	for _, input := range inputs {
		if _, corrupted := stack.isLineCorrupted(input); !corrupted {
			// Note: This is an imperfect solution as it runs through some of the input line twice but will do for now
			stack := Stack{}
			lineScore := calculateLineScore(stack.completeLine(input))
			scores = append(scores, lineScore)
		}
	}

	sort.Sort(sort.IntSlice(scores))

	return scores[(len(scores))/2]
}
