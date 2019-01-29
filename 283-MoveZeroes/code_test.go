package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	num1     []int
	expected []int
}{
	{
		num1:     []int{0, 1, 0, 3, 12},
		expected: []int{1, 3, 12, 0, 0},
	},
	{
		num1:     []int{0, 0, 0, 0},
		expected: []int{0, 0, 0, 0},
	},
	{
		num1:     []int{0, 1, 0, 2},
		expected: []int{1, 2, 0, 0},
	},
	{
		num1:     []int{0, 0, 0, 1},
		expected: []int{1, 0, 0, 0},
	},
	{
		num1:     []int{1, 0, 0, 0, 3},
		expected: []int{1, 3, 0, 0, 0},
	},
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		moveZeroes(c.num1)
		ast.Equal(c.expected, c.num1)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			moveZeroes(c.num1)
		}
	}
}
