package main

import "fmt"

func main() {
	fmt.Println(findDuplicates([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}

func findDuplicates(nums []int) []int {
	m := make(map[int]bool)
	out := []int{}

	for _, n := range nums {
		if _, ok := m[n]; ok {
			out = append(out, n)
		} else {
			m[n] = true
		}
	}

	return out
}
