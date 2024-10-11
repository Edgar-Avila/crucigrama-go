package core

import "math/rand"

func MatrixEffectInit(grid [][]string) [][]string {
	newGrid := make([][]string, len(grid))
	for i := range grid {
		newGrid[i] = make([]string, len(grid[i]))
		copy(newGrid[i], grid[i])
	}

	rows := len(grid)
	cols := 0
	if rows > 0 {
		cols = len(grid[0])
	}

	// Delete some values for each column
	for i := 0; i < cols; i++ {
		start := rand.Intn(rows)
		n := rand.Intn(rows / 3) + rows / 2

		for j := start; j < start+n; j++ {
			newGrid[j%rows][i] = " "
		}
	}

	return newGrid
}

func MatrixEffectNext(grid [][]string) [][]string {
	lastRowIdx := len(grid) - 1
	lastRow := grid[lastRowIdx]
	newGrid := make([][]string, len(grid))
	newGrid[0] = make([]string, len(lastRow))
	copy(newGrid[0], lastRow)

	for i := 1; i < len(grid); i++ {
		newGrid[i] = make([]string, len(grid[i]))
		copy(newGrid[i], grid[i-1])
	}

	return newGrid
}
