package core

import (
	"errors"
	"math/rand"
)

type Direction struct {
	dr, dc int
}

type Cell struct {
	r, c int
}

func Crossword(words []string, size int) ([][]string, error) {
	// Create a grid of size x size with all cells empty
	grid := make([][]string, size)
	for i := 0; i < size; i++ {
		grid[i] = make([]string, size)
		for j := 0; j < size; j++ {
			grid[i][j] = "_"
		}
	}

	// Define the directions in which words can be placed
	directions := []Direction{
		{0, 1},  // right
		{1, 0},  // down
		{1, 1},  // diagonal down-right
		{-1, 1}, // diagonal up-right
	}

	// Create a list of every cell in the grid
	everyCell := make([]Cell, 0)
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			everyCell = append(everyCell, Cell{r, c})
		}
	}

	placedCount := 0
	// Place each word in the grid
	for _, word := range words {
		wordlen := len(word)

		// Shuffle the list of cells
		rand.Shuffle(len(everyCell), func(i, j int) {
			everyCell[i], everyCell[j] = everyCell[j], everyCell[i]
		})

		// Shuffle the list of directions
		rand.Shuffle(len(directions), func(i, j int) {
			directions[i], directions[j] = directions[j], directions[i]
		})

		// Try to place the word in a random cell with a random direction
	cellLoop:
		for _, cell := range everyCell {
			r, c := cell.r, cell.c
			for _, dir := range directions {
				canPlace := true
				for i := 0; i < wordlen; i++ {
					newR, newC := r+dir.dr*i, c+dir.dc*i
					if newR < 0 || newR >= size || newC < 0 || newC >= size {
						canPlace = false
						break
					}
					currentLetter := string(word[i])
					gridLetter := grid[newR][newC]

					if gridLetter != "_" && gridLetter != currentLetter {
						canPlace = false
						break
					}
				}
				if canPlace {
					for i := 0; i < wordlen; i++ {
						newR, newC := r+dir.dr*i, c+dir.dc*i
						grid[newR][newC] = string(word[i])
					}
					placedCount++
					break cellLoop
				}
			}
		}
	}

	if placedCount < len(words) {
		return nil, errors.New("no se pudieron colocar todas las palabras en el crucigrama")
	}

	// Fill the empty cells with random letters
	letters := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	for r := 0; r < size; r++ {
		for c := 0; c < size; c++ {
			if grid[r][c] == "_" {
				letterIndex := rand.Intn(len(letters))
				grid[r][c] = string(letters[letterIndex])
			}
		}
	}

	return grid, nil
}
