package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var cases = []struct {
	num1     string
	expected int
}{
	{
		num1:     "abc",
		expected: 3,
	},
	{
		num1:     "aaa",
		expected: 6,
	},
	{
		num1:     "",
		expected: 0,
	},
	{
		num1:     "longtimenosee",
		expected: 14,
	},
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(c.expected, countSubstrings(c.num1))
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			countSubstrings(c.num1)
		}
	}
}
