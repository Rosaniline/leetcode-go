package leetcode


import (
	"testing"
	"github.com/stretchr/testify/assert"
)


var cases = []struct{
	num1		int
	expected	bool
}{
	{
		num1:		10,
		expected: 	false,
	},
	{
		num1:		16,
		expected: 	true,
	},
	{
		num1:		24,
		expected: 	false,
	},
	{
		num1:		46,
		expected:	true,
	},
	{
		num1:		821,
		expected:	true,
	},
}


func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(reorderedPowerOf2(c.num1), c.expected)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		for _, c := range cases {
			reorderedPowerOf2(c.num1)
		}
	}
}