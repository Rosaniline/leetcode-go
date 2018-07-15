package leetcode


import (
	"testing"
	"github.com/stretchr/testify/assert"
)


var cases = []struct{
	num1		[]int
	expected	int
}{
	{
		num1:		[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		expected: 	6,
	},
	{
		num1:		[]int{},
		expected: 	0,
	},
	{
		num1:		[]int{-1, -2, -3, -4, -5, -6},
		expected: 	-1,
	},
	{
		num1:		[]int{1, 1, 1, 1, -999, 1, 1, 1, 1},
		expected: 	4,
	},
}


func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(c.expected, maxSubArray(c.num1))
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		for _, c := range cases {
			maxSubArray(c.num1)
		}
	}
}