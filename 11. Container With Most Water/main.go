package main

import "fmt"

func main() {
	fmt.Println(maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
}

func maxArea(height []int) int {
	max := 0
	l := 0
	r := len(height) - 1

	for l < r {
		min := min(height[l], height[r])
		area := min * (r - l)
		if area > max {
			max = area
		}

		if height[l] < height[r] {
			l++
			continue
		}

		r--
	}

	return max
}
