package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var cases = []struct {
	num1     string
	num2     string
	expected int
}{
	{
		num1:     "aA",
		num2:     "aAAbbbb",
		expected: 3,
	},
	{
		num1:     "z",
		num2:     "ZZ",
		expected: 0,
	},
	{
		num1:     "123awefzsvs",
		num2:     "a",
		expected: 1,
	},
	{
		num1:     "",
		num2:     "123",
		expected: 0,
	},
	{
		num1:     "123",
		num2:     "",
		expected: 0,
	},
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(c.expected, numJewelsInStones(c.num1, c.num2))
	}
}
