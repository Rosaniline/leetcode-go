package leetcode


import (
	"testing"
	"github.com/stretchr/testify/assert"
)


var cases = []struct{
	num1		[]int
	expected	bool
}{
	{
		num1:		[]int{1, 5, 11, 5},
		expected: 	true,
	},
	{
		num1:		[]int{1, 2, 3, 5},
		expected:	false,
	},
	{
		num1:		[]int{1, 2, 5},
		expected:	false,
	},
}


func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(c.expected, canPartition(c.num1))
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		for _, c := range cases {
			canPartition(c.num1)
		}
	}
}