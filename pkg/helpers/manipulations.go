package helpers

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

// IsLocationValid returns if the provided x,y coordinates are within the range of the provided 2d array.
func IsLocationValid(arr [][]rune, x, y int) bool {
	//TODO change this to be a generic 2d array rather than rune specifically
	return (x >= 0) && (x < len(arr[0])) && (y >= 0) && (y < len(arr))
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

// Copy2dInt creates a copy of a 2d array
// TODO convert to a generic in the future with an applicable golang version
func Copy2dInt(array [][]int) (copied [][]int) {
	for _, i := range array {
		tmp := make([]int, len(i))
		copy(tmp, i)
		copied = append(copied, tmp)
	}

	return
}
