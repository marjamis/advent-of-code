package advent2021

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsLineCorruptedFalse(t *testing.T) {
	tests := []string{
		"[]",
		"<>",
		"{}",
		"()",
		"[<>]",
		"{()()()}",
		"([])",
		"<([{}])>",
		"[<>({}){}[([])<>]]",
		"(((((((((())))))))))",
	}

	for _, test := range tests {
		stack := Stack{}
		_, corrupted := stack.isLineCorrupted(test)
		assert.False(t, corrupted)
	}
}

func TestIsLineCorruptedTrue(t *testing.T) {
	tests := []string{
		"[>",
		"[<>}",
		"{()()()>",
		"(((()))}",
		"<([]){()}[{}])",
	}

	for _, test := range tests {
		stack := Stack{}
		_, corrupted := stack.isLineCorrupted(test)
		assert.True(t, corrupted)
	}
}

func TestLineCorruptedPoppedValue(t *testing.T) {
	tests := []struct {
		input    string
		expected rune
	}{
		{
			"{([(<{}[<>[]}>{[]{[(<()>",
			'}',
		},
		{
			"[[<[([]))<([[{}[[()]]] ",
			')',
		},
		{
			"[{[{({}]{}}([{[{{{}}([]",
			']',
		},
		{
			"[<(<(<(<{}))><([]([]()",
			')',
		},
		{
			"<{([([[(<>()){}]>(<<{{",
			'>',
		},
	}

	for _, test := range tests {
		stack := Stack{}
		firstCorruptedRune, corrupted := stack.isLineCorrupted(test.input)
		assert.True(t, corrupted)
		assert.Equal(t, test.expected, firstCorruptedRune)
	}
}

func TestCompleteLine(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			"[({(<(())[]>[[{[]{<()<>>",
			"}}]])})]",
		},
		{
			"[(()[<>])]({[<{<<[]>>(",
			")}>]})",
		},
		{
			"(((({<>}<{<{<>}{[]{[]{}",
			"}}>}>))))",
		},
		{
			"{<[[]]>}<{[{[{[]{()[[[]",
			"]]}}]}]}>",
		},
		{
			"<{([{{}}[<[[[<>{}]]]>[]]",
			"])}>",
		},
	}

	for _, test := range tests {
		stack := Stack{}
		assert.Equal(t, test.expected, stack.completeLine(test.input))
	}
}

func TestCalculateLineScore(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"}}]])})]",
			288957,
		},
		{
			")}>]})",
			5566,
		},
		{
			"}}>}>))))",
			1480781,
		},
		{
			"]]}}]}]}>",
			995444,
		},
		{
			"])}>",
			294,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, calculateLineScore(test.input))
	}
}

func TestDay10Part1(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"[({(<(())[]>[[{[]{<()<>>",
				"[(()[<>])]({[<{<<[]>>(",
				"{([(<{}[<>[]}>{[]{[(<()>",
				"(((({<>}<{<{<>}{[]{[]{}",
				"[[<[([]))<([[{}[[()]]]",
				"[{[{({}]{}}([{[{{{}}([]",
				"{<[[]]>}<{[{[{[]{()[[[]",
				"[<(<(<(<{}))><([]([]()",
				"<{([([[(<>()){}]>(<<{{",
				"<{([{{}}[<[[[<>{}]]]>[]]",
			},
			26397,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day10Part1(test.input))
	}
}

func TestDay10Part2(t *testing.T) {
	tests := []struct {
		input    []string
		expected int
	}{
		{
			[]string{
				"[({(<(())[]>[[{[]{<()<>>",
				"[(()[<>])]({[<{<<[]>>(",
				"{([(<{}[<>[]}>{[]{[(<()>",
				"(((({<>}<{<{<>}{[]{[]{}",
				"[[<[([]))<([[{}[[()]]]",
				"[{[{({}]{}}([{[{{{}}([]",
				"{<[[]]>}<{[{[{[]{()[[[]",
				"[<(<(<(<{}))><([]([]()",
				"<{([([[(<>()){}]>(<<{{",
				"<{([{{}}[<[[[<>{}]]]>[]]",
			},
			288957,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.expected, Day10Part2(test.input))
	}
}
