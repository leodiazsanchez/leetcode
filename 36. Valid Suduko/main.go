package main

import "fmt"

const length = 9

func main() {
	board := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board))
}

func isValidSudoku(board [][]byte) bool {
	for _, row := range board {
		if !valid(row) {
			return false
		}
	}

	for i := 0; i < length; i++ {
		col := make([]byte, 0, length)
		for j := 0; j < length; j++ {
			col = append(col, board[j][i])
		}
		if !valid(col) {
			return false
		}
	}

	for rowStart := 0; rowStart < length; rowStart += 3 {
		for colStart := 0; colStart < length; colStart += 3 {
			grid := make([]byte, 0, length)
			for i := 0; i < 3; i++ {
				for j := 0; j < 3; j++ {
					grid = append(grid, board[rowStart+i][colStart+j])
				}
			}
			if !valid(grid) {
				return false
			}
		}
	}

	return true
}

func valid(s []byte) bool {
	seen := make(map[byte]bool)
	for _, b := range s {
		if b == '.' {
			continue
		}
		if seen[b] {
			return false
		}
		seen[b] = true
	}
	return true
}
