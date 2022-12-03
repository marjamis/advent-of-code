package advent2022

func calculatePriority(value int) (priority int) {
	if value >= 97 {
		// Converts lowercase chars to the correct priority
		priority = value - 96
	} else {
		// Converts uppercase chars to the correct priority
		priority = value - 38
	}

	return
}

func createRucksackMap(rucksack string) (rucksackMap map[rune]int) {
	rucksackMap = map[rune]int{}
	for _, item := range rucksack {
		rucksackMap[item] = 0
	}

	return
}

// Day3Part1 return the total priority value across all rucksacks
func Day3Part1(rucksacks []string) (total int) {
	for _, rucksack := range rucksacks {
		compartementSize := len(rucksack) / 2
		compartmentOne := rucksack[0:compartementSize]
		compartmentTwo := rucksack[compartementSize:]

		compartmentOneMap := createRucksackMap(compartmentOne)
		for _, value := range compartmentTwo {
			if _, containsKey := compartmentOneMap[value]; containsKey {
				total += calculatePriority(int(value))
				break
			}
		}
	}

	return
}

// Day3Part2 returns the total of the priorities of the badge per elf group
func Day3Part2(rucksacks []string) (total int) {
	for rucksack := 0; rucksack < len(rucksacks); rucksack += 3 {
		elf1Rucksack := createRucksackMap(rucksacks[rucksack])
		elf2Rucksack := createRucksackMap(rucksacks[rucksack+1])
		elf3Rucksack := createRucksackMap(rucksacks[rucksack+2])

		for item := range elf1Rucksack {
			_, elf2RucksackContainsItem := elf2Rucksack[item]
			_, elf3RucksackContainsItem := elf3Rucksack[item]

			// If the item from elf1's rucksack also exists in the groups other two elf's rucksacks this must be the badge
			if elf2RucksackContainsItem && elf3RucksackContainsItem {
				total += calculatePriority(int(item))
				break
			}
		}
	}

	return
}
