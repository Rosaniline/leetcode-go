package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math/big"
	"testing"
)

func from_string(s string) *big.Int {
	num, _ := new(big.Int).SetString(s, 10)

	return num
}

var cases = []struct {
	n        int
	expected *big.Int
}{
	{
		n:        2,
		expected: big.NewInt(1),
	},
	{
		n:        3,
		expected: big.NewInt(2),
	},
	{
		n:        4,
		expected: big.NewInt(3),
	},
	{
		n:        100,
		expected: from_string("354224848179261915075"),
	},
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(c.expected, fib(c.n))
	}

	fmt.Println(fib(8181))
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			fib(c.n)
		}
	}
}
