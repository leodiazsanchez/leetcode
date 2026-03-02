package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
}

func containsNearbyDuplicate(nums []int, k int) bool {
	seen := make(map[int]int)

	for i, n := range nums {
		if _, ok := seen[n]; !ok {
			seen[n] = i
			continue
		}

		if math.Abs(float64(i-seen[n])) <= float64(k) {
			return true
		}

		seen[n] = i
	}

	return false
}
