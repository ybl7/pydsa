import (
	dq "github.com/idsulik/go-collections/deque"
)

type Point struct {
	X int
	Y int
}

func shortestPathBinaryMatrix(grid [][]int) int {
	// Check if start or end is blocked
	if grid[0][0] == 1 || grid[len(grid)-1][len(grid[0])-1] == 1 {
		return -1
	}

	nRows := len(grid)
	nCols := len(grid[0])

	if nRows == 1 && nCols == 1 {
		return 1
	}

	visited := make([][]bool, nRows)
	for i := range visited {
		visited[i] = make([]bool, nCols)
	}

	q := dq.New[Point](nRows * nCols)
	q.PushBack(Point{0, 0})
	visited[0][0] = true

	pathLength := 1

	for !q.IsEmpty() {
		size := q.Len()

		for i := 0; i < size; i++ {
			p, _ := q.PopFront()

			// Check if we reached the destination
			if p.X == nRows-1 && p.Y == nCols-1 {
				return pathLength
			}

			// Get valid neighbors
			for _, neighbor := range getValidNeighbors(grid, p, nRows, nCols, visited) {
				q.PushBack(neighbor)
				visited[neighbor.X][neighbor.Y] = true
			}
		}
		pathLength++
	}

	return -1
}

func isInBounds(x, y, rows, cols int) bool {
	return x >= 0 && x < rows && y >= 0 && y < cols
}

func getValidNeighbors(grid [][]int, p Point, rows, cols int, visited [][]bool) []Point {
	// The 8 directions for moving in a grid (including diagonals)
	directions := []Point{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	var neighbors []Point

	for _, dir := range directions {
		newX := p.X + dir.X
		newY := p.Y + dir.Y

		if isInBounds(newX, newY, rows, cols) && grid[newX][newY] == 0 && !visited[newX][newY] {
			neighbors = append(neighbors, Point{newX, newY})
		}
	}

	return neighbors
}

// Binary matrix so only 0 and 1s
// Path must consist of 0s, start is topmost left corner, end is bottomost right corner
// Each cell has 8 neigbours
// This is a matrix traversal problem with certain conditions
// The conditions are that we cannot traverse a path containing 1s
// We will need to explore ALL paths systemically, since we aren't JUST trying to reach the end
// So we can use BFS instead of DFS to systemically move towards the end, the very first time we find the endpoint
// We also don't want to revisit a node, right? Because if we revist the node it means there was a shorter path to it
// So any path that includes an already visisted node will not be optimal
// we can exit out since this is by definition the shortest path, the first path via BFS to find is must be shortest
