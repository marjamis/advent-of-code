package advent2022

func findStartingIndex(message string, numberOfDistinctCharacters int) (startIndex int) {
	for index := range message {
		messageSegment := map[byte]int{}

		for i := 0; i < numberOfDistinctCharacters; i++ {
			messageSegment[message[index+i]] = 0
		}

		if len(messageSegment) == numberOfDistinctCharacters {
			return index
		}
	}

	return -1
}

// Day6Part1 returns the start of the packet index
func Day6Part1(message string) (startOfPacketIndex int) {
	numberOfDistinctCharactersForPacket := 4

	return findStartingIndex(message, numberOfDistinctCharactersForPacket) + numberOfDistinctCharactersForPacket
}

// Day6Part2 returns the start of the message index
func Day6Part2(message string) (startOfMessageIndex int) {
	numberOfDistinctCharactersForMessage := 14

	return findStartingIndex(message, numberOfDistinctCharactersForMessage) + numberOfDistinctCharactersForMessage
}
