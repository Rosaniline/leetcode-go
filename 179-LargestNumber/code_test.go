package leetcode


import (
	"testing"
	"github.com/stretchr/testify/assert"
)


var cases = []struct{
	n			[]int
	expected	string
}{
	{
		n: 			[]int{10, 2},
		expected: 	"210",
	},
	{
		n: 			[]int{3, 30, 34, 5, 9},
		expected: 	"9534330",
	},
	{
		n:			[]int{0,0,0,0,0,0},
		expected: 	"0",
	},
	{
		n:			[]int{0,0,0,0,0,0,1},
		expected: 	"1000000",
	},
}


func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(largestNumber(c.n), c.expected)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		for _, c := range cases {
			largestNumber(c.n)
		}
	}
}