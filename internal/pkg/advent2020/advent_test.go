package advent2020

import (
	"testing"

	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func init() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{
		DisableColors:    true,
		DisableTimestamp: true,
	})

	log.SetReportCaller(false)
}

func TestDay1(t *testing.T) {
	var expenseReport = []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}

	t.Run("Day 1 Test", func(t *testing.T) {
		assert.Equal(t, 514579, Day1(expenseReport))
	})
	t.Run("Day 1 Part 2 Test", func(t *testing.T) {
		assert.Equal(t, 241861950, Day1Part2(expenseReport))
	})
}

func TestDay2(t *testing.T) {
	var passwordPolicies = []string{
		"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc",
	}

	tests := []struct {
		testName string
		check    Day2CheckOption
		expected int
	}{
		{
			"Day 2 Test",
			Day2CheckOptionGeneral,
			2,
		},
		{
			"Day 2 Part 2 Test",
			Day2CheckOptionSpecial,
			1,
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			assert.Equal(t, test.expected, Day2(passwordPolicies, test.check))
		})
	}
}

func TestDay3(t *testing.T) {
	var tobMap = [][]string{
		{".", ".", "#", "#", ".", ".", ".", ".", ".", ".", "."},
		{"#", ".", ".", ".", "#", ".", ".", ".", "#", ".", "."},
		{".", "#", ".", ".", ".", ".", "#", ".", ".", "#", "."},
		{".", ".", "#", ".", "#", ".", ".", ".", "#", ".", "#"},
		{".", "#", ".", ".", ".", "#", "#", ".", ".", "#", "."},
		{".", ".", "#", ".", "#", "#", ".", ".", ".", ".", "."},
		{".", "#", ".", "#", ".", "#", ".", ".", ".", ".", "#"},
		{".", "#", ".", ".", ".", ".", ".", ".", ".", ".", "#"},
		{"#", ".", "#", "#", ".", ".", ".", "#", ".", ".", "."},
		{"#", ".", ".", ".", "#", "#", ".", ".", ".", ".", "#"},
		{".", "#", ".", ".", "#", ".", ".", ".", "#", ".", "#"},
	}

	var tests = []struct {
		testName string
		expected int
		tobMove  ToboganMovement
	}{
		{
			"Day 3 Test",
			7,
			ToboganMovement{3, 1},
		},
		{
			"Day 3 Test Part 2 - A",
			2,
			ToboganMovement{1, 1},
		},
		{
			"Day 3 Test Part 2 - B",
			7,
			ToboganMovement{3, 1},
		},
		{
			"Day 3 Test Part 2 - C",
			3,
			ToboganMovement{5, 1},
		},
		{
			"Day 3 Test Part 2 - D",
			4,
			ToboganMovement{7, 1},
		},
		{
			"Day 3 Test Part 2 - E",
			2,
			ToboganMovement{1, 2},
		},
	}

	for _, test := range tests {
		t.Run(test.testName, func(t *testing.T) {
			assert.Equal(t, test.expected, day3Counter(tobMap, test.tobMove))
		})
	}

	t.Run("Wrapper test", func(t *testing.T) {
		assert.Equal(t, 336, Day3(tobMap, []ToboganMovement{
			{1, 1},
			{3, 1},
			{5, 1},
			{7, 1},
			{1, 2},
		}))
	})
}

func TestDay4(t *testing.T) {
	var passportData = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

	var validatePassportData = `eyr:1972 cid:100
hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926

iyr:2019
hcl:#602927 eyr:1967 hgt:170cm
ecl:grn pid:012533040 byr:1946

hcl:dab227 iyr:2012
ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277

hgt:59cm ecl:zzz
eyr:2038 hcl:74454a iyr:2023
pid:3556412378 byr:2007

pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980
hcl:#623a2f

eyr:2029 ecl:blu cid:129 byr:1989
iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm

hcl:#888785
hgt:164cm byr:2001 iyr:2015 cid:88
pid:545766238 ecl:hzl
eyr:2022

iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719`

	t.Run("No advanced validation", func(t *testing.T) {
		assert.Equal(t, 2, Day4(passportData, false))
	})

	t.Run("Advanced validation", func(t *testing.T) {
		assert.Equal(t, 4, Day4(validatePassportData, true))
	})
}

func TestDay5(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"FBFBBFFRLR",
			357,
		},
		{
			"BFFFBBFRRR",
			567,
		},
		{
			"FFFBBBFRRR",
			119,
		},
		{
			"BBFFBBFRLL",
			820,
		},
	}

	t.Run("Testing return the seat ID", func(t *testing.T) {
		for _, test := range tests {
			assert.Equal(t, test.expected, day5SeatID(test.input))
		}
	})

	seatLocations := []string{}
	for _, test := range tests {
		seatLocations = append(seatLocations, test.input)
	}

	t.Run("Testing which is the highest seat ID", func(t *testing.T) {
		assert.Equal(t, 820, Day5(seatLocations))
	})

	tests = []struct {
		input    string
		expected int
	}{
		{
			"BBFFBBFRLL",
			820,
		},
		{
			"BBFFBBFRRL",
			822,
		},
		{
			"BBFFBBFRRR",
			823,
		},
	}

	seatLocations = []string{}
	for _, test := range tests {
		seatLocations = append(seatLocations, test.input)
	}

	t.Run("Testing which is missing seat ID", func(t *testing.T) {
		assert.Equal(t, 821, Day5Part2(seatLocations))
	})

}

func TestDay6(t *testing.T) {
	var declartionForms = `abc

a
b
c

ab
ac

a
a
a
a

b`
	t.Run("Unique count of anyone answering yes", func(t *testing.T) {
		assert.Equal(t, 11, Day6(declartionForms, false))
	})

	t.Run("Unique count of everyone answering yes", func(t *testing.T) {
		assert.Equal(t, 6, Day6(declartionForms, true))
	})

	declartionForms += `

ijmp
dmjp
pjm
pmidj
lpjafmzv
`

	t.Run("Unique count of everyone answering yes with ending newline character which causes issue", func(t *testing.T) {
		assert.Equal(t, 9, Day6(declartionForms, true))
	})
}

func TestDay7(t *testing.T) {
	bagName := "shiny gold"
	rules := `light red bags contain 1 bright white bag, 2 muted yellow bags.
dark orange bags contain 3 bright white bags, 4 muted yellow bags.
bright white bags contain 1 shiny gold bag.
muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.
shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.
dark olive bags contain 3 faded blue bags, 4 dotted black bags.
vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.
faded blue bags contain no other bags.
dotted black bags contain no other bags.`

	rules2 := `shiny gold bags contain 2 dark red bags.
dark red bags contain 2 dark orange bags.
dark orange bags contain 2 dark yellow bags.
dark yellow bags contain 2 dark green bags.
dark green bags contain 2 dark blue bags.
dark blue bags contain 2 dark violet bags.
dark violet bags contain no other bags.`

	t.Run("Testing Part 1 rules", func(t *testing.T) {
		assert.Equal(t, 4, Day7(rules, bagName, Day7SearchOptionIsIn))
	})

	t.Run("Testing how many bags would be required", func(t *testing.T) {
		assert.Equal(t, 32, Day7(rules, bagName, Day7SearchOptionContains))
		assert.Equal(t, 126, Day7(rules2, bagName, Day7SearchOptionContains))
	})
}

func TestDay8(t *testing.T) {
	programData := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}

	t.Run("Getting the accumlator value before fixing the program execution", func(t *testing.T) {
		assert.Equal(t, 5, Day8(programData, false))
	})

	t.Run("Getting the accumlator value after fixing the program execution", func(t *testing.T) {
		assert.Equal(t, 8, Day8(programData, true))
	})
}

func TestDay9(t *testing.T) {
	data := []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}

	t.Run("Find the unsummable number", func(t *testing.T) {
		assert.Equal(t, 127, Day9(data, 5))
	})

	t.Run("Find the encryption weakness", func(t *testing.T) {
		assert.Equal(t, 62, Day9Part2(data, 5))
	})
}

func TestDay10(t *testing.T) {
	testData1 := []int{
		16,
		10,
		15,
		5,
		1,
		11,
		7,
		19,
		6,
		12,
		4,
	}

	testData2 := []int{
		28,
		33,
		18,
		42,
		31,
		14,
		46,
		20,
		48,
		47,
		24,
		23,
		49,
		45,
		19,
		38,
		39,
		11,
		1,
		32,
		25,
		35,
		8,
		17,
		7,
		9,
		4,
		2,
		34,
		10,
		3,
	}

	personalTest := []int{
		1,
		2,
		3,
		4,
		6,
		7,
		9,
		12,
	}

	t.Run("Joltage returned for most adaptors", func(t *testing.T) {
		assert.Equal(t, 35, Day10(testData1))
		assert.Equal(t, 220, Day10(testData2))
	})

	t.Run("How many possible combinations?", func(t *testing.T) {
		assert.Equal(t, 29, Day10Part2(personalTest))
		assert.Equal(t, 8, Day10Part2(testData1))
		assert.Equal(t, 19208, Day10Part2(testData2))
	})
}

func TestDay11(t *testing.T) {
	testSeating := [][]rune{
		{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
		{'L', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
		{'L', '.', 'L', '.', 'L', '.', '.', 'L', '.', '.'},
		{'L', 'L', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
		{'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L'},
		{'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
		{'.', '.', 'L', '.', 'L', '.', '.', '.', '.', '.'},
		{'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L'},
		{'L', '.', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L'},
		{'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L'},
	}

	t.Run("Number of occuppied seats with basic rules", func(t *testing.T) {
		assert.Equal(t, 37, Day11(testSeating, Day11RuleOptionBasic))
	})

	t.Run("Number of occuppied seats with advanced rules", func(t *testing.T) {
		assert.Equal(t, 26, Day11(testSeating, Day11RuleOptionAdvanced))
	})
}

func TestDay12(t *testing.T) {
	navigationInstructions := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}

	navigationInstructions2 := []string{
		"F10",
		"N3",
		"F2",
		"L90",
		"F7",
		"R90",
		"F11",
		"L180",
		"S4",
		"E19",
		"L270",
		"R180",
		"W8",
		"F6",
		"R270",
	}

	type rotationTestData struct {
		input     string
		expectedX int64
		expectedY int64
	}

	rotateClockwiseTests := []rotationTestData{
		{"R90", 1, 10},
		{"R90", -10, 1},
		{"R90", -1, -10},
	}

	rotateCounterClockwiseTests := []rotationTestData{
		{"L90", -1, -10},
		{"L90", -10, 1},
		{"L90", 1, 10},
	}

	t.Run("Rotation test", func(t *testing.T) {
		// Setup for default start position
		t.Run("Right", func(t *testing.T) {
			new := day12Position{X: 10, Y: -1}
			for _, test := range rotateClockwiseTests {
				new.X, new.Y = rotateClockwise(new)
				assert.Equal(t, test.expectedX, new.X)
				assert.Equal(t, test.expectedY, new.Y)
			}
		})

		t.Run("Left", func(t *testing.T) {
			new := day12Position{X: 10, Y: -1}
			for _, test := range rotateCounterClockwiseTests {
				new.X, new.Y = rotateCounterClockwise(new)
				assert.Equal(t, test.expectedX, new.X)
				assert.Equal(t, test.expectedY, new.Y)
			}
		})
	})

	t.Run("Manhattan distance without waypoints", func(t *testing.T) {
		assert.Equal(t, 25, Day12(navigationInstructions, Day12MovementTypeShip))
		assert.Equal(t, 34, Day12(navigationInstructions2, Day12MovementTypeShip))
	})

	t.Run("Manhattan distance with waypoints", func(t *testing.T) {
		assert.Equal(t, 286, Day12(navigationInstructions, Day12MovementTypeWaypoint))
	})
}

func TestDay13(t *testing.T) {
	testdata := []string{
		"939",
		"7,13,x,x,59,x,31,19",
	}
	assert.Equal(t, 295, Day13(testdata))

	testdata2 := []struct {
		input            []string
		startingPosition int64 // This is for speed of algorithm
		expected         int
	}{
		{[]string{"Empty", "7,13,x,x,59,x,31,19"}, 0, 1068781},
		{[]string{"Empty", "17,x,13,19"}, 0, 3417},
		{[]string{"Empty", "67,7,59,61"}, 0, 754018},
		{[]string{"Empty", "67,x,7,59,61"}, 0, 779210},
		{[]string{"Empty", "67,7,x,59,61"}, 0, 1261476},
		{[]string{"Empty", "1789,37,47,1889"}, 600000, 1202161486}, // Works just takes an incredible long time from 0. Start at about 600,000 to be acceptable for testing.
	}

	t.Run("Finding the initial timestamps for when these leave at specific intervals", func(t *testing.T) {
		for _, test := range testdata2 {
			assert.Equal(t, test.expected, Day13Part2(test.input, test.startingPosition))
		}
	})
}

func TestDay14(t *testing.T) {
	testdata := `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
	mem[8] = 11
	mem[7] = 101
	mem[8] = 0`

	testdata2 := `mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[7] = 101
mask = 111XXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 0`

	t.Run("Testing V1 of the mask", func(t *testing.T) {
		assert.Equal(t, 165, Day14(testdata))
		assert.Equal(t, 60129542309, Day14(testdata2))
	})

	testdataV2 := `mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1`
	t.Run("Testing V2 of the mask", func(t *testing.T) {
		assert.Equal(t, 208, Day14Part2(testdataV2))
	})

}

func TestDay15(t *testing.T) {
	tests := []struct {
		input         []int
		expected2020  int
		expected30mil int
	}{
		{
			[]int{0, 3, 6},
			436,
			175594,
		},
		{
			[]int{1, 3, 2},
			1,
			2578,
		},
		{
			[]int{2, 1, 3},
			10,
			3544142,
		},
		{
			[]int{1, 2, 3},
			27,
			261214,
		},
		{
			[]int{2, 3, 1},
			78,
			6895259,
		},
		{
			[]int{3, 2, 1},
			438,
			18,
		},
		{
			[]int{3, 1, 2},
			1836,
			362,
		},
	}

	t.Run("Position @ 2020", func(t *testing.T) {
		for _, test := range tests {
			assert.Equal(t, test.expected2020, Day15(test.input, Day15PositionOption2020))
		}
	})

	t.Run("Position @ 30000000", func(t *testing.T) {
		if testing.Short() {
			t.Skip("Skipping during short tests as these take quite a while.")
		}
		for _, test := range tests {
			assert.Equal(t, test.expected30mil, Day15(test.input, Day15PositionOption30mil))
		}
	})
}

func TestDay16(t *testing.T) {
	testdata := `class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12`

	assert.Equal(t, 71, Day16(testdata))

	testdata2 := `departure class: 0-1 or 4-19
departure running: 9-11 or 17-20
row: 0-5 or 8-19
departure seat: 0-13 or 16-19

your ticket:
11,12,13,10

nearby tickets:
77,77,77,77
3,9,18,18
15,1,5,19
5,14,9,10
99,99,99,99`

	assert.Equal(t, int64(1560), Day16Part2(testdata2))
}

func TestDay17(t *testing.T) {
	// 	testdata := `.#.
	// ..#
	// ###`

	// 	assert.Equal(t, 112, Day17(testdata))
}

func TestDay18(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			"1 + 2 * 3 + 4 * 5 + 6",
			71,
		},
		{
			"1 + (2 * 3) + (4 * (5 + 6))",
			51,
		},
		{
			"2 * 3 + (4 * 5)",
			26,
		},
		{
			"5 + (8 * 3 + 9 + 3 * 4 * 3)",
			437,
		},
		{
			"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			12240,
		},
		{
			"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			13632,
		},
		{
			"2 + (8 + (2 + 9) * 7 * 4 + (8 * 6 + 8 * 6 * 4 + 5) + 6) + 4 * (9 + (3 + 4) * (3 * 9 * 7 * 6 + 9 + 7) + 9 * 6 * 3)",
			627268266,
		},
	}

	t.Run("Calculate without precedence order", func(t *testing.T) {
		for _, test := range tests {
			assert.Equal(t, test.expected, Day18(test.input, false))
		}
	})

	inputs := make([]string, len(tests))
	for i := range tests {
		inputs[i] = tests[i].input
	}

	t.Run("Find the sum of all above data", func(t *testing.T) {
		assert.Equal(t, 627294723, Day18Wrapper(inputs))
	})

	testsP2 := []struct {
		input    string
		expected int
	}{
		{
			"1 + 2 * 3 + 4 * 5 + 6",
			231,
		},
		{
			"1 + (2 * 3) + (4 * (5 + 6))",
			51,
		},
		{
			"2 * 3 + (4 * 5)",
			46,
		},
		{
			"5 + (8 * 3 + 9 + 3 * 4 * 3)",
			1445,
		},
		{
			"5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))",
			669060,
		},
		{
			"((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2",
			23340,
		},
		{
			"2 + (8 + (2 + 9) * 7 * 4 + (8 * 6 + 8 * 6 * 4 + 5) + 6) + 4 * (9 + (3 + 4) * (3 * 9 * 7 * 6 + 9 + 7) + 9 * 6 * 3)",
			966941349120,
		},
		{
			"3 * 8 + 9 * (3 * 4)",
			612,
		},
		{
			"(4 * 2 + 4 * 8 + 4) + (3 * 9 + (2 * 4 + 2 * 7) * 9 + 7) * (5 + 9) * (7 + (3 * 4 * 4 * 6) + 8 * 7) + 8 + 8",
			142170336,
		},
		{
			"(6 * (5 + 8 * 7 * 8 + 4) * (7 + 7 * 3 * 5)) + 5 * (8 + (8 + 3 + 5 + 5) + (3 + 2 + 7 * 2 * 9) + 6 * 5 + (2 * 6)) * ((4 * 3) + 3) + 9 * (3 + 6 * 2 + 3 * 8)",
			50726061864000,
		},
		{
			"((9 + 7 + 9 + 7 * 7 * 4) * 6 * 5) + 5 + (8 + (2 + 6 * 9 * 7 * 9 + 9) * (9 * 3 * 8 + 9 * 6 * 2) + 6 * 7) * (2 + 6 + (3 * 4 + 4) * 6 * 5) * 9 * 5",
			15141458520000,
		},
	}

	t.Run("Calculate with +'s before *'s", func(t *testing.T) {
		for _, test := range testsP2 {
			assert.Equal(t, test.expected, Day18(test.input, true))
		}
	})

	// As I know the length change this approach rather than expensive appends
	inputs = make([]string, len(testsP2))
	result := 0
	for i := range testsP2 {
		inputs[i] = testsP2[i].input
		result += testsP2[i].expected
	}

	assert.Equal(t, result, Day18WrapperPart2(inputs))
}

func TestDay19(t *testing.T) {
	testdata := `0: 4 1 5
1: 2 3 | 3 2
2: 4 4 | 5 5
3: 4 5 | 5 4
4: "a"
5: "b"

ababbb
bababa
abbbab
aaabbb
aaaabbb`

	testdata2 := `0: 2 5
1: 3 3 | 4 4
2: 1 1 | 3 5
3: "a"
4: "b"
5: 3 4 | 4 3

ababbb
bababa
abbbab
aabbba
aaabbb
aabbab
aaaabbb
abaab`

	t.Run("Find the valid messages", func(t *testing.T) {
		assert.Equal(t, 2, Day19(testdata, false))
		assert.Equal(t, 3, Day19(testdata2, false))
	})

	testdataP2 := `42: 9 14 | 10 1
9: 14 27 | 1 26
10: 23 14 | 28 1
1: "a"
11: 42 31
5: 1 14 | 15 1
19: 14 1 | 14 14
12: 24 14 | 19 1
16: 15 1 | 14 14
31: 14 17 | 1 13
6: 14 14 | 1 14
2: 1 24 | 14 4
0: 8 11
13: 14 3 | 1 12
15: 1 | 14
17: 14 2 | 1 7
23: 25 1 | 22 14
28: 16 1
4: 1 1
20: 14 14 | 1 15
3: 5 14 | 16 1
27: 1 6 | 14 18
14: "b"
21: 14 1 | 1 14
25: 1 1 | 1 14
22: 14 14
8: 42
26: 14 22 | 1 20
18: 15 15
7: 14 5 | 1 21
24: 14 1

abbbbbabbbaaaababbaabbbbabababbbabbbbbbabaaaa
bbabbbbaabaabba
babbbbaabbbbbabbbbbbaabaaabaaa
aaabbbbbbaaaabaababaabababbabaaabbababababaaa
bbbbbbbaaaabbbbaaabbabaaa
bbbababbbbaaaaaaaabbababaaababaabab
ababaaaaaabaaab
ababaaaaabbbaba
baabbaaaabbaaaababbaababb
abbbbabbbbaaaababbbbbbaaaababb
aaaaabbaabaaaaababaa
aaaabbaaaabbaaa
aaaabbaabbaaaaaaabbbabbbaaabbaabaaa
babaaabbbaaabaababbaabababaaab
aabbbbbaabbbaaaaaabbbbbababaaaaabbaaabba`

	t.Run("Find the valid messages", func(t *testing.T) {
		assert.Equal(t, 3, Day19(testdataP2, false))
		assert.Equal(t, 12, Day19(testdataP2, true))
	})
}

func TestDay20(t *testing.T) {
	testdata := `Tile 2311:
..##.#..#.
##..#.....
#...##..#.
####.#...#
##.##.###.
##...#.###
.#.#.#..##
..#....#..
###...#.#.
..###..###

Tile 1951:
#.##...##.
#.####...#
.....#..##
#...######
.##.#....#
.###.#####
###.##.##.
.###....#.
..#.#..#.#
#...##.#..

Tile 1171:
####...##.
#..##.#..#
##.#..#.#.
.###.####.
..###.####
.##....##.
.#...####.
#.##.####.
####..#...
.....##...

Tile 1427:
###.##.#..
.#..#.##..
.#.##.#..#
#.#.#.##.#
....#...##
...##..##.
...#.#####
.#.####.#.
..#..###.#
..##.#..#.

Tile 1489:
##.#.#....
..##...#..
.##..##...
..#...#...
#####...#.
#..#.#.#.#
...#.#.#..
##.#...##.
..##.##.##
###.##.#..

Tile 2473:
#....####.
#..#.##...
#.##..#...
######.#.#
.#...#.#.#
.#########
.###.#..#.
########.#
##...##.#.
..###.#.#.

Tile 2971:
..#.#....#
#...###...
#.#.###...
##.##..#..
.#####..##
.#..####.#
#..#.#..#.
..####.###
..#.#.###.
...#.#.#.#

Tile 2729:
...#.#.#.#
####.#....
..#.#.....
....#..#.#
.##..##.#.
.#.####...
####.#.#..
##.####...
##..#.##..
#.##...##.

Tile 3079:
#.#.#####.
.#..######
..#.......
######....
####.#..#.
.#...#.##.
#.#####.##
..#.###...
..#.......
..#.###...`

	t.Run("Rotations", func(t *testing.T) {
		t.Run("Testing that the right (and left) rotations works in a full loop", func(t *testing.T) {
			testRotateExpected1Right := day20Tile{"1", [][]rune{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			}, map[string]day20RotationFingerprint{}}
			testRotateExpected2Right := day20Tile{"1", [][]rune{
				{9, 8, 7},
				{6, 5, 4},
				{3, 2, 1},
			}, map[string]day20RotationFingerprint{}}

			testRotateExpected3Right := day20Tile{"1", [][]rune{
				{3, 6, 9},
				{2, 5, 8},
				{1, 4, 7},
			}, map[string]day20RotationFingerprint{}}

			testRotateExpected4Right := day20Tile{"1", [][]rune{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, map[string]day20RotationFingerprint{}}

			testRotateInput := day20Tile{"1", [][]rune{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, map[string]day20RotationFingerprint{}}

			assert.Equal(t, testRotateExpected1Right, testRotateInput.day20RotateRight())
			assert.Equal(t, testRotateExpected2Right, testRotateInput.day20RotateRight().day20RotateRight())
			assert.Equal(t, testRotateExpected3Right, testRotateInput.day20RotateRight().day20RotateRight().day20RotateRight())
			assert.Equal(t, testRotateExpected4Right, testRotateInput.day20RotateRight().day20RotateRight().day20RotateRight().day20RotateRight())
		})

		t.Run("Testing horizontal flip", func(t *testing.T) {
			testHFlipInput := day20Tile{"1", [][]rune{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, map[string]day20RotationFingerprint{}}

			testHFlipExpected := day20Tile{"1", [][]rune{
				{3, 2, 1},
				{6, 5, 4},
				{9, 8, 7},
			}, map[string]day20RotationFingerprint{}}

			assert.Equal(t, testHFlipExpected, testHFlipInput.day20FlipHorizontal())
			assert.Equal(t, testHFlipInput, testHFlipInput.day20FlipHorizontal().day20FlipHorizontal())
		})

		t.Run("Testing vertical flip", func(t *testing.T) {
			testVFlipInput := day20Tile{"1", [][]rune{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			}, map[string]day20RotationFingerprint{}}

			testVFlipExpected := day20Tile{"1", [][]rune{
				{7, 8, 9},
				{4, 5, 6},
				{1, 2, 3},
			}, map[string]day20RotationFingerprint{}}

			assert.Equal(t, testVFlipExpected, testVFlipInput.day20FlipVertical())
			assert.Equal(t, testVFlipInput, testVFlipInput.day20FlipVertical().day20FlipVertical())
		})
	})

	t.Run("Fingerprinting", func(t *testing.T) {
		testInput := day20Tile{"1", [][]rune{
			{'#', '.', '.'},
			{'.', '.', '#'},
			{'#', '#', '#'},
		}, map[string]day20RotationFingerprint{}}

		assert.Equal(t, 73, fingerprintInput{'#', '.', '.', '#', '.', '.', '#'}.day20GenerateFingerprint())
		assert.Equal(t, 11, fingerprintInput{'#', '.', '#', '#'}.day20GenerateFingerprint())
		t.Run("Generating all fingerprints from tile", func(t *testing.T) {
			testInput.day20TileFingerprintInputData("R=1")
			assert.Equal(t, day20RotationFingerprint{
				rotation: "R=1",
				rotatedImage: [][]rune{
					{'#', '.', '.'},
					{'.', '.', '#'},
					{'#', '#', '#'},
				},
				top:    4,
				bottom: 7,
				left:   5,
				right:  3,
			}, testInput.fingerprints["R=1"])
		})

		t.Run("Generating all fingerprints from tile", func(t *testing.T) {
			testInput2 := day20Tile{"1", [][]rune{
				{'#', '.', '.'},
				{'.', '.', '#'},
				{'#', '#', '#'},
			}, map[string]day20RotationFingerprint{}}

			testInput2 = testInput2.day20RotateRight()
			testInput2.day20TileFingerprintInputData("R=1")
			assert.Equal(t, day20RotationFingerprint{
				rotation: "R=1",
				rotatedImage: [][]rune{
					{'#', '.', '#'},
					{'#', '.', '.'},
					{'#', '#', '.'},
				},
				top:    5,
				bottom: 6,
				left:   7,
				right:  4,
			}, testInput2.fingerprints["R=1"])

			testInput2 = testInput2.day20RotateRight()
			testInput2.day20TileFingerprintInputData("R=1")
			assert.Equal(t, day20RotationFingerprint{
				rotation: "R=1",
				rotatedImage: [][]rune{
					{'#', '#', '#'},
					{'#', '.', '.'},
					{'.', '.', '#'},
				},
				top:    7,
				bottom: 1,
				left:   6,
				right:  5,
			}, testInput2.fingerprints["R=1"])
		})

	})

	t.Run("Finding the multiplication of corners", func(t *testing.T) {
		assert.Equal(t, 20899048083289, Day20(testdata))
	})
}

func TestDay22(t *testing.T) {
	testdata := `Player 1:
9
2
6
3
1

Player 2:
5
8
4
7
10`

	recursiveLoop := `Player 1:
43
19

Player 2:
2
29
14`

	t.Run("", func(t *testing.T) {
		assert.Equal(t, 306, Day22(testdata, false))
		assert.Equal(t, 0, Day22(recursiveLoop, true))
		assert.Equal(t, 291, Day22(testdata, true))
	})
}

func TestDay23(t *testing.T) {
	t.Skip()
	testdata := `389125467`

	t.Run("", func(t *testing.T) {
		assert.Equal(t, "92658374", Day23(testdata, 10))
		assert.Equal(t, "67384529", Day23(testdata, 100))
		if testing.Short() {
			t.Skip("Skipping during short tests as these take quite a while.")
		}
		assert.Equal(t, "149245887792", Day23(testdata, 10000000))
	})
}
