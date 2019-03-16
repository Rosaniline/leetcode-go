package leetcode

import (
	"strings"
)

func toLowerCase(str string) string {
	return strings.Map(func(r rune) rune {
		if 'A' <= r && r <= 'Z' {
			return r + 'a' - 'A'
		}

		return r
	}, str)
}
