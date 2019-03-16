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
		num1:     []int{},
		expected: []int{},
	},
	{
		num1:     []int{1, 2, 3, 4},
		expected: []int{24, 12, 8, 6},
	},
	{
		num1:     []int{0},
		expected: []int{0},
	},
	{
		num1:     []int{1, 2},
		expected: []int{2, 1},
	},
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(c.expected, productExceptSelf(c.num1))
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			productExceptSelf(c.num1)
		}
	}
}
