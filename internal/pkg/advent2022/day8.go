package advent2022

// TreeCoordinates has details about where the tree is position in the grid
type TreeCoordinates struct {
	row int
	col int
}

func checkLeft(treeGrid [][]int, tree *TreeCoordinates) (bool, int) {
	height := treeGrid[tree.row][tree.col]

	for col := tree.col - 1; col >= 0; col-- {
		if !(treeGrid[tree.row][col] < height) {
			return false, tree.col - col
		}
	}

	return true, tree.col
}

func checkRight(treeGrid [][]int, tree *TreeCoordinates) (bool, int) {
	height := treeGrid[tree.row][tree.col]

	for col := tree.col + 1; col < len(treeGrid[tree.row]); col++ {
		if !(treeGrid[tree.row][col] < height) {
			return false, col - tree.col
		}
	}

	return true, (len(treeGrid[tree.row]) - 1) - tree.col
}

func checkUp(treeGrid [][]int, tree *TreeCoordinates) (bool, int) {
	height := treeGrid[tree.row][tree.col]

	for row := tree.row - 1; row >= 0; row-- {
		if !(treeGrid[row][tree.col] < height) {
			return false, tree.row - row
		}
	}

	return true, tree.row
}

func checkDown(treeGrid [][]int, tree *TreeCoordinates) (bool, int) {
	height := treeGrid[tree.row][tree.col]

	for row := tree.row + 1; row < len(treeGrid); row++ {
		if !(treeGrid[row][tree.col] < height) {
			return false, row - tree.row
		}
	}

	return true, (len(treeGrid) - 1) - tree.row
}

// Day8Part1 returns the number of tree positions that can see around it
func Day8Part1(treeGrid [][]int) (visibleTreePositionCount int) {
	// Top and bottom which are visible from outside the grid
	visibleTreePositionCount = len(treeGrid[0]) * 2
	// Sides of middle rows which are visible from outside the grid
	visibleTreePositionCount += (len(treeGrid) - 2) * 2

	for row := 1; row < len(treeGrid)-1; row++ {
		for col := 1; col < len(treeGrid[row])-1; col++ {
			tree := &TreeCoordinates{
				row: row,
				col: col,
			}
			left, _ := checkLeft(treeGrid, tree)
			right, _ := checkRight(treeGrid, tree)
			up, _ := checkUp(treeGrid, tree)
			down, _ := checkDown(treeGrid, tree)

			if left || right || up || down {
				visibleTreePositionCount++
			}
		}
	}

	return visibleTreePositionCount
}

// Day8Part2 returns the highest scenic score a tree position has
func Day8Part2(treeGrid [][]int) (highestScenicScore int) {
	for row := 1; row < len(treeGrid)-1; row++ {
		for col := 1; col < len(treeGrid[row])-1; col++ {
			tree := &TreeCoordinates{
				row: row,
				col: col,
			}
			_, left := checkLeft(treeGrid, tree)
			_, right := checkRight(treeGrid, tree)
			_, up := checkUp(treeGrid, tree)
			_, down := checkDown(treeGrid, tree)

			score := left * right * up * down
			if score > highestScenicScore {
				highestScenicScore = score
			}
		}
	}

	return highestScenicScore
}
