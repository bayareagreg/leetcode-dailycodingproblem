package main

import "fmt"

func numIslands(grid [][]byte) int {
	count := 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '1' {
				count++
				// now find all adjacent land and set it to
				// something else to avoid counting twice
				findAdjacentLand(grid, i, j)
			}
		}
	}
	return count
}

func findAdjacentLand(grid [][]byte, row, col int) {
	if 0 <= row && row < len(grid) &&
		0 <= col && col < len(grid[0]) &&
		grid[row][col] == '1' {
		grid[row][col] = '0'
		findAdjacentLand(grid, row-1, col)
		findAdjacentLand(grid, row+1, col)
		findAdjacentLand(grid, row, col-1)
		findAdjacentLand(grid, row, col+1)
	}
}

func main() {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '1', '0', '0'},
		{'0', '0', '0', '1', '1'}}

	fmt.Println(numIslands(grid))
}
