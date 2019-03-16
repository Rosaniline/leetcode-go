package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	n        []int
	expected [][]int
}{
	{
		n: []int{1, 2, 3, 4, 5, 6, 7},
		expected: [][]int{
			{1, 2, 3},
		},
	},
}

func TestOK(t *testing.T) {
	for _, c := range cases {
		t.Run("test", func(t *testing.T) {
			ast := assert.New(t)

			ast.Equal(c.expected, permute(c.n))
		})
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			permute(c.n)
		}
	}
}
