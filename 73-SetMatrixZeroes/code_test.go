package leetcode


import (
	"testing"
	"github.com/stretchr/testify/assert"
)


var cases = []struct{
	n			[][]int
	expected	[][]int
}{
	{
		n: [][]int{
			{1, 2, 3, 4},
			{5, 0, 7, 0},
			{9, 1, 1, 1},
		},
		expected: [][]int{
			{1, 0, 3, 0},
			{0, 0, 0, 0},
			{9, 0, 1, 0},
		},
	},
}


func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		setZeroes(c.n)

		ast.Equal(c.n, c.expected)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		for _, c := range cases {
			setZeroes(c.n)
		}
	}
}