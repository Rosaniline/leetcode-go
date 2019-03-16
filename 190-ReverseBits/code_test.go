package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var cases = []struct {
	str      uint32
	expected uint32
}{
	{
		str:      43261596,
		expected: 964176192,
	},
	{
		str:      4294967293,
		expected: 3221225471,
	},
	{
		str:      0,
		expected: 0,
	},
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(reverseBits(c.str), c.expected)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			reverseBits(c.str)
		}
	}
}
