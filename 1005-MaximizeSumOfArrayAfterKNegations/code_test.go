package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	num1     []int
	num2     int
	expected int
}{
	{
		num1:     []int{4, 2, 3},
		num2:     1,
		expected: 5,
	},
	{
		num1:     []int{3, -1, 0, 2},
		num2:     3,
		expected: 6,
	},
	{
		num1:     []int{2, -3, -1, 5, -4},
		num2:     2,
		expected: 13,
	},
	{
		num1:     []int{-2, 9, 9, 8, 4},
		num2:     5,
		expected: 32,
	},
	{
		num1:     []int{-1, -2, -3},
		num2:     4,
		expected: 4,
	},
	{
		num1:     []int{-5, -6, -2, -8, 6, 9},
		num2:     4,
		expected: 36,
	},
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(c.expected, largestSumAfterKNegations(c.num1, c.num2))
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			largestSumAfterKNegations(c.num1, c.num2)
		}
	}
}
