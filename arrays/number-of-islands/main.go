package main

func NumberOfIslands(grid [][]int) int {
	if len(grid) == 0 && len(grid[0]) == 0 {
		return 0
	}

	n := 0
	rows, cols := len(grid), len(grid[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 1 {
				DFS(grid, col, row, rows, cols)
				n++
			}
		}
	}

	return n
}

func DFS(grid [][]int, x, y int, rows, cols int) {
	if x < 0 || x >= rows || y < 0 || y >= cols || grid[x][y] != 1 {
		return
	}

	grid[y][x] = 0

	DFS(grid, x+1, y, rows, cols)
	DFS(grid, x-1, y, rows, cols)
	DFS(grid, x, y+1, rows, cols)
	DFS(grid, x, y-1, rows, cols)
}
