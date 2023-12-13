package advent2023

import (
	"strconv"
	"strings"

	"github.com/marjamis/advent-of-code/pkg/structures"
)

var mapOrderKeys = []string{
	"seed-to-soil",
	"soil-to-fertilizer",
	"fertilizer-to-water",
	"water-to-light",
	"light-to-temperature",
	"temperature-to-humidity",
	"humidity-to-location",
}

type originalMapping struct {
	destination int
	source      int
	distance    int
}

type distance struct {
	start int
	end   int
}

type originalMappings []originalMapping

func extractMappings(mappings []string) (distances originalMappings) {
	for _, mapping := range mappings {
		values := strings.Split(mapping, " ")

		to, err := strconv.Atoi(values[0])
		if err != nil {
			panic("Failed to convert to a string")
		}

		from, err := strconv.Atoi(values[1])
		if err != nil {
			panic("Failed to convert to a string")
		}

		distance, err := strconv.Atoi(values[2])
		if err != nil {
			panic("Failed to convert to a string")
		}

		distances = append(distances, originalMapping{
			destination: to,
			source:      from,
			distance:    distance,
		})
	}

	return
}

func createMappings(data []string) (mappings map[string]originalMappings) {
	mappings = map[string]originalMappings{}

	for _, d := range data {
		lines := strings.Split(d, "\n")
		mapType := strings.Split(lines[0], " ")

		switch mapType[0] {
		case "seed-to-soil":
			mappings["seed-to-soil"] = extractMappings(lines[1:])
		case "soil-to-fertilizer":
			mappings["soil-to-fertilizer"] = extractMappings(lines[1:])
		case "fertilizer-to-water":
			mappings["fertilizer-to-water"] = extractMappings(lines[1:])
		case "water-to-light":
			mappings["water-to-light"] = extractMappings(lines[1:])
		case "light-to-temperature":
			mappings["light-to-temperature"] = extractMappings(lines[1:])
		case "temperature-to-humidity":
			mappings["temperature-to-humidity"] = extractMappings(lines[1:])
		case "humidity-to-location":
			mappings["humidity-to-location"] = extractMappings(lines[1:])
		}
	}

	return
}

func findLocation(mappings originalMappings, sourceLocation int) (location int) {
	for _, mapping := range mappings {
		if (sourceLocation >= mapping.source) && (sourceLocation <= mapping.source+mapping.distance) {
			return mapping.destination + (sourceLocation - mapping.source)
		}
	}

	return sourceLocation
}

func findLocations(locations *structures.Stack, mappings originalMappings) *structures.Stack {
	newLocations := structures.CreateStack()

	for !locations.IsEmpty() {
		if location := locations.Pop(); location != nil {
			anyMatches := false
			seedBatch := location.(distance)

			for _, mapping := range mappings {
				destinationStart := mapping.destination
				destinationEnd := mapping.destination + mapping.distance

				sourceStart := mapping.source
				sourceEnd := mapping.source + mapping.distance

				// Complete overlap
				if (seedBatch.start >= sourceStart) && (seedBatch.end <= sourceEnd) {
					newLocations.Push(distance{
						start: destinationStart + (seedBatch.start - sourceStart),
						end:   destinationStart + (seedBatch.end - sourceStart),
					})

					anyMatches = true
				} else if (seedBatch.end > sourceStart) && (seedBatch.start < sourceStart) {
					// Left partial overlap
					newLocations.Push(distance{
						start: destinationStart,
						end:   destinationStart + (seedBatch.end - sourceStart),
					})

					locations.Push(distance{
						start: seedBatch.start,
						end:   sourceStart - 1,
					})

					anyMatches = true
				} else if (seedBatch.start < sourceEnd) && (sourceStart < seedBatch.start) {
					// Right partial overlap
					newLocations.Push(distance{
						start: destinationStart + (sourceEnd - seedBatch.start),
						end:   destinationEnd,
					})

					locations.Push(distance{
						start: seedBatch.start + (sourceEnd - seedBatch.start),
						end:   seedBatch.end,
					})

					anyMatches = true
				}
			}

			// No overlap
			if !anyMatches {
				newLocations.Push(seedBatch)
			}
		}
	}

	return newLocations
}

// Day5Part1 returns the closest seed location
func Day5Part1(data string) (closestLocationNumber int) {
	sections := strings.Split(data, "\n\n")
	seeds := strings.Fields(sections[0])
	mappings := createMappings(sections[1:])

	closestLocationNumber = 999999999999999999
	for _, seed := range seeds[1:] {
		location, _ := strconv.Atoi(seed)

		for _, key := range mapOrderKeys {
			location = findLocation(mappings[key], location)
		}

		if location < closestLocationNumber {
			closestLocationNumber = location
		}
	}

	return
}

// Day5Part2 returns the closest seed location, based off the seed ranges
func Day5Part2(data string) (closestLocationNumber int) {
	sections := strings.Split(data, "\n\n")
	seeds := strings.Fields(sections[0])
	mappings := createMappings(sections[1:])

	closestLocationNumber = 999999999999999999
	for i := 1; i < len(seeds); i += 2 {
		seedStart, err := strconv.Atoi(seeds[i])
		if err != nil {
			panic("Failed to convert to a string")
		}

		seedRange, err := strconv.Atoi(seeds[i+1])
		if err != nil {
			panic("Failed to convert to a string")
		}

		seedRanges := structures.CreateStack()
		seedRanges.Push(distance{
			start: seedStart,
			end:   seedStart + seedRange,
		})

		for _, key := range mapOrderKeys {
			seedRanges = findLocations(seedRanges, mappings[key])
		}

		for !seedRanges.IsEmpty() {
			location := seedRanges.Pop().(distance)

			if location.start < closestLocationNumber && location.start != 0 {
				closestLocationNumber = location.start
			}
		}
	}

	return
}
