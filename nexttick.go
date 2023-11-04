package coderetreat2023

import "fmt"

func validateInput(grid [][]bool) (success bool) {
	rowCount := len(grid)

	if rowCount < 1 {
		fmt.Println("Grid too small")
		return false
	}

	colCount := len(grid[0])

	if rowCount != colCount {
		fmt.Println("Grid not square")
		return false
	}
	return true
}

func isLiving(grid [][]bool, rowIndex int, colIndex int) (living bool) {
	lengthOfGrid := len(grid)
	if rowIndex >= 0 && colIndex >= 0 && rowIndex < lengthOfGrid && colIndex < lengthOfGrid {
		return grid[rowIndex][colIndex]
	}

	return false
}
func getLivingNeighboursCount(grid [][]bool, rowIndex int, colIndex int) (count int) {
	count = 0

	for r := -1; r <= 1; r++ {
		for c := -1; c <= 1; c++ {
			if r != 0 && c != 0 && isLiving(grid, rowIndex+r, colIndex+c) {
				count++
			}
		}
	}

	return count
}

func getPossibleNeighboursCount(rowIndex int, colIndex int, size int) (possibleNeighboursCount int) {
	possibleNeighboursCount = 8
	if rowIndex == 0 || rowIndex == size-1 {
		possibleNeighboursCount -= 3
	}

	if colIndex == 0 || colIndex == size-1 {
		possibleNeighboursCount -= 3
	}

	if rowIndex == 0 && colIndex == 0 ||
		rowIndex == size-1 && colIndex == 0 ||
		rowIndex == 0 && colIndex == size-1 ||
		rowIndex == size-1 && colIndex == size-1 {
		possibleNeighboursCount += 1
	}

	return possibleNeighboursCount
}

func getDeadNeighboursCount(grid [][]bool, rowIndex int, colIndex int) (count int) {
	return getPossibleNeighboursCount(rowIndex, colIndex, len(grid)) - getLivingNeighboursCount(grid, rowIndex, colIndex) // TODO - fix for bounds
}

func shouldMarkLiving(liveNeighbours int, deadNeighbours int) (living bool) {
	if liveNeighbours == 2 || liveNeighbours == 3 || deadNeighbours == 3 {
		return true
	}

	return false
}

func nextTick(grid [][]bool) (outputGrid [][]bool, success bool) {
	success = validateInput(grid)
	if !success {
		return outputGrid, false
	}

	outputGrid = make([][]bool, len(grid))
	for r, rows := range grid {
		outputGrid[r] = make([]bool, len(rows))
		for c := range rows {
			liveNeighbours := getLivingNeighboursCount(grid, r, c)
			deadNeighbours := getDeadNeighboursCount(grid, r, c)
			outputGrid[r][c] = shouldMarkLiving(liveNeighbours, deadNeighbours)
		}
	}

	return outputGrid, true
}
