package game_of_life_simple

import ()

type Coord struct {
	x int
	y int
}

func runGame(grid [][]int, iters int) [][]int {
	for i := 0; i < iters; i++ {
		grid = nextGen(grid)
	}
	return grid
}

func nextGen(grid [][]int) [][]int {
	newGrid := make([][]int, len(grid))
	for i := range newGrid {
		newGrid[i] = make([]int, len(grid[0]))
	}
	for i := range grid {
		for j := range grid[i] {
			newGrid[i][j] = nextStatus(grid[i][j],
				countNeighbors(grid, Coord{i, j}))
		}
	}
	return newGrid
}

func countNeighbors(grid [][]int, loc Coord) int {
	offsets := []int{-1, 0, 1}
	count := 0
	for _, dx := range offsets {
		x := loc.x + dx
		if x < 0 || x >= len(grid) {
			continue
		}
		for _, dy := range offsets {
			if dx == 0 && dy == 0 {
				continue
			}
			y := loc.y + dy
			if y < 0 || y >= len(grid[0]) {
				continue
			}
			count += grid[x][y]
		}
	}
	return count
}

func nextStatus(currStatus, numNeighbors int) int {
	if numNeighbors == 3 || (currStatus == 1 && numNeighbors == 2) {
		return 1
	}
	return 0
}
