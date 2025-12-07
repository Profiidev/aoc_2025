package internal

import (
	"log"
	"strconv"
	"strings"
)

type cellType int

const (
	empty cellType = iota
	splitter
	beam
)

type cell struct {
	cellType cellType
	count    int
}

func Day07(input string) {
	grid := [][]cell{}

	for line := range strings.SplitSeq(input, "\n") {
		row := []cell{}
		for _, char := range line {
			switch char {
			case '.':
				row = append(row, cell{cellType: empty})
			case '^':
				row = append(row, cell{cellType: splitter})
			case 'S':
				row = append(row, cell{cellType: beam, count: 1})
			}
		}

		grid = append(grid, row)
	}

	total := 0
	updated := true
	splits := 0

	for updated {
		updated, splits, grid = simulateBeams(grid)
		total += splits
	}

	result_2 := 0
	for j := 0; j < len(grid[0]); j++ {
		if grid[len(grid)-1][j].cellType == beam {
			result_2 += grid[len(grid)-1][j].count
		}
	}

	log.Printf("Result: %d", total)
	log.Printf("Result 2: %d", result_2)
}

func simulateBeams(grid [][]cell) (bool, int, [][]cell) {
	updated := false
	splits := 0
	newGrid := make([][]cell, len(grid))
	for i := range grid {
		newGrid[i] = make([]cell, len(grid[i]))
		copy(newGrid[i], grid[i])
	}

	for i := 0; i < len(grid)-1; i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j].cellType == beam {
				current := &newGrid[i][j]
				next := &newGrid[i+1][j]

				switch grid[i+1][j].cellType {
				case empty:
					count := current.count + next.count
					*next = cell{cellType: beam, count: count}
					*current = cell{cellType: empty}
					updated = true
				case splitter:
					updated = true
					splits++
					if j > 0 {
						count := current.count + newGrid[i+1][j-1].count
						newGrid[i+1][j-1] = cell{cellType: beam, count: count}
					}
					if j < len(grid[i])-1 {
						count := current.count + newGrid[i+1][j+1].count
						newGrid[i+1][j+1] = cell{cellType: beam, count: count}
					}
					*current = cell{cellType: empty}
				}
			}
		}
	}

	return updated, splits, newGrid
}

func printGrid(grid *[][]cell) {
	for i := 0; i < len(*grid); i++ {
		row := ""
		for j := 0; j < len((*grid)[i]); j++ {
			switch (*grid)[i][j].cellType {
			case empty:
				row += "."
			case splitter:
				row += "^"
			case beam:
				row += strconv.Itoa((*grid)[i][j].count)
			}
		}
		log.Println(row)
	}
}
