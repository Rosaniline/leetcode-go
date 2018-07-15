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
		num1:		[]int{7, 1, 5, 3, 6, 4},
		expected: 	5,
	},
	{
		num1:		[]int{7, 6, 4, 3, 1},
		expected: 	0,
	},
	{
		num1:		[]int{},
		expected:	0,
	},
}


func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(c.expected, maxProfit(c.num1))
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		for _, c := range cases {
			maxProfit(c.num1)
		}
	}
}