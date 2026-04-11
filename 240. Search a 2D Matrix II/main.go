package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 4, 7, 11, 15},
		{2, 5, 8, 12, 19},
		{3, 6, 9, 16, 22},
		{10, 13, 14, 17, 24},
		{18, 21, 23, 26, 30},
	}
	fmt.Printf("%t\n", searchMatrix(matrix, 5))
}

func searchMatrix(matrix [][]int, target int) bool {
	row := 0
	col := len(matrix[0]) - 1

	for row < len(matrix) && col >= 0 {
		val := matrix[row][col]

		if val == target {
			return true
		} else if val > target {
			col--
		} else {
			row++
		}
	}

	return false
}
