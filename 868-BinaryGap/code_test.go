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
		num1:		22,
		expected: 	2,
	},
	{
		num1:		0,
		expected: 	0,
	},
	{
		num1:		1,
		expected: 	0,
	},
}


func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(binaryGap(c.num1), c.expected)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		for _, c := range cases {
			binaryGap(c.num1)
		}
	}
}