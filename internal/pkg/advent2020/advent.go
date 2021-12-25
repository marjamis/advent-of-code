package advent2020

import (
	"bufio"
	"fmt"
	"math"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"github.com/marjamis/advent-of-code/pkg/helpers"
	log "github.com/sirupsen/logrus"
)

func day1CheckListAddsTo2020(pointers []int, expenseReport []int) bool {
	plus := expenseReport[pointers[0]]
	for _, pointer := range pointers[1:] {
		plus += expenseReport[pointer]
	}

	if plus == 2020 {
		return true
	}

	return false
}

func day1MultiplyListValues(pointers []int, expenseReport []int) int {
	multi := expenseReport[pointers[0]]
	for _, pointer := range pointers[1:] {
		multi *= expenseReport[pointer]
	}

	return multi
}

// Day1 entry function
func Day1(expenseReport []int) (twoMulti int) {
	sort.Sort(sort.Reverse(sort.IntSlice(expenseReport)))
out:
	for index := range expenseReport {
		for index2 := range expenseReport[index:] {
			pointers := []int{index, index + index2}
			if day1CheckListAddsTo2020(pointers, expenseReport) {
				twoMulti = day1MultiplyListValues(pointers, expenseReport)
				break out
			}
		}
	}

	return
}

// Day1Part2 entry function
func Day1Part2(expenseReport []int) (threeMulti int) {
	sort.Sort(sort.Reverse(sort.IntSlice(expenseReport)))
out:
	for index := range expenseReport {
		for index2 := range expenseReport[index:] {
			for index3 := range expenseReport[index2:] {
				pointers := []int{index, index + index2, index2 + index3}
				if day1CheckListAddsTo2020(pointers, expenseReport) {
					threeMulti = day1MultiplyListValues(pointers, expenseReport)
					break out
				}
			}
		}
	}

	return
}

// Day2CheckOption option type
type Day2CheckOption int

const (
	// Day2CheckOptionGeneral option
	Day2CheckOptionGeneral = 0
	// Day2CheckOptionSpecial option
	Day2CheckOptionSpecial = 1
)

func day2GenerateConfiguration(policy string) (rangeLower int64, rangeUpper int64, character string, password string) {
	split := strings.Split(policy, " ")
	rant := strings.Split(split[0], "-")
	rangeLower, _ = strconv.ParseInt(rant[0], 0, 64)
	rangeUpper, _ = strconv.ParseInt(rant[1], 0, 64)
	character = strings.Split(split[1], ":")[0]
	password = split[2]

	return
}

var day2checks = map[Day2CheckOption](func(int64, int64, string, string) bool){
	Day2CheckOptionGeneral: (func(rangeLower int64, rangeUpper int64, character string, password string) bool {
		numberOfInstances := int64(strings.Count(password, character))
		if numberOfInstances >= rangeLower && numberOfInstances <= rangeUpper {
			return true
		}

		return false
	}),
	Day2CheckOptionSpecial: (func(rangeLower int64, rangeUpper int64, character string, password string) bool {
		var X, Y bool

		// The -1's are due to the task as the customer starts from 1 rather than 0.
		if string(password[rangeLower-1]) == character {
			X = true
		}
		if string(password[rangeUpper-1]) == character {
			Y = true
		}

		if (X || Y) && !(X && Y) {
			return true
		}

		return false
	}),
}

// Day2 entry function
func Day2(passwordPolicies []string, check Day2CheckOption) (count int) {
	f := day2checks[check]
	for _, policy := range passwordPolicies {
		if f(day2GenerateConfiguration(policy)) {
			count++
		}
	}

	return
}

// ToboganMovement is the X, Y coordinates of the toboggan
type ToboganMovement struct {
	X int64
	Y int64
}

func day3Counter(tobMap [][]string, tobMovement ToboganMovement) (count int) {
	var x int64
	var y int64
	for y < (int64(len(tobMap)) - 1) {
		log.Debugf("x: %d - tobMove.X: %d - len(tobMap): %d\n", x, tobMovement.X, int64(len(tobMap)))
		if x+tobMovement.X >= int64(len(tobMap[0])) {
			x = (x + tobMovement.X) - int64(len(tobMap[0]))
		} else {
			x += tobMovement.X
		}

		y += tobMovement.Y

		if tobMap[y][x] == "#" {
			count++
		}
	}

	return
}

// Day3 entry function
func Day3(tobMap [][]string, runs []ToboganMovement) (multi int) {
	multi = 1
	for _, run := range runs {
		multi *= day3Counter(tobMap, run)
	}

	return
}

// Day4 entry function
func Day4(passportData string, advancedValidation bool) (count int) {
	type Validation struct {
		Type         string
		MinimumValue int
		MaximumValue int
	}

	re := `([[:ascii:]][^ \n]*)`
	requiredFields := []struct {
		FieldName          string
		FieldFinder        *regexp.Regexp
		ValidationFunction (func(string) bool)
	}{
		{ // (Birth Year)
			"byr",
			regexp.MustCompile(`byr:` + re),
			func(input string) bool {
				num, _ := strconv.Atoi(input)
				return num <= 2002 && num >= 1920
			},
		},
		{ // (Issue Year)
			"iyr",
			regexp.MustCompile(`iyr:` + re),
			func(input string) bool {
				num, _ := strconv.Atoi(input)
				return num <= 2020 && num >= 2010
			},
		},
		{ // (Expiration Year)
			"eyr",
			regexp.MustCompile(`eyr:` + re),
			func(input string) bool {
				num, _ := strconv.Atoi(input)
				return num <= 2030 && num >= 2020
			},
		},
		{ // (Height)
			"hgt",
			regexp.MustCompile(`hgt:` + re),
			func(input string) bool {
				num, _ := strconv.ParseInt(input[:len(input)-2], 0, 64)
				if strings.Contains(input, "cm") {
					if num >= 150 && num <= 193 {
						return true
					}
				}
				if strings.Contains(input, "in") {
					if num >= 59 && num <= 76 {
						return true
					}
				}
				return false
			},
		},
		{ // (Hair Color)
			"hcl",
			regexp.MustCompile(`hcl:` + re),
			func(input string) bool {
				re := regexp.MustCompile(`#[0-9a-f]{6}`)
				return re.MatchString(input)
			},
		},
		{ // (Eye Color)
			"ecl",
			regexp.MustCompile(`ecl:` + re),
			func(input string) bool {
				colours := []string{
					"amb",
					"blu",
					"brn",
					"gry",
					"grn",
					"hzl",
					"oth",
				}
				for _, colour := range colours {
					if input == colour {
						return true
					}
				}
				return false
			},
		},
		{ // (Passport ID)
			"pid",
			regexp.MustCompile(`pid:` + re),
			func(input string) bool {
				return len(input) == 9
			},
		},
		// cid (Country ID)
	}

	for _, passport := range strings.Split(passportData, "\n\n") {
		rfCheck := 0
	out:
		for _, rf := range requiredFields {
			output := rf.FieldFinder.FindStringSubmatch(passport)
			log.Debugf("Output: %+v", output)
			if len(output) != 2 {
				continue out
			}

			if advancedValidation && rf.ValidationFunction != nil && !rf.ValidationFunction(output[1]) {
				continue out
			}
			rfCheck++
		}
		if rfCheck == len(requiredFields) {
			count++
		}
	}

	return
}

func day5Parser(lower int, upper int, direction rune) (int, int) {
	// fmt.Printf("Lower: %d, Upper: %d\n", lower, upper)
	lf := float64(lower)
	uf := float64(upper)
	val := int(lf + (math.RoundToEven(uf-lf) / 2))
	if direction == 'F' || direction == 'L' {
		return lower, val
	}
	return val, upper
}

func day5Wrapper(lower int, upper int, data string) int {
	for _, d := range data {
		lower, upper = day5Parser(lower, upper, d)
	}

	return upper
}

func day5SeatID(seatLocation string) (id int) {
	rowData := seatLocation[:7]
	colData := seatLocation[7:10]

	return (day5Wrapper(0, 127, rowData) * 8) + day5Wrapper(0, 7, colData)
}

// Day5 entry function
func Day5(seatLocations []string) (id int) {
	for _, seatLocation := range seatLocations {
		if num := day5SeatID(seatLocation); num > id {
			id = num
		}
	}
	return
}

// Day5Part2 entry function
func Day5Part2(seatLocations []string) (id int) {
	seatIDs := []int{}
	for _, seatLocation := range seatLocations {
		seatIDs = append(seatIDs, day5SeatID(seatLocation))
	}
	sort.Sort(sort.IntSlice(seatIDs))

	seat := seatIDs[0] - 1
	for _, seatID := range seatIDs {
		// log.Debug("Index: %d, SeatID: %d\n", index, seatID)
		if (seatID - seat) > 1 {
			id = seatID - 1
			break
		}
		seat = seatID
	}

	return
}

func day6FindCount(declartionForm string, isEveryone bool) int {
	questions := map[rune]int{}
	log.Debugf("%s", declartionForm)

	re := regexp.MustCompile(`[[:alpha:]]+`)
	persons := re.FindAllString(declartionForm, -1)
	for _, person := range persons {
		for _, char := range person {
			questions[char]++
		}
	}

	count := len(questions)
	if isEveryone {
		count = 0
		for _, question := range questions {
			if question == len(persons) {
				count++
			}
		}
		log.Debugf("Number of people: %d - Count: %d - Mapping: %+v", len(persons), count, questions)
	}

	return count
}

// Day6 entry function
func Day6(declartionForms string, isEveryone bool) (count int) {
	forms := strings.Split(declartionForms, "\n\n")
	for _, form := range forms {
		count += day6FindCount(form, isEveryone)
	}

	return
}

type day7Contains struct {
	key      string
	numberOf int
}

type day7Bag struct {
	Contains []day7Contains
	IsIn     []string
}

// Day7SearchOption option type
type Day7SearchOption int

const (
	// Day7SearchOptionIsIn option
	Day7SearchOptionIsIn = 0
	// Day7SearchOptionContains option
	Day7SearchOptionContains = 1
)

func day7SearchIsIn(bags map[string]day7Bag, findBag string, fbs *[]string) {
	for _, bag := range bags[findBag].IsIn {
		*fbs = append(*fbs, bag)
		day7SearchIsIn(bags, bag, fbs)
	}

	return
}

func day7SearchContains(bags map[string]day7Bag, findBag string) (count int) {
	for _, bag := range bags[findBag].Contains {
		if bag.numberOf != 0 {
			count += bag.numberOf + (bag.numberOf * day7SearchContains(bags, bag.key))
		}
	}

	return
}

// Day7 entry function
func Day7(rulesList string, findBag string, typeOfCheck Day7SearchOption) (count int) {
	rules := strings.Split(rulesList, "\n")
	bags := map[string]day7Bag{}
	for _, rule := range rules {
		bag := strings.Split(rule, " bags contain ")[0]
		// Not required in tests but required in actual data. Should be fixed.
		if len(bag) < 1 {
			continue
		}
		contains := strings.Split(rule, " bags contain ")[1]
		re := regexp.MustCompile(`(.*? bag)s?[,.]`)
		nc := re.FindAllStringSubmatch(contains, -1)
		currentBags := []day7Contains{}

		for _, n := range nc {
			trmmed := strings.Trim(n[1], " ")
			count, _ := strconv.Atoi(strings.Split(trmmed, " ")[0])
			key := strings.Split(trmmed, " ")[1] + " " + strings.Split(trmmed, " ")[2]

			existingBag, ok := bags[key]
			if !ok {
				bags[key] = day7Bag{
					nil,
					[]string{
						bag,
					},
				}
			} else {
				var isin = existingBag.IsIn
				isin = append(isin, bag)
				bags[key] = day7Bag{
					existingBag.Contains,
					isin,
				}
			}

			currentBags = append(currentBags, day7Contains{key, count})
		}
		existingBag, ok := bags[bag]
		if !ok {
			bags[bag] = day7Bag{
				currentBags,
				nil,
			}
		} else {
			contains := existingBag.Contains
			contains = append(contains, currentBags...)
			bags[bag] = day7Bag{
				contains,
				existingBag.IsIn,
			}
		}
	}

	switch typeOfCheck {
	case Day7SearchOptionIsIn:
		fbs := []string{}
		day7SearchIsIn(bags, findBag, &fbs)
		count = len(helpers.RemoveDuplicates(fbs))
	case Day7SearchOptionContains:
		count = day7SearchContains(bags, findBag)
	}

	return
}

type day8OpData struct {
	opcode    string
	direction int
	visited   bool
}

const (
	day8OpCodenop = "nop"
	day8OpCodejmp = "jmp"
	day8Opcodeacc = "acc"
)

func day8ExecuteProgram(operations []day8OpData) (acc int, ok bool) {
	pos := 0
	for pos != len(operations) {
		// log.Debug("Opcode: %s - Direction: %d\n", changedOperations[pos].opcode, changedOperations[pos].direction)
		if operations[pos].visited {
			break
		}

		operations[pos].visited = true

		switch operations[pos].opcode {
		case day8OpCodenop:
			pos++
		case day8Opcodeacc:
			acc += operations[pos].direction
			pos++
		case day8OpCodejmp:
			pos += operations[pos].direction
		}
	}

	if pos == len(operations) {
		ok = true
		return
	}

	return
}

// Day8 entry function
func Day8(programData []string, rotateClockwise bool) (acc int) {
	operations := []day8OpData{}
	for _, d := range programData {
		da, _ := strconv.Atoi(strings.Split(d, " ")[1])
		operations = append(operations, day8OpData{
			strings.Split(d, " ")[0],
			da,
			false,
		})
	}

	if !rotateClockwise {
		acc, _ = day8ExecuteProgram(operations)
	} else {
		for index := range operations {
			changedOperations := make([]day8OpData, len(operations))
			copy(changedOperations, operations)

			if changedOperations[index].opcode == day8OpCodenop || changedOperations[index].opcode == day8OpCodejmp {
				if changedOperations[index].opcode == day8OpCodenop {
					changedOperations[index].opcode = day8OpCodejmp
				} else if changedOperations[index].opcode == day8OpCodejmp {
					changedOperations[index].opcode = day8OpCodenop
				}
			} else {
				continue
			}

			var ok bool
			acc, ok = day8ExecuteProgram(changedOperations)
			if ok {
				return
			}
		}
	}

	return
}

func day9IsSummable(data []int, sumValue int) bool {
	log.Debugf("Sum required to calculate: %d", sumValue)
	for startingIndex := range data {
		for currentIndex := range data {
			// As it can't be checked against itself for now simply skip.
			if startingIndex == currentIndex {
				continue
			}
			if data[startingIndex]+data[currentIndex] == sumValue {
				log.Debugf("P: %d - S: %d - sumValue: %d", data[startingIndex], data[currentIndex], sumValue)
				return true
			}
		}
	}

	return false
}

// Day9 entry function
func Day9(data []int, preceding int) (firstNonSummableNumber int) {
	for index, number := range data[preceding:] {
		if !day9IsSummable(data[index:index+preceding], number) {
			firstNonSummableNumber = number
			break
		}
	}

	return firstNonSummableNumber
}

// Day9Part2 entry function
func Day9Part2(data []int, preceding int) int {
	unsummableNumber := Day9(data, preceding)

	smallest := 0
	largest := 0
out:
	for startingIndex := range data {
		// Resets the values for a new round of trying to find the smallest and largest values in the contigous array of array values that add up to unsummableNumber.
		// Using a count here to simplify the summing of values rather than looping through the numbers array constantly.
		count := 0
		numbers := []int{}
		for currentIndex := range data[startingIndex:] {
			count += data[currentIndex+startingIndex]
			numbers = append(numbers, data[currentIndex+startingIndex])
			if count == unsummableNumber {
				// Sort to make it easier to find the smallest and largest values, as required by the problem.
				sort.Sort(sort.IntSlice(numbers))
				smallest = numbers[0]
				largest = numbers[len(numbers)-1]
				break out
			}

			// If the count is larger than the number we're adding to no point in continuing in this loop and we break out to try another starting point.
			if count >= unsummableNumber {
				break
			}
		}
	}

	return smallest + largest
}

// Day10 entry function
func Day10(adaptors []int) int {
	sort.Sort(sort.IntSlice(adaptors))
	// Adding in the outlet of value 0 to the beginning of the array
	adaptors = append([]int{0}, adaptors...)
	// As the device is always 3 more jolts than the largest adaptors joltage we add that number to the end
	adaptors = append(adaptors, adaptors[len(adaptors)-1]+3)

	ones := 0
	threes := 0
	currentJoltage := 0
	for adaptor := range adaptors {
		log.Debugf("Adaptor number: %d - Adaptor: %d", adaptor, adaptors[adaptor])

		// To ensure we dont pass the boundaries of the array we do set a default, which is adding the three and then if we exceed the array simply use the last element of the array.
		checkEnd := adaptor + 3
		if checkEnd >= len(adaptors) {
			checkEnd = (len(adaptors))
		}

		for index := range adaptors[adaptor:checkEnd] {
			log.Debugf("Current Joltage: %d", currentJoltage)
			if adaptors[adaptor+index] == currentJoltage+1 {
				currentJoltage++
				ones++
				break
			}
			if adaptors[adaptor+index] == currentJoltage+3 {
				currentJoltage += 3
				threes++
				break
			}
		}
	}
	log.Debugf("Ones: %d and Threes: %d", ones, threes)

	return ones * threes
}

// Day10Part2 entry function
func Day10Part2(adaptors []int) int {
	// Adding in the outlet of value 0 to the beginning of the array
	adaptors = append([]int{0}, adaptors...)
	// As the device is always 3 more than the largest adaptors joltage we add that number to the end
	adaptors = append(adaptors, adaptors[len(adaptors)-1]+3)
	sort.Sort(sort.IntSlice(adaptors))

	type node struct {
		node     int
		branches []int
	}
	nodes := []node{}

	// Go through the list of adaptors to see which parent adaptor can be used with the possible following adaptors, that meet the puzzles criteria, within the 2 jolts.
	for i, adaptor := range adaptors {
		node := node{
			node:     adaptor,
			branches: []int{},
		}
		for j := 1; j <= 3; j++ {
			if i+j <= (len(adaptors)-1) && adaptors[i+j]-adaptor <= 3 {
				node.branches = append(node.branches, adaptors[i+j])
			}
		}
		nodes = append(nodes, node)
	}

	// Create a map of adaptors which return the number possible other adaptors it can be connected to.
	adaptorToBranchesMap := map[int]int{}

	// Loops through the adaptors list backwards to calculate the number of branches each adaptor has and slowly progress back to the root node for the overal number of possibilities.
	for i := (len(nodes) - 2); i >= 0; i-- {
		currentNode := nodes[i]
		log.Debugf("Node Number: %d", currentNode.node)

		count := 0
		for _, branch := range currentNode.branches {
			if val, ok := adaptorToBranchesMap[branch]; ok {
				count += val
			} else {
				count++
			}
		}
		adaptorToBranchesMap[currentNode.node] = count
	}

	// Sort the keys here to ensure an ordered map.
	var orderedKeys []int
	for k := range adaptorToBranchesMap {
		orderedKeys = append(orderedKeys, k)
	}
	sort.Sort(sort.IntSlice(orderedKeys))

	// Returns the data from the first key, which is the root node. This contains the number of all branches, hence combinations that exist for the given data.
	return adaptorToBranchesMap[orderedKeys[0]]
}

const (
	// Day11RuleOptionBasic option
	Day11RuleOptionBasic = 0
	// Day11RuleOptionAdvanced option
	Day11RuleOptionAdvanced = 1
)

type day11SeatPosition struct {
	X int
	Y int
}

// Day11RuleOption option type
type Day11RuleOption int

var (
	day11Directions = []struct {
		xdiff int
		ydiff int
	}{
		{0, -1},  // Up
		{0, 1},   // Down
		{-1, 0},  // Left
		{1, 0},   // Right
		{-1, -1}, // Diagonal Up/Left
		{1, -1},  // Diagonal Up/Right
		{-1, 1},  // Diagonal Down/Left
		{1, 1},   // Diagonal Down/Right
	}

	day11RuleFunctions = map[Day11RuleOption](func(seatMap [][]rune, currentSeat day11SeatPosition) (validSeatPositions []day11SeatPosition)){
		Day11RuleOptionBasic: (func(seatMap [][]rune, currentSeat day11SeatPosition) (validSeatPositions []day11SeatPosition) {
			for _, direction := range day11Directions {
				// Assign these for easier usage throughout
				x := currentSeat.X + direction.xdiff
				y := currentSeat.Y + direction.ydiff
				//General check to make sure it's a valid location on the 2D array
				if helpers.IsLocationValid(seatMap, x, y) {
					// If the type is a empty spot then simply skip as not valid for checking.
					if seatMap[y][x] != '.' {
						validSeatPositions = append(validSeatPositions, day11SeatPosition{X: x, Y: y})
					}
				}
			}

			return
		}),
		Day11RuleOptionAdvanced: (func(seatMap [][]rune, currentSeat day11SeatPosition) (validSeatPositions []day11SeatPosition) {
			for _, direction := range day11Directions {
				for x, y := (currentSeat.X + direction.xdiff), (currentSeat.Y + direction.ydiff); helpers.IsLocationValid(seatMap, x, y); x, y = x+direction.xdiff, y+direction.ydiff {
					if seatMap[y][x] != '.' {
						validSeatPositions = append(validSeatPositions, day11SeatPosition{X: x, Y: y})
						break
					}
				}
			}

			return
		}),
	}
)

func day11NumberOfOccupiedSeatsVisible(seatMap [][]rune, seats []day11SeatPosition) (count int) {
	for _, seat := range seats {
		if seatMap[seat.Y][seat.X] == '#' {
			count++
		}
	}

	return
}

func day11NumberOfOccupiedSeats(seatMap [][]rune) (count int) {
	for y, seat := range seatMap {
		for x := range seat {
			if seatMap[y][x] == '#' {
				count++
			}
		}
	}

	return
}

func day11Rules(seatMap [][]rune, wSeatMap [][]rune, currentSeat day11SeatPosition, ruleOption Day11RuleOption) {
	adjacentSeatCount := day11NumberOfOccupiedSeatsVisible(seatMap, day11RuleFunctions[ruleOption](seatMap, currentSeat))

	switch seatMap[currentSeat.Y][currentSeat.X] {
	case '.':
		return
	case 'L':
		if adjacentSeatCount == 0 {
			wSeatMap[currentSeat.Y][currentSeat.X] = '#'
		}
	case '#':
		if (adjacentSeatCount >= 4 && ruleOption == Day11RuleOptionBasic) || (adjacentSeatCount >= 5 && ruleOption == Day11RuleOptionAdvanced) {
			wSeatMap[currentSeat.Y][currentSeat.X] = 'L'
		}
	}
}

// Day11 entry function
func Day11(seatMap [][]rune, ruleOption Day11RuleOption) (occupiedSeats int) {
	// Basically run forever until there is a break which is caused by a stabilisation in the numbers as per the puzzle.
	for true {
		wSeatMap := make([][]rune, len(seatMap))
		for index := range seatMap {
			wSeatMap[index] = make([]rune, len(seatMap[index]))
			copy(wSeatMap[index], seatMap[index])
		}

		for y := range seatMap {
			for x := range seatMap[y] {
				// log.Infof("X: (%d/%d), Y: (%d/%d)", x, len(seatMap[x]), y, len(seatMap))
				day11Rules(seatMap, wSeatMap, day11SeatPosition{X: x, Y: y}, ruleOption)
			}
		}

		seatMap = wSeatMap

		if val := day11NumberOfOccupiedSeats(seatMap); val != occupiedSeats {
			occupiedSeats = val
		} else {
			break
		}
	}

	return
}

const (
	north day12Direction = 0
	east  day12Direction = 1
	south day12Direction = 2
	west  day12Direction = 3
	// Day12MovementTypeShip option
	Day12MovementTypeShip day12MovementType = 0
	// Day12MovementTypeWaypoint option
	Day12MovementTypeWaypoint day12MovementType = 1
)

type day12Direction int64
type day12MovementType int

type day12Position struct {
	direction day12Direction
	X         int64
	Y         int64
}

func day12CompassPoint(d day12Direction) (s string) {
	switch d {
	case north:
		s = "north"
	case south:
		s = "south"
	case east:
		s = "east"
	case west:
		s = "west"
	}

	return
}

func (shipPosition *day12Position) Day12MovementTypeShip(action rune, value int64) {
	switch action {
	case 'N':
		shipPosition.Y -= value
	case 'S':
		shipPosition.Y += value
	case 'E':
		shipPosition.X += value
	case 'W':
		shipPosition.X -= value
	case 'L':
		switch value {
		case 90:
			shipPosition.direction = (shipPosition.direction + 3) % 4
		case 180:
			shipPosition.direction = (shipPosition.direction + 2) % 4
		case 270:
			shipPosition.direction = (shipPosition.direction + 1) % 4
		}
	case 'R':
		switch value {
		case 90:
			shipPosition.direction = (shipPosition.direction + 1) % 4
		case 180:
			shipPosition.direction = (shipPosition.direction + 2) % 4
		case 270:
			shipPosition.direction = (shipPosition.direction + 3) % 4
		}
	case 'F':
		switch shipPosition.direction {
		case north:
			shipPosition.Y -= value
		case east:
			shipPosition.X += value
		case south:
			shipPosition.Y += value
		case west:
			shipPosition.X -= value
		}
	}
}

func rotateClockwise(waypoint day12Position) (int64, int64) {
	return (waypoint.Y * -1), waypoint.X
}

func rotateCounterClockwise(waypoint day12Position) (int64, int64) {
	return waypoint.Y, (waypoint.X * -1)
}

func (shipPosition *day12Position) Day12MovementTypeWaypoint(action rune, value int64, waypoint *day12Position) {
	switch action {
	case 'N':
		waypoint.Y -= value
	case 'S':
		waypoint.Y += value
	case 'E':
		waypoint.X += value
	case 'W':
		waypoint.X -= value
	case 'L':
		switch value {
		case 90:
			x, y := rotateCounterClockwise(*waypoint)
			waypoint.X = x
			waypoint.Y = y
		case 180:
			for i := 0; i < 2; i++ {
				x, y := rotateCounterClockwise(*waypoint)
				waypoint.X = x
				waypoint.Y = y
			}
		case 270:
			for i := 0; i < 3; i++ {
				x, y := rotateCounterClockwise(*waypoint)
				waypoint.X = x
				waypoint.Y = y
			}
		}
	case 'R':
		switch value {
		case 90:
			x, y := rotateClockwise(*waypoint)
			waypoint.X = x
			waypoint.Y = y
		case 180:
			for i := 0; i < 2; i++ {
				x, y := rotateClockwise(*waypoint)
				waypoint.X = x
				waypoint.Y = y
			}
		case 270:
			for i := 0; i < 3; i++ {
				log.Debugf("Index: %d", i)
				x, y := rotateClockwise(*waypoint)
				waypoint.X = x
				waypoint.Y = y
			}
		}
	case 'F':
		shipPosition.X += value * waypoint.X
		shipPosition.Y += value * waypoint.Y
	}
}

// Day12 entry function
func Day12(navigationInstructions []string, movementType day12MovementType) (manhattanDistance int) {
	shipPosition := day12Position{direction: 1}
	waypoint := day12Position{X: 10, Y: -1}
	for _, instruction := range navigationInstructions {
		action := rune(instruction[0])
		value, _ := strconv.ParseInt(instruction[1:], 0, 64)

		switch movementType {
		case Day12MovementTypeShip:
			shipPosition.Day12MovementTypeShip(action, value)
		case Day12MovementTypeWaypoint:
			shipPosition.Day12MovementTypeWaypoint(action, value, &waypoint)
		}

		log.Debugf("Action: %c - Count: %d - Ships direction: %+v (X:%d/Y:%d)", action, value, day12CompassPoint(shipPosition.direction), shipPosition.X, shipPosition.Y)
		log.Debugf("Action: %c - Count: %d - Waypoint direction: %+v (X:%d/Y:%d)", action, value, day12CompassPoint(waypoint.direction), waypoint.X, waypoint.Y)
	}

	return helpers.ManhattansDistance(int(shipPosition.X), int(shipPosition.Y), 0, 0)
}

// Day13 entry function
func Day13(data []string) (n int) {
	initialTimestamp, _ := strconv.ParseInt(data[0], 0, 64)
	buses := strings.Split(data[1], ",")

	smallest := struct {
		timestamp int64
		busID     int64
	}{
		timestamp: 99999999999999999,
		busID:     0,
	}
	for _, bus := range buses {
		if bus == "x" {
			continue
		}
		log.Debugf("Bus id: %s", bus)
		busID, _ := strconv.ParseInt(bus, 0, 64)

		for i := initialTimestamp; i <= busID+initialTimestamp; i++ {
			completeNumber := float64(i) / float64(busID)
			if math.Mod(completeNumber, 1) == 0 {
				log.Debugf("Found the earliest. Time: %d - Number: %f", i, completeNumber)
				// If the current timestamp is smaller that the smallest.timestamp (across all buses computed) then this bus route and time becomes the smallest.
				if i < smallest.timestamp {
					log.Debugf("Adding bus %d to smallest", busID)
					smallest.busID = busID
					smallest.timestamp = i
				}
				break
			}
		}
	}

	log.Debugf("BusID: %d - Timestamp: %d - Initial time: %d", smallest.busID, smallest.timestamp, initialTimestamp)
	return int(smallest.busID * (smallest.timestamp - initialTimestamp))
}

// Day13Part2 entry function
func Day13Part2(data []string, startPosition int64) (earliestTimestamp int) {
	timingData := map[int64]int{}
	buses := strings.Split(data[1], ",")
	for i, bus := range buses {
		if bus == "x" {
			continue
		}
		log.Infof("Bus '%s' needs to depart 't+%d'", bus, i)
		busID, _ := strconv.ParseInt(bus, 0, 64)
		timingData[busID] = i
	}

	var largest int64
	for k := range timingData {
		if int64(k) > largest {
			largest = k
		}
	}
	log.Infof("Length: %d, Values: %+v, Largest: %d", len(timingData), timingData, largest)

	//rearrange tPlys as per the largest
	for k, v := range timingData {
		if k != largest {
			timingData[k] = v - timingData[largest]
		}
		log.Infof("Bus '%d' needs to depart 't+%d'", k, timingData[k])
	}
	timingData[largest] = 0
	log.Infof("Length: %d, Values: %+v, Largest: %d", len(timingData), timingData, largest)

	// Starting iteration - normally 0 but allows for easy skipping.
	batchSize := int64(1000000000)
	timestampToStartBatchAt := startPosition
	for true {

		initialValidNumber := int64((float64(largest) - math.Mod(float64(timestampToStartBatchAt), float64(largest))) + float64(timestampToStartBatchAt))
		log.Infof("Timestamp: %d", timestampToStartBatchAt)
		// TODO fix for a small batch size that this would still work, for example a batch size of 1 which currently wouldn't work I think or would it just be not optimised? need to test
		for timestampOfLargestNumber := initialValidNumber; timestampOfLargestNumber < timestampToStartBatchAt+batchSize; timestampOfLargestNumber += largest {
			countOfAligningTimes := 1
			earliestTimestamp := timestampOfLargestNumber

			for busID, busTime := range timingData {
				// fmt.Printf("BusID: %d, BusTime: %d, BusTimestamp: %d, ModINput: %f", busID, busTime, timestampMinusCurrentBusTimePosition, float64(timestampMinusCurrentBusTimePosition)/float64(busID))
				if int64(busID) == largest {
					// log.Debugf("Skipping largest bus...")
					continue
				}
				timestampMinusCurrentBusTimePosition := timestampOfLargestNumber + int64(busTime)
				modValue := math.Mod(float64(timestampMinusCurrentBusTimePosition)/float64(busID), float64(1))
				if modValue == 0 {
					// log.Info("Time Match...")
					countOfAligningTimes++
					if timestampMinusCurrentBusTimePosition < earliestTimestamp {
						earliestTimestamp = timestampMinusCurrentBusTimePosition
					}
				}
			}

			if countOfAligningTimes == len(timingData) {
				log.Infof("Match found starting timestamp is: %d", earliestTimestamp)
				return int(earliestTimestamp)
			}
		}
		timestampToStartBatchAt += batchSize
	}

	return
}

type day14Bits [36]byte

func day14ConvertToBinary(number int) day14Bits {
	// t := number
	new := day14Bits{}
	max := int(math.Pow(2, 35))
	for p, l := max, 0; (p >= 1) && (l < len(new)); p, l = p/2, l+1 {
		// log.Debugf("Power of: %d with number: %d - Result: %v", p, number, number >= p)
		if number >= p {
			new[l] = 1
			number -= p
			// log.Debugf("Flipped: %d", p)
		} else {
			new[l] = 0
		}

		if number == 0 {
			break
		}
	}

	// log.Debugf("Number: %d into binary: %+v", t, new)

	return new
}

func day14ApplyMask(num day14Bits, mask day14Bits) day14Bits {
	// log.Debugf("Applying mask: %+v", mask)
	for i, v := range mask {
		if v != 'X' {
			num[i] = v
		}
	}

	return num
}

func day14MaskCreation(stringMask string) (mask day14Bits) {
	for i, char := range stringMask {
		if char == '0' {
			mask[i] = 0
		} else if char == '1' {
			mask[i] = 1
		} else {
			mask[i] = byte(char)
		}
	}

	return
}

func (mem day14Bits) day14ConvertToInt() (val int) {
	max := int(math.Pow(2, 35))

	for p, l := max, 0; (p >= 1) && (l < len(mem)); p, l = p/2, l+1 {
		if mem[l] == 1 {
			val += p
		}
	}

	return
}

// Day14 entry function
func Day14(data string) (sum int) {
	memory := map[string]day14Bits{}
	str := strings.Split(data, "mask = ")
	for _, entry := range str {
		mask := day14MaskCreation(strings.Split(entry, "\n")[0])

		log.Debug("Printing memory allocations...")
		memMask := regexp.MustCompile(`mem\[(.*)\] = (.*)`)
		additions := memMask.FindAllStringSubmatch(entry, -1)
		for _, addition := range additions {
			// log.Debugf("Memory location: %+v to: %+v", addition[1], addition[2])
			// Initialise
			_, ok := memory[addition[1]]
			if !ok {
				new := day14Bits{}
				for i := range new {
					new[i] = 0
				}
				memory[addition[1]] = new
			}
			n, _ := strconv.ParseInt(addition[2], 0, 64)

			memory[addition[1]] = day14ApplyMask(day14ConvertToBinary(int(n)), mask)
		}
	}

	log.Debug("Printing memory locations...")
	for _, mem := range memory {
		v := mem.day14ConvertToInt()
		// log.Debugf("Memory: %s == %d int: %d", key, mem, v)
		sum += v
	}

	return
}

func day14ApplyMaskV2(templateAddress day14Bits, mask day14Bits) (addresses []day14Bits) {
	// log.Debugf("Applying mask: %+v", mask)
	// templateAddress := day14Bits{}
	for i, v := range mask {
		switch v {
		case 'X':
			templateAddress[i] = 'X'
		case 1:
			templateAddress[i] = 1
		case 0:
			// Do nothing
		}
	}

	new := templateAddress
	addresses = append(addresses, new)
	for i, add := range templateAddress {
		if add == 'X' {
			for _, existing := range addresses {
				again := existing
				addresses = append(addresses, again)
			}

			middle := len(addresses) / 2
			for j := range addresses[:middle] {
				addresses[j][i] = 0
			}

			for j := range addresses[middle:] {
				addresses[j][i] = 1
			}
		}
	}

	// log.Debugf("Mask: %+v", templateAddress)
	// for _, i := range addresses {
	// 	log.Debug(i)
	// 	log.Debug(i.day14ConvertToInt())
	// }

	return
}

// Day14Part2 entry function
func Day14Part2(data string) (sum int) {
	memory := map[int]day14Bits{}
	str := strings.Split(data, "mask = ")
	for _, entry := range str {
		mask := day14MaskCreation(strings.Split(entry, "\n")[0])

		log.Debug("Printing memory allocations...")
		memMask := regexp.MustCompile(`mem\[(.*)\] = (.*)`)
		additions := memMask.FindAllStringSubmatch(entry, -1)
		for _, addition := range additions {
			log.Debugf("Memory location: %+v to: %+v", addition[1], addition[2])
			// Initialise
			write, _ := strconv.Atoi(addition[2])
			address, _ := strconv.Atoi(addition[1])
			_, ok := memory[address]
			if !ok {
				new := day14Bits{}
				for i := range new {
					new[i] = 0
				}
				memory[address] = new
			}

			for _, addre := range day14ApplyMaskV2(day14ConvertToBinary(address), mask) {
				newAddress := addre.day14ConvertToInt()
				_, ok := memory[newAddress]
				if !ok {
					new := day14Bits{}
					for i := range new {
						new[i] = 0
					}
					memory[newAddress] = new
				}
				log.Debug(newAddress)
				memory[newAddress] = day14ConvertToBinary(write)
			}
		}
	}

	log.Debug("Printing memory locations...")
	for _, mem := range memory {
		v := mem.day14ConvertToInt()
		// log.Debugf("Memory: %d == %d int: %d", key, mem, v)
		sum += v
	}

	return
}

type day15Positions struct {
	mostRecentPosition       int
	secondMostRecentPosition int
}

type day15PositionsMap map[int]day15Positions

func (positions day15PositionsMap) updatePositions(key int, value int) {
	_, ok := positions[key]
	if !ok {
		positions[key] = day15Positions{
			mostRecentPosition:       value,
			secondMostRecentPosition: -1,
		}
	} else {
		positions[key] = day15Positions{
			value,
			positions[key].mostRecentPosition,
		}
	}
}

const (
	// Day15PositionOption2020 option
	Day15PositionOption2020 Day15PositionOption = 2020
	// Day15PositionOption30mil option
	Day15PositionOption30mil Day15PositionOption = 30000000
)

// Day15PositionOption option type
type Day15PositionOption int

// Day15 entry function
func Day15(startingNumbers []int, position Day15PositionOption) (number int) {
	positions := day15PositionsMap{}
	for i, n := range startingNumbers {
		positions.updatePositions(n, i)
	}

out:
	for i := len(startingNumbers); i < int(position); i++ {
		numberToFind := startingNumbers[len(startingNumbers)-1]
		// log.Debugf("Number to find in map with key: %d: %+v", numberToFind, positions[numberToFind])
		if positions[numberToFind].mostRecentPosition != -1 &&
			positions[numberToFind].secondMostRecentPosition != -1 {
			val := positions[numberToFind].mostRecentPosition - positions[numberToFind].secondMostRecentPosition

			// log.Debugf("1. Last was: %d", val)
			startingNumbers = append(startingNumbers, val)
			positions.updatePositions(val, i)
			continue out
		}

		// log.Debugf("2. Last was: %d", 0)
		startingNumbers = append(startingNumbers, 0)
		positions.updatePositions(0, i)
	}

	return startingNumbers[len(startingNumbers)-1]
}

type day16Range struct {
	rangeFrom int64
	rangeTo   int64
}

type day16FieldData struct {
	typeOf  string
	ranges  []day16Range
	columns []int
}

// Day16 entry function
func Day16(ticketData string) (errorRate int) {
	splits := strings.Split(ticketData, "\n\n")
	dataFields := strings.Split(splits[0], "\n")
	nearbyTickets := strings.Split(splits[2], "\n")[1:]

	fieldData := []day16FieldData{}

	for _, df := range dataFields {
		fieldType := strings.Split(df, ":")[0]
		log.Debug(fieldType)
		f := strings.Split(strings.Split(df, ": ")[1], " or ")
		new := day16FieldData{
			typeOf: fieldType,
		}
		for _, b := range f {
			vs := strings.Split(b, "-")
			from, _ := strconv.ParseInt(vs[0], 0, 64)
			to, _ := strconv.ParseInt(vs[1], 0, 64)
			new.ranges = append(new.ranges, day16Range{from, to})
			// log.Debugf("From: %d to: %d", from, to)
		}
		fieldData = append(fieldData, new)
	}

	// for _, f := range fieldData {
	// 	log.Debugf("Field data: %+v", f)
	// }

	invalids := []int64{}
	for _, nbt := range nearbyTickets {
		// log.Debugf("NBT: %s", nbt)
		nbtFields := strings.Split(nbt, ",")
		for _, f := range nbtFields {
			c, _ := strconv.ParseInt(f, 0, 64)
			// log.Debugf("Nbt: %d", c)
			valids := 0
			for _, mfields := range fieldData {
				for _, fields := range mfields.ranges {
					if c >= fields.rangeFrom && c <= fields.rangeTo {
						// log.Debugf("From: %d to: %d - Number: %d", fields.rangeFrom, fields.rangeTo, c)
						valids++
					}
				}
			}
			if valids == 0 {
				invalids = append(invalids, c)
			}
		}
	}

	var sum int64
	for _, i := range invalids {
		sum += i
	}

	return int(sum)
}

func remove(fieldData []day16FieldData, col int, fieldType string) {
	for index, i := range fieldData {
		for s, j := range i.columns {
			if j == col && i.typeOf != fieldType {
				fieldData[index].columns = append(fieldData[index].columns[:s], fieldData[index].columns[s+1:]...)
			}
		}
	}
}

func day16FieldToTicketsValid(field day16FieldData, ticketData []int64) bool {
	count := 0
	for _, i := range ticketData {
		for _, r := range field.ranges {
			if i >= r.rangeFrom && i <= r.rangeTo {
				count++
				break
			}
		}
	}

	if count == len(ticketData) {
		log.Debugf("Count: %d - Field: %s - TicketData: %+v", count, field.typeOf, ticketData)
		return true
	}

	return false
}

// Day16Part2 entry function
func Day16Part2(ticketData string) (errorRate int64) {
	splits := strings.Split(ticketData, "\n\n")
	dataFields := strings.Split(splits[0], "\n")
	myTicket := strings.Split(splits[1], "\n")[1]
	nearbyTickets := strings.Split(splits[2], "\n")[1:]

	fieldData := []day16FieldData{}

	for _, df := range dataFields {
		fieldType := strings.Split(df, ":")[0]
		log.Debug(fieldType)
		f := strings.Split(strings.Split(df, ": ")[1], " or ")
		new := day16FieldData{
			typeOf: fieldType,
		}
		for _, b := range f {
			vs := strings.Split(b, "-")
			from, _ := strconv.ParseInt(vs[0], 0, 64)
			to, _ := strconv.ParseInt(vs[1], 0, 64)
			new.ranges = append(new.ranges, day16Range{from, to})
			// log.Debugf("From: %d to: %d", from, to)
		}
		fieldData = append(fieldData, new)
	}

	for _, f := range fieldData {
		log.Debugf("Field data: %+v", f)
	}

	invalids := []int64{}
	valids := []string{}
	for _, nbt := range nearbyTickets {
		// log.Debugf("NBT: %s", nbt)
		nbtFields := strings.Split(nbt, ",")
		throwaway := false
		for _, f := range nbtFields {
			c, _ := strconv.ParseInt(f, 0, 64)
			// log.Debugf("Nbt: %d", c)
			valids := 0
			for _, mfields := range fieldData {
				for _, fields := range mfields.ranges {
					if c >= fields.rangeFrom && c <= fields.rangeTo {
						// log.Debugf("From: %d to: %d - Number: %d", fields.rangeFrom, fields.rangeTo, c)
						valids++
					}
				}
			}
			if valids == 0 {
				invalids = append(invalids, c)
				// Throw away the ticket
				throwaway = true
			}
		}
		if !throwaway {
			valids = append(valids, nbt)
		}
	}

	nbts := [][]string{}
	for _, valid := range valids {
		nbtFields := strings.Split(valid, ",")
		nbts = append(nbts, nbtFields)
	}

	for col := 0; col < len(nbts[0]); col++ {
		ticketCol := []int64{}
		for _, ticket := range nbts {
			t, _ := strconv.ParseInt(ticket[col], 0, 64)
			ticketCol = append(ticketCol, t)
		}
		for i, field := range fieldData {
			if day16FieldToTicketsValid(field, ticketCol) {
				// log.Infof("Found type '%s' for column data", field.typeOf)
				fieldData[i].columns = append(fieldData[i].columns, col)
			}
		}
	}

	count := 0
	for count < len(fieldData) {
		count = 0
		for _, field := range fieldData {
			if len(field.columns) == 1 {
				remove(fieldData, field.columns[0], field.typeOf)
				count++
			}
		}
	}

	for _, field := range fieldData {
		log.Infof("%+v", field)
	}

	var multi int64 = 1

	blah := strings.Split(myTicket, ",")
	for _, field := range fieldData {
		if strings.HasPrefix(field.typeOf, "departure") {
			log.Info(field)
			mtv, _ := strconv.ParseInt(blah[field.columns[0]], 0, 64)
			multi *= mtv
		}
	}

	log.Infof("Multi is: %d", multi)

	return multi
}

// type energySource [][][]bool

// func (array energySource) day17Print() {
// 	zs := (len(array) / 2) * -1
// 	for _, z := range array {
// 		log.Debugf("Z: %d", zs)
// 		zs++
// 		for _, y := range z {
// 			for _, x := range y {
// 				var v rune
// 				if x {
// 					v = '#'
// 				} else {
// 					v = '.'
// 				}
// 			}
// 		}
// 	}
// }

// func (array energySource) day17CountActive() (activeCount int) {
// 	for _, z := range array {
// 		for _, y := range z {
// 			for _, x := range y {
// 				if x {
// 					activeCount++
// 				}
// 			}
// 		}
// 	}

// 	return
// }

// func day17CreateBlank2D(z, y, x int) [][]bool {
// 	arr := [][]bool{}
// 	for i := 0; i < y; i++ {
// 		temp := make([]bool, x)
// 		arr = append(arr, temp)
// 	}

// 	return arr
// }

// // This will expand the energySource in every direction (x, y and z) by 1.
// func (array energySource) day17Expand() energySource {
// 	zLen := len(array)
// 	yLen := len(array[0])
// 	xLen := len(array[0][0])
// 	log.Printf("Z: %d, Y: %d, X: %d", zLen, yLen, xLen)
// 	plus := 0
// 	if zLen != 1 {
// 		plus = 2
// 	}

// 	newZ := energySource{}
// 	newZ = append(newZ, day17CreateBlank2D(zLen, yLen+plus, xLen+plus))
// 	for z := 0; z < zLen; z++ {
// 		newY := [][]bool{}
// 		if plus != 0 {
// 			newY = append(newY, make([]bool, yLen+plus))
// 		}
// 		for y := 0; y < yLen; y++ {
// 			// Takes initial row data and adds a false on either side for new inactive states
// 			if zLen != 1 {
// 				// Append false to the end of the pulled array data
// 				fromData := append(array[z][y], false)
// 				// Prepend false
// 				fromData = append([]bool{false}[0:], fromData...)
// 				newY = append(newY, fromData)
// 			} else {
// 				newY = append(newY, array[z][y])
// 			}
// 		}
// 		if plus != 0 {
// 			newY = append(newY, make([]bool, yLen+plus))
// 		}
// 		newZ = append(newZ, newY)
// 	}
// 	newZ = append(newZ, day17CreateBlank2D(zLen, yLen+plus, xLen+plus))

// 	return newZ
// }

// type day17Coordinates struct {
// 	X int
// 	Y int
// 	Z int
// }

// func (array energySource) day17FindValidPositions(z, y, x int) []day17Coordinates {
// 	log.Debugf("Position at: z:%d y:%d x:%d", z, y, x)
// 	zLen := len(array)
// 	yLen := len(array[0])
// 	xLen := len(array[0][0])

// 	positions := []day17Coordinates{
// 		// Same plane
// 		{x - 1, y - 1, z},
// 		{x, y - 1, z},
// 		{x + 1, y - 1, z},
// 		{x - 1, y, z},
// 		{x + 1, y, z},
// 		{x - 1, y + 1, z},
// 		{x, y + 1, z},
// 		{x + 1, y + 1, z},

// 		// "Left" plane
// 		{x - 1, y - 1, z - 1},
// 		{x, y - 1, z - 1},
// 		{x + 1, y - 1, z - 1},
// 		{x - 1, y, z - 1},
// 		{x + 1, y, z - 1},
// 		{x - 1, y + 1, z - 1},
// 		{x, y + 1, z - 1},
// 		{x + 1, y + 1, z - 1},
// 		{x, y, z - 1},

// 		// "Right" plane
// 		{x - 1, y - 1, z + 1},
// 		{x, y - 1, z + 1},
// 		{x + 1, y - 1, z + 1},
// 		{x - 1, y, z + 1},
// 		{x + 1, y, z + 1},
// 		{x - 1, y + 1, z + 1},
// 		{x, y + 1, z + 1},
// 		{x + 1, y + 1, z + 1},
// 		{x, y, z + 1},
// 	}

// 	validPositions := []day17Coordinates{}
// 	for i, position := range positions {
// 		// log.Debugf("Position check: X: %d Y: %d Z: %d", position.X, position.Y, position.Z)
// 		if (position.X >= 0 && position.X < xLen) && (position.Y >= 0 && position.Y < yLen) && (position.Z >= 0 && position.Z < zLen) {
// 			validPositions = append(validPositions, position)
// 			log.Debugf("%d - %+v", i, position)
// 		}
// 	}

// 	return validPositions
// }

// func (array energySource) day17Count(positions []day17Coordinates) (count int) {
// 	for _, pos := range positions {
// 		if array[pos.Z][pos.Y][pos.X] {
// 			count++
// 		}
// 	}
// 	log.Debugf("Count is: %d", count)
// 	return
// }

// func (array energySource) day17Copy() (replacement energySource) {
// 	for _, z := range array {
// 		newZ := [][]bool{}
// 		for _, y := range z {
// 			newY := []bool{}
// 			for _, x := range y {
// 				newY = append(newY, x)
// 			}
// 			newZ = append(newZ, newY)
// 		}
// 		replacement = append(replacement, newZ)
// 	}

// 	return
// }

// func (array energySource) day17Flip() energySource {
// 	replacement := array.day17Copy()

// 	log.Debugf("Replacement length: %d - Array length: %d", len(replacement), len(array))
// 	for zi, z := range array {
// 		for yi, y := range z {
// 			for xi, x := range y {
// 				vp := array.day17FindValidPositions(zi, yi, xi)
// 				count := array.day17Count(vp)
// 				if x && !(count == 2 || count == 3) {
// 					replacement[zi][yi][xi] = false
// 				} else if !x && (count == 3) {
// 					replacement[zi][yi][xi] = true
// 				}
// 			}
// 		}
// 	}

// 	return replacement
// }

// // Day17 entry function
// func Day17(data string) (activeCube int) {
// 	array := energySource{}

// 	re := regexp.MustCompile(`(.)+`)
// 	rows := re.FindAllString(data, -1)

// 	// Takes the initial state and replicates it over the 3x3x3, instead of the 3x3x1
// 	for i := 0; i < 3; i++ {
// 		zAxis := [][]bool{}
// 		for _, row := range rows {
// 			// Try a make here for the initial size while not making it a type?
// 			newRow := []bool{}
// 			for _, col := range row {
// 				if col == '#' {
// 					newRow = append(newRow, true)
// 				} else {
// 					newRow = append(newRow, false)
// 				}
// 			}
// 			zAxis = append(zAxis, newRow)
// 		}
// 		array = append(array, zAxis)
// 	}

// 	// Testing
// 	// array = array.day17Expand()
// 	// replacement := array.day17Flip()
// 	log.Debug("Original")
// 	array.day17Print()
// 	array = array.day17Flip()
// 	log.Debug("Updated")
// 	array.day17Print()
// 	// replacement.day17Print()

// 	// TODO works, I think, uncomment
// 	// for i := 1; i <= 1; i++ {
// 	// 	array = array.day17Expand()
// 	// 	array = array.day17Flip()
// 	// 	log.Debugf("Step: %d", i)
// 	// }

// 	return array.day17CountActive()
// }

func calculateGroup(expression string) (result int) {
	operations := strings.Split(expression, " ")
	t, _ := strconv.ParseInt(operations[0], 0, 64)
	result = int(t)
	nextOp := ""
	for _, operation := range operations {
		val, err := strconv.ParseInt(operation, 0, 64)
		num := int(val)
		if err == nil && nextOp != "" {
			// log.Debugf("Results: %d using operation: %s with currentNumber: %d", result, nextOp, vali)
			switch nextOp {
			case "+":
				result += num
			case "*":
				result *= num
			}
		} else {
			nextOp = operation
		}
	}

	// log.Debugf("Calculated value: %d", result)

	return
}

var groups []string

func day18FindGroups(expression string) string {
	for i := 0; i < len(expression); i++ {
		c := rune(expression[i])
		if c == '(' {
			group := day18FindGroups(expression[i+1:])
			// log.Debugf("Group: %+v", group)
			i = i + 1 + len(group)
			// Only appends subgroups rather than parent groups as well to reduce the group calculations and future checks.
			if !strings.Contains(group, "(") {
				groups = append(groups, group)
			}
		} else if c == ')' {
			return expression[:i]
		}
	}

	return expression
}

// Day18Wrapper entry function
func Day18Wrapper(expressions []string) (result int) {
	for _, expression := range expressions {
		result += Day18(expression, false)
	}
	return
}

func calculateGroupAdvanced(expression string) (result int) {
	initial := expression
	log.Infof("Initial expression: %s", initial)
	for strings.Contains(expression, " + ") {
		operations := strings.Split(expression, " ")
		t, _ := strconv.ParseInt(operations[0], 0, 64)
		result = int(t)
		nextOp := ""
		for i, operation := range operations {
			val, err := strconv.ParseInt(operation, 0, 64)
			if err == nil && nextOp == "+" {
				// log.Debugf("Results: %d using operation: %s with currentNumber: %d", result, nextOp, vali)
				vali := int(val)
				prev, _ := strconv.ParseInt(operations[i-2], 0, 64)
				previ := int(prev)
				rep := operations[i-2] + " + " + operation
				pv := strconv.Itoa(previ + vali)
				expression = strings.Replace(expression, rep, pv, 1)
				break
			} else {
				nextOp = operation
			}
		}
	}

	log.Infof("After +'s expression: %s", expression)

	operations := strings.Split(expression, " ")
	t, _ := strconv.ParseInt(operations[0], 0, 64)
	result = int(t)
	nextOp := ""
	for _, operation := range operations {
		val, err := strconv.ParseInt(operation, 0, 64)
		if err == nil && nextOp == "*" {
			vali := int(val)
			result *= vali
		} else {
			nextOp = operation
		}
	}

	log.Infof("expression is: %s with result: %d", expression, result)

	return
}

// Day18 entry function
func Day18(expression string, advanced bool) (result int) {
	log.Debugf("Main expression: %s", expression)
	expression = "(" + expression + ")"

	var arithmeticTypeFunction func(string) int
	if advanced {
		arithmeticTypeFunction = calculateGroupAdvanced
	} else {
		arithmeticTypeFunction = calculateGroup
	}

	for true {
		groups = []string{}
		day18FindGroups(expression)
		if len(groups) == 0 {
			log.Debugf("Current expression: %s and result: %d", expression, result)
			return
		}

		// log.Debugf("Groups: %+v", groups)
		for _, group := range groups {
			result = arithmeticTypeFunction(group)
			expression = strings.Replace(expression, "("+group+")", strconv.Itoa(result), 1)
		}
	}

	return
}

// Day18WrapperPart2 entry function
func Day18WrapperPart2(expressions []string) (result int) {
	for _, expression := range expressions {
		result += Day18(expression, true)
	}
	return
}

type day19RuleMap map[string]string

func (rulemap day19RuleMap) day19GetRules(rule string) (s string) {
	if rulemap[rule] == `"a"` || rulemap[rule] == `"b"` || rulemap[rule] == `"R1"` || rulemap[rule] == `"R2"` {
		return strings.Trim(rulemap[rule], `"`)
	}

	options := strings.Split(rulemap[rule], "|")
	s = "("
	for i, k := range options {
		suboptions := strings.Split(strings.Trim(k, " "), " ")
		for _, l := range suboptions {
			s += rulemap.day19GetRules(l)
		}
		if i != len(options)-1 {
			s += "|"
		}
	}
	s += ")"

	return
}

// Day19 entry function
func Day19(data string, ruleChange bool) (validMessages int) {
	split := strings.Split(data, "\n\n")
	rules := strings.Split(split[0], "\n")
	rulemap := day19RuleMap{}

	for _, rule := range rules {
		s := strings.Split(rule, ":")
		rulemap[s[0]] = strings.Trim(s[1], " ")
	}

	// As per part 2 two rules are being replaced.
	// 8: 42 | 42 8
	// 11: 42 31 | 42 11 31
	if ruleChange {
		rulemap["8"] = `42 | 42 8R`
		rulemap["11"] = `42 31 | 42 11R 31`
		rulemap["8R"] = `"R1"`
		rulemap["11R"] = `"R2"`
	}

	rev := rulemap.day19GetRules("0")
	log.Infof("Regex: %s", rev)

	re2 := regexp.MustCompile("^" + rev + "$")
	validMessages = 0
	for _, v := range strings.Split(split[1], "\n") {
		kl := re2.FindAllString(v, -1)
		log.Debug(kl)
		if len(kl) > 0 {
			validMessages++
			log.Info(v)
		}
	}

	return
}

type day20Tile struct {
	id           string
	image        [][]rune
	fingerprints map[string]day20RotationFingerprint
}

type day20RotationFingerprint struct {
	rotation     string
	rotatedImage [][]rune
	top          fingerprint
	bottom       fingerprint
	left         fingerprint
	right        fingerprint
}

type day20RotationLock struct {
	from day20Tile
	to   day20Tile
}

type day20Tiles []day20Tile
type day20Image [][]day20Tile
type fingerprintInput []rune
type fingerprint int

func (tile day20Tile) day20Print() {
	for _, y := range tile.image {
		for _, x := range y {
			fmt.Printf("%c", x)
		}
		fmt.Println()
	}
}

func (image day20Image) Print() {
	for imageY := 0; imageY < len(image); imageY++ {
		for tileY := 0; tileY < len(image[imageY][0].image); tileY++ {
			for imageX := 0; imageX < len(image[0]); imageX++ {
				for tileX := 0; tileX < len(image[imageY][imageX].image[tileY]); tileX++ {
					fmt.Printf("%c", image[imageY][imageX].image[tileY][tileX])
				}
				fmt.Print(" | ")
			}
			fmt.Println()
		}
		fmt.Println("---")
	}
}

func (tile day20Tile) day20RotateRight() (new day20Tile) {
	new.id = tile.id
	new.fingerprints = tile.fingerprints

	for x := 0; x < len(tile.image[0]); x++ {
		newX := []rune{}
		for y := len(tile.image) - 1; y >= 0; y-- {
			t := tile.image[y][x]
			newX = append(newX, t)
		}
		new.image = append(new.image, newX)
	}

	return
}

func (tile day20Tile) day20FlipHorizontal() (new day20Tile) {
	new.id = tile.id
	new.fingerprints = tile.fingerprints

	for yi, y := range tile.image {
		newY := []rune{}
		for x := len(y) - 1; x >= 0; x-- {
			newY = append(newY, tile.image[yi][x])
		}
		new.image = append(new.image, newY)
	}

	return
}

func (tile day20Tile) day20FlipVertical() (new day20Tile) {
	new.id = tile.id
	new.fingerprints = tile.fingerprints

	for y := len(tile.image) - 1; y >= 0; y-- {
		newY := []rune{}
		for x := 0; x < len(tile.image[0]); x++ {
			newY = append(newY, tile.image[y][x])
		}
		new.image = append(new.image, newY)
	}

	return
}

func (input fingerprintInput) day20GenerateFingerprint() (fingerprint int) {
	// log.Debugf("input: %+v - len of input: %d", input, len(input)-1)
	for index, pow := len(input)-1, 1; index >= 0; index, pow = index-1, pow*2 {
		if input[index] == '#' {
			fingerprint += pow
		}
	}

	// log.Debugf("Outcome: %d", fingerprint)

	return
}

func (tile day20Tile) day20TileFingerprintInputData(rotation string) {
	// Left column of tile
	left := []rune{}
	for y := range tile.image {
		left = append(left, tile.image[y][0])
	}

	// Right column of tile
	right := []rune{}
	end := len(tile.image[0]) - 1
	for y := range tile.image {
		right = append(right, tile.image[y][end])
	}

	rotatedImage := make([][]rune, len(tile.image))
	for i, y := range tile.image {
		newY := make([]rune, len(y))
		for j, x := range y {
			newY[j] = x
		}
		rotatedImage[i] = newY
	}

	fingerprints := day20RotationFingerprint{
		rotation:     rotation,
		rotatedImage: rotatedImage,
		top:          fingerprint(fingerprintInput(tile.image[0]).day20GenerateFingerprint()),
		bottom:       fingerprint(fingerprintInput(tile.image[len(tile.image)-1]).day20GenerateFingerprint()),
		left:         fingerprint(fingerprintInput(left).day20GenerateFingerprint()),
		right:        fingerprint(fingerprintInput(right).day20GenerateFingerprint()),
	}

	tile.fingerprints[rotation] = fingerprints

	// log.Debugf("tile values: top: %+v bottom: %+v left: %+v right: %+v", tile.image[0], tile.image[len(tile.image)-1], left, right)
}

func (tile day20Tile) day20GenerateRotationData() {
	originalImage := tile.image

	pos := []string{
		"First",
		"Second",
		"Third",
		"Original",
	}

	for _, i := range pos {
		tile = tile.day20RotateRight()
		tile.day20TileFingerprintInputData("Rotate" + i)

		tile = tile.day20FlipHorizontal()
		tile.day20TileFingerprintInputData("HorizontalRotate" + i)
		//Resets position for vertical
		tile = tile.day20FlipHorizontal()
		//
		tile = tile.day20FlipVertical()
		tile.day20TileFingerprintInputData("VerticalRotate" + i)
		//Resets position for vertical
		tile = tile.day20FlipVertical()
		//
		tile = tile.day20FlipVertical()
		tile = tile.day20FlipHorizontal()
		tile.day20TileFingerprintInputData("HorizontalVerticalRotate" + i)
		//Resets position for vertical
		tile = tile.day20FlipHorizontal()
		tile = tile.day20FlipVertical()
	}

	tile.image = originalImage
}

type day20RotationFingerprintConnection struct {
	from day20RotationFingerprint
	to   day20RotationFingerprint
}

// This is used to find a corner
func day20MatchFingerPrint(from, to day20Tile) (fingerprints []day20RotationFingerprintConnection) {
	if from.id == to.id {
		// log.Debug("Trying to compare the same tile. Skipping...")
		return
	}

	log.Debugf("From rotations:")
	for k, fp := range from.fingerprints {
		log.Debugf("Rotation: %s - Top: %d Bottom: %d, Left: %d, Right: %d", k, fp.top, fp.bottom, fp.left, fp.right)
	}

	log.Debugf("To rotations:")
	for k, fp := range to.fingerprints {
		log.Debugf("Rotation: %s - Top: %d Bottom: %d, Left: %d, Right: %d", k, fp.top, fp.bottom, fp.left, fp.right)
	}

	for _, pfingerprint := range from.fingerprints {
		for _, sfingerprint := range to.fingerprints {
			if pfingerprint.left == sfingerprint.right {
				// log.Debugf("Tile %s is connected to tile %s from left to right.", from.id, to.id)
				fingerprints = append(fingerprints, day20RotationFingerprintConnection{
					pfingerprint,
					sfingerprint,
				})
			}

			if pfingerprint.right == sfingerprint.left {
				// log.Debugf("Tile %s is connected to tile %s from right to left.", from.id, to.id)
				fingerprints = append(fingerprints, day20RotationFingerprintConnection{
					pfingerprint,
					sfingerprint,
				})
			}

			if pfingerprint.bottom == sfingerprint.top {
				// log.Debugf("Tile %s is connected to tile %s from bottom to top.", from.id, to.id)
				fingerprints = append(fingerprints, day20RotationFingerprintConnection{
					pfingerprint,
					sfingerprint,
				})
			}

			if pfingerprint.top == sfingerprint.bottom {
				// log.Debugf("Tile %s is connected to tile %s from top to bottom.", from.id, to.id)
				fingerprints = append(fingerprints, day20RotationFingerprintConnection{
					pfingerprint,
					sfingerprint,
				})
			}
		}

	}

	return
}

func (tiles day20Tiles) day20Find() (result int) {
	// max := int(math.Sqrt(float64(len(tiles))))
	// log.Debugf("Max size is: %d", max)
	image := day20Image{}
	// for i := 0; i < max; i++ {
	// 	image = append(image, make([]day20Tile, max))
	// }
	//
	// for i := 0; i < len(image); i++ {
	// 	for j := 0; j < len(image[i]); j++ {
	// 		log.Debugf("i=%d j=%d", i, j)
	// 	}
	// }

	// image = append(image, []day20Tile{tiles[0]}) // tile is 0,0
	// tiles = tiles[1:]

	// update := func() {
	//
	// }
	for _, from := range tiles {
		count := 0
		for _, to := range tiles {
			if from.id == "2311" && to.id == "1951" {
				if len(day20MatchFingerPrint(from, to)) > 0 {
					count++
				}
			}
		}
		if count > 0 {
			log.Debugf("Tile: %s has %d connections", from.id, count)
		}
	}

	// total:
	// 	for {
	// 		for ii := 0; ii < len(image); ii++ {
	// 			for jj := 0; jj < len(image[ii]); jj++ {
	// 				current := image[ii][jj]
	// 			out:
	// 				for _, pfingerprint := range current.fingerprints {
	// 					for ni, stile := range tiles {
	// 						for _, sfingerprint := range stile.fingerprints {
	// 							if current.id != stile.id {
	// 								if pfingerprint.left == sfingerprint.right {
	// 									current.image = pfingerprint.rotatedImage
	// 									stile.image = sfingerprint.rotatedImage
	// 									image[ii] = append([]day20Tile{stile}, image[ii]...)
	// 									break out
	// 								}
	//
	// 								if pfingerprint.right == sfingerprint.left {
	// 									current.image = pfingerprint.rotatedImage
	// 									stile.image = sfingerprint.rotatedImage
	// 									image[ii] = append(image[ii], stile)
	// 									break out
	// 								}
	//
	// 								if pfingerprint.bottom == sfingerprint.top {
	// 									current.image = pfingerprint.rotatedImage
	// 									stile.image = sfingerprint.rotatedImage
	// 									// Create new row below (of same length as current) and insert into the right position
	// 									new := make([]day20Tile, len(image[0]))
	// 									new[0] = stile
	// 									image = append(image, new)
	// 									break out
	// 								}
	//
	// 								if pfingerprint.top == sfingerprint.bottom {
	// 									current.image = pfingerprint.rotatedImage
	// 									stile.image = sfingerprint.rotatedImage
	// 									// Create new row below (of same length as current) and insert into the right position
	// 									new := make([]day20Tile, len(image))
	// 									new[0] = stile
	// 									image = append(day20Image{new}, image...)
	// 									break out
	// 								}
	// 							}
	// 						}
	// 						log.Debugf("Len %d ni %d", len(tiles), ni)
	// 						if ni+1 < len(tiles) {
	// 							// Remove the tile from tiles as it's now in the image
	// 							tiles = append(tiles[:ni], tiles[ni+1:]...)
	// 						} else {
	// 							tiles = nil
	// 							break total
	// 						}
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}

	image.Print()

	return
}

// Day20 entry function
func Day20(data string) (val int) {
	tiles := day20Tiles{}

	re := regexp.MustCompile(`Tile (.*):`)
	tilesStr := strings.Split(data, "\n\n")
	for _, tile := range tilesStr {
		tileID := re.FindStringSubmatch(tile)
		imagedata := strings.Split(tile, "\n")
		image := [][]rune{}

		for _, il := range imagedata[1:] { // Skips the Title line which is unneeded atm
			ti := []rune{}
			for _, c := range il[1:] {
				ti = append(ti, c)
			}
			image = append(image, ti)
		}

		// TODO hack due to end of line
		if len(tileID) < 2 {
			break
		}

		t := day20Tile{
			id:           tileID[1],
			image:        image,
			fingerprints: map[string]day20RotationFingerprint{},
		}
		t.day20GenerateRotationData()
		tiles = append(tiles, t)
	}

	return tiles.day20Find()
}

func day22Play(p1, p2 []int64, recursive bool, gamenumber int) ([]int64, []int64) {
	p1History := [][]int64{}
	p2History := [][]int64{}

	for i := 1; len(p1) > 0 && len(p2) > 0; i++ {
		// Take card from deck.
		p1Card := p1[0]
		p1 = p1[1:]
		p2Card := p2[0]
		p2 = p2[1:]

		log.Debugf("Game %d - Round %d...", gamenumber, i)
		log.Debugf("Player 1's deck: %+v", p1)
		log.Debugf("Player 2's deck: %+v", p2)
		log.Debugf("Player 1 plays: %d", p1Card)
		log.Debugf("Player 2 plays: %d", p2Card)

		if recursive {
			// A check to see if same cards same order and then give the victory to player 1.
			for index := range p1History {
				if reflect.DeepEqual(p1History[index], p1) && reflect.DeepEqual(p2History[index], p2) {
					log.Debug("Player 1 automatically wins.")
					return []int64{0}, nil
				}
			}
		}

		// Save pattern in history for later usage in case of infinte loops in the recursion.
		p1History = append(p1History, p1)
		p2History = append(p2History, p2)

		if recursive {
			// If both players have at least as many cards remaining in their deck as the value of the card they just drew, the winner of the round is determined by playing a new game of Recursive Combat (see below).
			if p1Card <= int64(len(p1)) && p2Card <= int64(len(p2)) {
				log.Debug("Playing a sub-game to determine the winner...")

				// Make copies of all data for recursive values
				p1c := make([]int64, p1Card)
				p2c := make([]int64, p2Card)
				for i := 0; i < len(p1c); i++ {
					p1c[i] = p1[i]
				}
				for i := 0; i < len(p2c); i++ {
					p2c[i] = p2[i]
				}

				tp1, tp2 := day22Play(p1c, p2c, recursive, gamenumber+1)

				// Based on the length of the returnned values will determine who will win this round.
				if len(tp1) != 0 {
					p1 = append(p1, p1Card, p2Card)
				} else if len(tp2) != 0 {
					p2 = append(p2, p2Card, p1Card)
				} else {
					log.Error("There was an error.")
					log.Errorf("TP1 %d TP2 %d", len(tp1), len(tp2))
				}

				// As the round winner is determined above we continue to next round to skip the high/low check.
				continue
			}
		}

		// High card for this round.
		if p1Card > p2Card {
			// log.Debug("Player 1 won this round.")
			p1 = append(p1, p1Card, p2Card)
		} else if p2Card > p1Card {
			// log.Debug("Player 2 won this round.")
			p2 = append(p2, p2Card, p1Card)
		}
	}

	return p1, p2
}

// Day22 entry function
func Day22(data string, recursive bool) int {
	playerdecks := map[string][]int64{}

	for _, player := range strings.Split(data, "\n\n") {
		scanner := bufio.NewScanner(strings.NewReader(player))
		scanner.Scan()
		playerName := scanner.Text()
		playerdecks[playerName] = []int64{}

		for scanner.Scan() {
			card, _ := strconv.ParseInt(scanner.Text(), 0, 64)
			playerdecks[playerName] = append(playerdecks[playerName], card)
		}
	}

	p1, p2 := day22Play(playerdecks["Player 1:"], playerdecks["Player 2:"], recursive, 1)
	name := ""
	winner := []int64{}
	if len(p1) != 0 {
		name = "1"
		winner = p1
	} else if len(p2) != 0 {
		name = "2"
		winner = p2
	}

	var score int64
	for i, j := len(winner)-1, 1; i >= 0; i, j = i-1, j+1 {
		score += winner[i] * int64(j)
	}

	log.Infof("Player %s won the game. with a score of: %d", name, score)

	return int(score)
}

// Day23 entry function
func Day23(cupOrder string, moves int) (finalOrder string) {
	order := make([]int, len(cupOrder))
	for index, cup := range cupOrder {
		nn, _ := strconv.ParseInt(string(cup), 0, 64)
		order[index] = int(nn)
	}

	if moves == 10000000 {
		temp := make([]int, 10000000-len(order))
		for i, j := 0, 10; i < len(temp); i, j = i+1, j+1 {
			temp[i] = j
		}
		order = append(order, temp...)
		log.Debugf("Added in a million positions")
	}

	// Setting up as much reusable variables as I can outside of the loop
	batchSize := moves
	if moves == 10000000 {
		batchSize = 5000
	}
	threeCups := make([]int, 3)
	threeCupsIndex := make([]int, 3)

	i := 0
	for outer := 0; outer < moves; outer += batchSize {
		log.Debugf("Batch: %d to %d", outer, outer+batchSize)
		for round := outer; round < outer+batchSize; round++ {
			orderLengthFull := len(order)
			currentCupIndex := i % orderLengthFull
			currentCup := order[currentCupIndex]

			for j := 0; j < 3; j++ {
				threeCupsIndex[j] = (i + j + 1) % orderLengthFull
				threeCups[j] = order[(i+j+1)%orderLengthFull]
			}

			// log.Debugf("Three cups index: %+v - IsSorted? %v", threeCupsIndex, sort.SliceIsSorted(threeCupsIndex, func(i, j int) bool { return threeCupsIndex[i] < threeCupsIndex[j] }))

			if sort.SliceIsSorted(threeCupsIndex, func(i, j int) bool { return threeCupsIndex[i] < threeCupsIndex[j] }) {
				order = append(order[0:threeCupsIndex[0]], order[threeCupsIndex[2]+1:]...)
			} else {
				for _, rc := range threeCups {
					for oindex := range order {
						// TODO need a way to be able to use current indexes with the cups while not looping through order
						if order[oindex] == rc {
							order = append(order[0:oindex], order[oindex+1:]...)
							break
						}
					}
				}
			}

			destinationValue := -1
			destinationCupIndex := -1
			largest := -1
			largestValue := -1
			for index, value := range order {
				if value < currentCup && value > destinationValue {
					// Find the next lowest under currentCup
					destinationCupIndex = index
					destinationValue = value
				}
				if value > largestValue {
					largest = index
					largestValue = value
				}
			}

			if destinationCupIndex == -1 {
				destinationCupIndex = largest
			}
			// log.Debugf("Current cup: %+v - Three Cups: %+v - Order: %+v - Destination Cup Index: %+v (%d)", currentCup, threeCups, order, destinationCupIndex, order[destinationCupIndex])

			// // Move cups to their new locations
			additions := append([]int{}, order[destinationCupIndex+1:]...)
			order = append(order[:destinationCupIndex+1], threeCups...)
			order = append(order, additions...)

			// If the cursor is in a different place recalculate the new index for the currentCup
			if order[currentCupIndex] != currentCup {
				if order[(currentCupIndex+1)%orderLengthFull] == currentCup {
					i += +1
				}
				if order[(currentCupIndex+2)%orderLengthFull] == currentCup {
					i += +2
				}
				if order[(currentCupIndex+3)%orderLengthFull] == currentCup {
					i += +3
				}
			}

			// No matter the next state increments the cursor position
			i += 1
		}
		break
	}

	if moves != 10000000 {
		startingIndex := 0

		for in, o := range order {
			if o == 1 {
				startingIndex = in
			}

		}
		for _, o := range order[startingIndex:] {
			finalOrder += strconv.Itoa(int(o))
		}
		for _, o := range order[:startingIndex] {
			finalOrder += strconv.Itoa(int(o))
		}

		finalOrder = strings.TrimLeft(finalOrder, "1")
	} else {
		for in, o := range order {
			if o == 1 {
				leng := len(order)
				two := in % leng
				three := in % leng
				finalOrder = strconv.Itoa(two * three)
				log.Debugf("Broken:   Data is: +1: %d (%d), +2: %d (%d)", two, order[two], three, order[three])
			}
		}
	}

	return
}
