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
		num1:     []int{1, 2, 3, 1},
		expected: 4,
	},
	{
		num1:     []int{2, 7, 9, 3, 1},
		expected: 12,
	},
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(c.expected, rob(c.num1))
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			rob(c.num1)
		}
	}
}
