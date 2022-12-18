package advent2022

import (
	"sort"
	"strconv"
	"strings"

	"github.com/marjamis/advent-of-code/pkg/helpers"
)

// Monkey hold keys information about the monkeys who take and throw items
type Monkey struct {
	items           []int
	operation       string
	operationValue  string
	testValue       int
	trueMonkey      int
	falseMonkey     int
	inspectionCount int
}

func createMonkey(monkeyDetails []string) *Monkey {
	testValue, err := strconv.Atoi(strings.Split(monkeyDetails[3], " by ")[1])
	if err != nil {
		return nil
	}

	trueMonkeyValue, _ := strconv.Atoi(strings.Split(monkeyDetails[4], " ")[9])
	if err != nil {
		return nil
	}

	falseMonkeyValue, _ := strconv.Atoi(strings.Split(monkeyDetails[5], " ")[9])
	if err != nil {
		return nil
	}

	items := []int{}
	for _, item := range strings.Split(strings.Split(monkeyDetails[1], ": ")[1], ", ") {
		val, err := strconv.Atoi(item)
		if err != nil {
			return nil
		}

		items = append(items, val)
	}

	return &Monkey{
		items:          items,
		operation:      strings.Split(strings.Split(monkeyDetails[2], ": ")[1], " ")[3],
		operationValue: strings.Split(strings.Split(monkeyDetails[2], ": ")[1], " ")[4],
		testValue:      testValue,
		trueMonkey:     trueMonkeyValue,
		falseMonkey:    falseMonkeyValue,
	}
}

func monkeyBusiness(monkeyDetails []string, part2 bool) int {
	monkeys := map[int]*Monkey{}
	// 6 lines for a monkey then +1 for the blank spacer line
	for monkey := 0; monkey < len(monkeyDetails); monkey += 7 {
		monkeys[len(monkeys)] = createMonkey(monkeyDetails[monkey : monkey+6])
	}

	// As this is a map order is important which we sort here. Could use an array but...
	var sortedMonkeyKeys []int
	for k := range monkeys {
		sortedMonkeyKeys = append(sortedMonkeyKeys, k)
	}
	sort.Sort(sort.IntSlice(sortedMonkeyKeys))

	rounds := 20

	// Used so can be referenced if part2 but doesn't run through any unneeded calculations for part 1
	var lcm int
	if part2 {
		rounds = 10000

		// For Part 2, as there is no division by 3 the numbers become too large to calculate
		// To overcome this we use the LCM of all the monkeys divisions as the basis of where the item should go
		monkeysTestValues := []int{}
		for _, monkey := range monkeys {
			monkeysTestValues = append(monkeysTestValues, monkey.testValue)
		}
		lcm = helpers.LCM(monkeysTestValues[0], monkeysTestValues[1], monkeysTestValues[2:]...)
	}

	for round := 1; round <= rounds; round++ {
		for _, value := range sortedMonkeyKeys {
			monkey := monkeys[value]

			for _, item := range monkey.items {
				monkey.inspectionCount++
				worry := item

				value, err := strconv.Atoi(monkey.operationValue)
				if err != nil {
					value = worry
				}

				switch monkey.operation {
				case "+":
					worry += value
				case "*":
					worry *= value
				}

				if !part2 {
					worry /= 3
				} else {
					worry %= lcm
				}

				var throwToMonkey int
				if (worry % int(monkey.testValue)) == 0 {
					throwToMonkey = monkey.trueMonkey
				} else {
					throwToMonkey = monkey.falseMonkey
				}

				// Adds the item to new monkey
				monkeys[throwToMonkey].items = append(monkeys[throwToMonkey].items, worry)
			}
			// By the end of the run there will be no more items, so emptying
			monkey.items = []int{}
		}
	}

	// Goes through all the monkeys and collects their inspection counts
	totals := []int{}
	for key := range sortedMonkeyKeys {
		totals = append(totals, monkeys[key].inspectionCount)
	}
	sort.Ints(totals)

	return totals[len(totals)-1] * totals[len(totals)-2]
}

// Day11Part1 returns the level of monkey business after 20 rounds
func Day11Part1(monkeyDetails []string) int {
	return monkeyBusiness(monkeyDetails, false)
}

// Day11Part2 returns the level of monkey business after 10,000 rounds
func Day11Part2(monkeyDetails []string) int {
	return monkeyBusiness(monkeyDetails, true)
}
