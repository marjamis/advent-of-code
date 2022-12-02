package helpers

// GetArrayIndexForValue will loop through an array to find the index that matches the given value. This assume only one element will have this value
func GetArrayIndexForValue(array []interface{}, valueToFind interface{}) int {
	for index, value := range array {
		if value == valueToFind {
			return index
		}
	}

	return -1
}

// GetValidAdjacentIndex Quick and dirty function to allow the array to be wrapped around if the adjacentShapeIndex goes above or below the arrays boundaries
func GetValidAdjacentIndex(array []interface{}, valueToFind interface{}, higher bool) (newIndex int) {
	valueToFindIndex := GetArrayIndexForValue(array, valueToFind)

	if higher {
		newIndex = valueToFindIndex + 1
		if newIndex > len(array)-1 {
			newIndex = 0
		}
	} else {
		newIndex = valueToFindIndex - 1
		if newIndex < 0 {
			newIndex = len(array) - 1
		}
	}

	return
}
