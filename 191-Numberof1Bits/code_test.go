package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var cases = []struct {
	str      uint32
	expected int
}{
	{
		str:      11,
		expected: 3,
	},
	{
		str:      256,
		expected: 1,
	},
	{
		str:      4294967293,
		expected: 31,
	},
	{
		str:      0,
		expected: 0,
	},
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(hammingWeight(c.str), c.expected)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			hammingWeight(c.str)
		}
	}
}
