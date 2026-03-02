package main

import "fmt"

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
}

func subsets(nums []int) [][]int {
	var result [][]int
	var path []int

	var backtrack func(start int)
	backtrack = func(start int) {
		tmp := append([]int{}, path...)
		result = append(result, tmp)

		for i := start; i < len(nums); i++ {
			path = append(path, nums[i])
			backtrack(i + 1)
			path = path[:len(path)-1]
		}
	}

	backtrack(0)
	return result
}
