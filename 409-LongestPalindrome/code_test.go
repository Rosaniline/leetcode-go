package leetcode


import (
	"testing"
	"github.com/stretchr/testify/assert"
)


var cases = []struct{
	str			string
	expected	int
}{
	{
		str:		"abccccdd",
		expected: 	7,
	},
	{
		str:		"aAbBcCdD",
		expected:	1,
	},
	{
		str:		"",
		expected:	0,
	},
}


func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(longestPalindrome(c.str), c.expected)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		for _, c := range cases {
			longestPalindrome(c.str)
		}
	}
}