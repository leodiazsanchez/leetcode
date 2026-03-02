package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Printf("%d\n", myAtoi("9223372036854775808"))
}

func myAtoi(s string) int {
	trim := strings.TrimSpace(s)
	if trim == "" {
		return 0
	}

	negative := trim[0] == '-'
	if negative || trim[0] == '+' {
		trim = trim[1:]
	}

	var res int64 = 0
	for _, r := range trim {
		if r >= '0' && r <= '9' {
			res = res*10 + int64(r-'0')

			if !negative && res > math.MaxInt32 {
				return math.MaxInt32
			}
			if negative && -res < math.MinInt32 {
				return math.MinInt32
			}

		} else {
			break
		}
	}

	if negative {
		res = -res
	}

	if res > math.MaxInt32 {
		return math.MaxInt32
	}
	if res < math.MinInt32 {
		return math.MinInt32
	}

	return int(res)
}
