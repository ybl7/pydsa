package main

import (
	dq "github.com/idsulik/go-collections/deque"
)

func RottenOranges(grid [][]int) int {
	maxR := len(grid) - 1
	maxL := len(grid[0]) - 1
	fresh := 0
	mins := 0

	q := dq.New[Coor](maxR * maxL)
	for r, row := range grid {
		for c, _ := range row {
			// Count fresh oranges
			if grid[r][c] == 1 {
				fresh++
				// Track position of rotten oranges
			} else if grid[r][c] == 2 {
				q.PushBack(Coor{
					Row: r,
					Col: c,
				})
			}
		}
	}

	for !q.IsEmpty() && fresh > 0 {
		l := q.Len()
		mins++

		for i := 0; i < l; i++ {
			c, _ := q.PopFront()

			rArr := []Coor{}
			rArr, fresh = RotNeighbours(c, grid, maxR, maxL, fresh)
			for _, r := range rArr {
				q.PushBack(r)
			}
		}
	}

	if fresh == 0 {
		// We were able to rot all oranges
		return mins
	}

	return -1
}

func RotNeighbours(c Coor, grid [][]int, maxR, maxL, fresh int) ([]Coor, int) {
	diff := []Coor{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	out := []Coor{}

	for _, d := range diff {
		dr := c.Row + d.Row
		dc := c.Col + d.Col
		if dr < 0 || dr > maxR || dc < 0 || dc > maxL {
			// Out of bounds, do nothing
		} else {
			if grid[dr][dc] == 1 {
				grid[dr][dc] = 2 // rot the nieghbour so we don't ever count it again as a fresh orange
				fresh--
				out = append(out, Coor{
					Row: dr,
					Col: dc,
				})
			}
		}
	}

	return out, fresh
}

type Coor struct {
	Row int
	Col int
}
