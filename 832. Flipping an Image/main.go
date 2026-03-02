package main

import "fmt"

func main() {
	fmt.Println(flipAndInvertImage([][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}))
}

func flipAndInvertImage(image [][]int) [][]int {
	for _, row := range image {
		flip(row)
		invert(row)
	}

	return image
}

func flip(row []int) {
	r := 0
	l := len(row) - 1

	for r < l {
		tmp := row[r]
		row[r] = row[l]
		row[l] = tmp

		r++
		l--
	}
}

func invert(row []int) {
	for i, n := range row {
		if n == 0 {
			row[i] = 1
		} else {
			row[i] = 0
		}
	}
}
