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
		num1:		[]int{10, 15, 20},
		expected: 	15,
	},
	{
		num1:		[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1},
		expected: 	6,
	},
	{
		num1:		[]int{},
		expected:	0,
	},
}


func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(c.expected, minCostClimbingStairs(c.num1))
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		for _, c := range cases {
			minCostClimbingStairs(c.num1)
		}
	}
}