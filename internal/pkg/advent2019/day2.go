package advent2019

// IntcodeComputer runs the provided intcode and returns the updated intcode after computation
func IntcodeComputer(intcode ...int) []int {
	address := 0

outerLoop:
	for true {
		switch intcode[address] {
		case 1:
			//Addition
			intcode[intcode[address+3]] = intcode[intcode[address+1]] + intcode[intcode[address+2]]
			address = address + 4
		case 2:
			//Multiplication
			intcode[intcode[address+3]] = intcode[intcode[address+1]] * intcode[intcode[address+2]]
			address = address + 4
		case 99:
			//End of app
			fallthrough
		default:
			break outerLoop
		}
	}

	return intcode
}

// Day2Part1 return the first position of the modified intcode after computation
func Day2Part1(intcode ...int) int {
	// This replaces the data as per the instructions
	usedCopy := make([]int, len(intcode))
	copy(usedCopy, intcode)
	usedCopy[1] = 12
	usedCopy[2] = 2

	return IntcodeComputer(usedCopy...)[0]
}

// Day2Part2 returns the noun and verb (address 1 and 2 respectively) times by 100 of the modified intcode after computation
func Day2Part2(intcode ...int) (val int) {
	usedCopy := make([]int, len(intcode))

outerLoop:
	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			copy(usedCopy, intcode)
			usedCopy[1] = i
			usedCopy[2] = j
			test := IntcodeComputer(usedCopy...)
			if test[0] == 19690720 {
				break outerLoop
			}
		}
	}

	return 100*usedCopy[1] + usedCopy[2]
}
