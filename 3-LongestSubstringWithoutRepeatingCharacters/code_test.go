package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	num1     string
	expected int
}{
	{
		num1:     "",
		expected: 0,
	},
	{
		num1:     "abc",
		expected: 3,
	},
	{
		num1:     "abcabcbb",
		expected: 3,
	},
	{
		num1:     "bbbbb",
		expected: 1,
	},
	{
		num1:     "pwwkew",
		expected: 3,
	},
	{
		num1:     "aab",
		expected: 2,
	},
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(c.expected, lengthOfLongestSubstring(c.num1))
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			lengthOfLongestSubstring(c.num1)
		}
	}
}
