package leetcode


import (
	"testing"
	"github.com/stretchr/testify/assert"
)


var cases = []struct{
	num1		int
	expected	[]int
}{
	{
		num1:		5,
		expected: 	[]int{0, 1, 1, 2, 1, 2},
	},
	{
		num1:		0,
		expected:	[]int{0},
	},
}


func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(c.expected, countBits(c.num1))
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		//for _ := range cases {
		countBits(1000)
		//}
	}
}