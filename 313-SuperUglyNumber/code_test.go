package leetcode

import (
	"testing"
	"github.com/stretchr/testify/assert"
)


var cases = []struct{
	n			int
	primes		[]int
	expected	int
}{
	{
		n: 			12,
		primes: 	[]int{2, 7, 13, 19},
		expected: 	32,
	},
	{
		n: 			100,
		primes: 	[]int{2, 3, 5},
		expected: 	1536,
	},
}


func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(nthSuperUglyNumber(c.n, c.primes), c.expected)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		for _, c := range cases {
			nthSuperUglyNumber(c.n, c.primes)
		}
	}
}