package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Printf("%v\n", merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}))
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return nil
	}

	slices.SortFunc(intervals, func(a, b []int) int {
		return a[0] - b[0]
	})

	out := make([][]int, 0)
	curr := intervals[0]

	for i := 1; i < len(intervals); i++ {
		next := intervals[i]

		if next[0] <= curr[1] {
			curr[1] = max(curr[1], next[1])
		} else {
			out = append(out, curr)
			curr = next
		}
	}

	out = append(out, curr)

	return out
}
