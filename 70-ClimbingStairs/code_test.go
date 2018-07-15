package leetcode


import (
	"testing"
	"github.com/stretchr/testify/assert"
)


var cases = []struct{
	num1		int
	expected	int
}{
	{
		num1:		2,
		expected: 	2,
	},
	{
		num1:		3,
		expected: 	3,
	},
	{
		num1:		0,
		expected:	0,
	},
}


func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(c.expected, climbStairs(c.num1))
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		for _, c := range cases {
			climbStairs(c.num1)
		}
	}
}