package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverseWords("the sky is blue"))
}

func reverseWords(s string) string {
	fields := strings.Fields(s)

	builder := strings.Builder{}

	for i := len(fields) - 1; i >= 0; i-- {
		builder.WriteString(fields[i])
		if i != 0 {
			builder.WriteRune(' ')
		}
	}

	return builder.String()
}
