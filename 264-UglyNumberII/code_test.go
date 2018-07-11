package _264


import (
	"testing"
	"github.com/stretchr/testify/assert"
)


var cases = []struct{
	n			int
	expected	int
}{
	{
		n: 			1,
		expected: 	1,
	},
	{
		n: 			2,
		expected: 	2,
	},
	{
		n: 			100,
		expected: 	1536,
	},
}


func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(nthUglyNumber(c.n), c.expected)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		for _, c := range cases {
			nthUglyNumber(c.n)
		}
	}
}