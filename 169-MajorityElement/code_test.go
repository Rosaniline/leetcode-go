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
		num1:     []int{3, 2, 3},
		expected: 3,
	},
	{
		num1:     []int{2, 2, 1, 1, 1, 2, 2},
		expected: 2,
	},
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(c.expected, majorityElement(c.num1))
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			majorityElement(c.num1)
		}
	}
}
