package main

import "fmt"

type point struct {
	x, y int
}

func main() {
	grid := [][]int{
		{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0},
		{0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0},
	}

	fmt.Println(maxAreaOfIsland(grid))
}

func maxAreaOfIsland(grid [][]int) int {
	maxArea := 0
	for i, row := range grid {
		for j := range row {
			if grid[i][j] == 1 {
				if a := bfs(point{i, j}, grid); a > maxArea {
					maxArea = a
				}
			}
		}
	}

	return maxArea
}

func bfs(p point, grid [][]int) int {
	area := 0

	dirs := [][2]int{
		{1, 0}, {-1, 0}, {0, 1}, {0, -1},
	}

	q := []point{p}
	grid[p.x][p.y] = 0

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		area++

		for _, d := range dirs {
			nr := curr.x + d[0]
			nc := curr.y + d[1]

			if nr >= 0 && nr < len(grid) && nc >= 0 && nc < len(grid[0]) {
				if grid[nr][nc] == 1 {
					grid[nr][nc] = 0
					q = append(q, point{nr, nc})
				}
			}
		}
	}

	return area
}
