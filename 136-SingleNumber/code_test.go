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
		num1:     []int{1, 1, 2},
		expected: 2,
	},
	{
		num1:     []int{4, 1, 2, 1, 2},
		expected: 4,
	},
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(c.expected, singleNumber(c.num1))
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			singleNumber(c.num1)
		}
	}
}
