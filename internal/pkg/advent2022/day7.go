package advent2022

import (
	"sort"
	"strconv"
	"strings"
)

const (
	part1SmallerThan = 100000
	filesystemSize   = 70000000
	updateRequires   = 30000000
)

// DirectoryDetails contains all the data required for a directory
type DirectoryDetails struct {
	files       map[string]*FileDetails
	directories map[string]*DirectoryDetails
	parent      *DirectoryDetails
}

// FileDetails contains all the data required for a file
type FileDetails struct {
	size int
}

func wrapper(structure *DirectoryDetails) (total int, directorySizes []int) {
	// Defining the variable first so the variable is defined before being initialised AND allows the closure to call itself
	var walk func(root *DirectoryDetails) int
	walk = func(root *DirectoryDetails) int {
		totalSize := 0

		// Loop through each directory and recursively calculate their size and return it
		for _, value := range root.directories {
			totalSize += walk(value)
		}

		// Loop through all the files of the current directory to get the total size of the files
		for _, value := range root.files {
			totalSize += value.size
		}

		// For part 1 calculating the total for directories smaller than part1SmallerThank
		if totalSize <= part1SmallerThan {
			total += totalSize
		}

		directorySizes = append(directorySizes, totalSize)

		return totalSize
	}

	walk(structure)

	return total, directorySizes
}

func buildTree(terminalOutput []string) (root *DirectoryDetails) {
	root = &DirectoryDetails{
		parent:      nil,
		files:       nil,
		directories: nil,
	}
	currentDirectory := root

	// TODO this likely should be broken into multiple functions
	for index := 0; index < len(terminalOutput); index++ {
		line := terminalOutput[index]

		// Checks if it's a command or not. This isn't really required with the current logic but doesn't hurt to check
		if line[0] == '$' {
			command := strings.Split(line, " ")[1]
			switch command {
			case "cd":
				dir := strings.Split(line, " ")[2]

				switch dir {
				case "/":
					// move pointer to root
					currentDirectory = root
				case "..":
					// move pointer up one
					currentDirectory = currentDirectory.parent
				default:
					// move pointer to directory specified
					currentDirectory = currentDirectory.directories[dir]
				}
			case "ls":
				subIndex := index + 1
				for subIndex < len(terminalOutput) {
					displayedLine := terminalOutput[subIndex]
					displayedLineSplit := strings.Split(displayedLine, " ")

					// Checks if the new displayedLine is a command which indicates the loop should be broken out of
					if displayedLineSplit[0] == "$" {
						index = subIndex - 1
						break
					}

					// Creates a dir if the dir keyword is specified
					if displayedLineSplit[0] == "dir" {

						// Checks if the specified directory already exists in the current directory
						if _, ok := currentDirectory.directories[displayedLineSplit[1]]; !ok {
							// Initialises the directories listing of the current directory if it hasn't been
							if currentDirectory.directories == nil {
								currentDirectory.directories = map[string]*DirectoryDetails{}
							}
							// Creates the new directory
							currentDirectory.directories[displayedLineSplit[1]] = &DirectoryDetails{
								parent: currentDirectory,
							}
						}
					} else {
						// Initialises the files listing of the current directory if it hasn't been
						if currentDirectory.files == nil {
							currentDirectory.files = map[string]*FileDetails{}
						}
						// Creates the new file with the provided size
						size, _ := strconv.Atoi(displayedLineSplit[0])
						currentDirectory.files[displayedLineSplit[1]] = &FileDetails{
							size: size,
						}
					}
					subIndex++
				}
			}
		}
	}

	return
}

// Day7Part1 returns the total of all the directories smaller than the value defined in the const part1SmallerThan
func Day7Part1(terminalOutput []string) (totalOfDirectoriesSmallerThan int) {
	root := buildTree(terminalOutput)
	totalOfDirectoriesSmallerThan, _ = wrapper(root)

	return totalOfDirectoriesSmallerThan
}

// Day7Part2 returns the optimal directory to delete which is the smallest directory while covering the additional space that needs to be deleted
func Day7Part2(terminalOutput []string) (optimalDirectoryToDelete int) {
	root := buildTree(terminalOutput)

	_, directorySizes := wrapper(root)
	sort.Sort(sort.IntSlice(directorySizes))

	// Takes the largest directory, which would be root as it encompasses everything, as the total used space
	usedSpace := directorySizes[len(directorySizes)-1]

	// Calculates how much space is unusedSpace in the filesystem
	unusedSpace := filesystemSize - usedSpace

	// Calculates how much additional space is requiredSpace on top of what is current available
	requiredSpace := updateRequires - unusedSpace

	// Finds the smallest directory than can be deleted which provides the required space
	optimalDirectoryToDelete = 0
	for i := 0; i < len(directorySizes); i++ {
		if directorySizes[i] >= requiredSpace {
			optimalDirectoryToDelete = directorySizes[i]
			break
		}
	}

	return optimalDirectoryToDelete
}
