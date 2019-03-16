package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	n        string
	expected int
}{
	{
		n:        "Hello World",
		expected: 5,
	},
	{
		n:        "here",
		expected: 4,
	},
	{
		n:        "",
		expected: 0,
	},
	{
		n:        "a ",
		expected: 1,
	},
}

func TestOK(t *testing.T) {
	for _, c := range cases {
		t.Run("test", func(t *testing.T) {
			ast := assert.New(t)

			ast.Equal(c.expected, lengthOfLastWord(c.n))
		})
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			lengthOfLastWord(c.n)
		}
	}
}
