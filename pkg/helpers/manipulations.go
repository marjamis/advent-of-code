package helpers

import "sort"

type ArrayTypes interface {
	int | int32 | int64 | float32 | float64 | ~string
}

type MatrixCoordinates struct {
	Row int
	Col int
}

// Permutations generates all possible combinations from the input data
func Permutations(xs []int16) (permuts [][]int16) {
	// Taken from: https://www.golangprograms.com/golang-program-to-generate-slice-permutations-of-number-entered-by-user.html
	var rc func([]int16, int16)
	rc = func(a []int16, k int16) {
		if k == int16(len(a)) {
			permuts = append(permuts, append([]int16{}, a...))
		} else {
			for i := k; i < int16(len(xs)); i++ {
				a[k], a[i] = a[i], a[k]
				rc(a, k+1)
				a[k], a[i] = a[i], a[k]
			}
		}
	}
	rc(xs, 0)

	return permuts
}

func IsArrayLocationValid[T ArrayTypes](arr [][]T, row, col int) bool {
	return (col >= 0) && (col < len(arr[0])) && (row >= 0) && (row < len(arr))
}

// IsLocationValid returns if the provided x,y coordinates are within the range of the provided 2d array.
//
// Deprecated: The x, y format is confusing when using with arrays as values have to be inversed.
// Use IsArrayLocationValid instead.
func IsLocationValid[T ArrayTypes](arr [][]T, x, y int) bool {
	return (x >= 0) && (x < len(arr[0])) && (y >= 0) && (y < len(arr))
}

func FindSurroundingCoordinates(currentPosition MatrixCoordinates, includeDiagonal bool) []MatrixCoordinates {
	surroundingLocations := []MatrixCoordinates{
		{
			Col: currentPosition.Col,
			Row: currentPosition.Row - 1,
		},
		{
			Col: currentPosition.Col,
			Row: currentPosition.Row + 1,
		},
		{
			Col: currentPosition.Col + 1,
			Row: currentPosition.Row,
		},
		{
			Col: currentPosition.Col - 1,
			Row: currentPosition.Row,
		},
	}

	if includeDiagonal {
		diagonal := []MatrixCoordinates{
			{
				Col: currentPosition.Col - 1,
				Row: currentPosition.Row - 1,
			},
			{
				Col: currentPosition.Col + 1,
				Row: currentPosition.Row + 1,
			},
			{
				Col: currentPosition.Col + 1,
				Row: currentPosition.Row - 1,
			},
			{
				Col: currentPosition.Col - 1,
				Row: currentPosition.Row + 1,
			},
		}
		surroundingLocations = append(surroundingLocations, diagonal...)
	}

	return surroundingLocations
}

// Abs is simple function to return the absolute value of an integer. Absolute value being essentially an always positive number.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// ManhattansDistance return thes Manhattan distance between two points
func ManhattansDistance(x1, y1, x2, y2 int) int {
	return Abs(x1-x2) + Abs(y1-y2)
}

// DecimalPositionOf returns a decimal position that a number can be times against to move the decimal position of any number
func DecimalPositionOf(desiredPosition int) (positionValue int) {
	/*
		Useful if you have an integer of 7 that needs to be in the 3rd decimal poition (i.e. be 700) you would use this as:
			return DecimalPositionOf(3) * 7   // 700
		Other examples:
			return DecimalPositionOf(1) * 7   // 7
			return DecimalPositionOf(2) * 7   // 70
			return DecimalPositionOf(7) * 7   // 7000000
	*/
	positionValue = 1
	for i := 0; i < desiredPosition; i++ {
		positionValue *= 10
	}

	return
}

// Copy2dArray creates a copy of a 2d array
func Copy2dArray[T ArrayTypes](array [][]T) (copied [][]T) {
	for _, i := range array {
		tmp := make([]T, len(i))
		copy(tmp, i)
		copied = append(copied, tmp)
	}

	return
}

// RemoveItemsAtIndexes takes an array of data and an array of indexes, loops
// through the array to remove those indexes, in a safe way, and returns the
// remaining elements
func RemoveItemsAtIndexes[T ArrayTypes](array []T, indexesForRemoval []int) []T {
	// Sorting in reverse means the removals doesn't change the indexes for other removals
	sort.Sort(sort.Reverse(sort.IntSlice(indexesForRemoval)))

	// Go through all the indexes to remove and remove them from the array
	for _, i := range indexesForRemoval {
		array = append(array[0:i], array[i+1:]...)
	}

	return array
}

// RemoveDuplicates takes a []string array and removes any duplicates strings in that array
func RemoveDuplicates(data []string) (uniques []string) {
	present := map[string]bool{}

	for _, d := range data {
		_, ok := present[d]
		if !ok {
			uniques = append(uniques, d)
			present[d] = true
		}
	}

	return uniques
}
