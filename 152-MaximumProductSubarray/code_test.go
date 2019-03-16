package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	num1     []int
	expected int
}{
	{
		num1:     []int{},
		expected: 0,
	},
	{
		num1:     []int{1, 2, 3, 4},
		expected: 24,
	},
	{
		num1:     []int{2, 3, -2, 4},
		expected: 6,
	},
	{
		num1:     []int{-2, 0, -1},
		expected: 0,
	},
	{
		num1:     []int{-1, -2, -3, 4, -5, -6},
		expected: 720,
	},
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(c.expected, maxProduct(c.num1))
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			maxProduct(c.num1)
		}
	}
}
