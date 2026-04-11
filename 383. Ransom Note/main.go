package main

import "fmt"

func main() {
	fmt.Printf("%t\n", canConstruct("a", "b"))
}

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) > len(magazine) {
		return false
	}

	m := make(map[rune]int)

	for _, r := range magazine {
		m[r]++
	}

	for _, r := range ransomNote {
		if m[r] == 0 {
			return false
		}
		m[r]--
	}

	return true
}
