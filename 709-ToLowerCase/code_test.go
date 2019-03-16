package leetcode

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func from_string(s string) *big.Int {
	num, _ := new(big.Int).SetString(s, 10)

	return num
}

var cases = []struct {
	n        string
	expected string
}{
	{
		n:        "Hello",
		expected: "hello",
	},
	{
		n:        "here",
		expected: "here",
	},
	{
		n:        "LOVELY",
		expected: "lovely",
	},
	{
		n:        "",
		expected: "",
	},
}

func TestOK(t *testing.T) {
	for _, c := range cases {
		t.Run("test", func(t *testing.T) {
			ast := assert.New(t)

			ast.Equal(c.expected, toLowerCase(c.n))
		})
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			toLowerCase(c.n)
		}
	}
}
