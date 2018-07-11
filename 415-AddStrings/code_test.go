package leetcode


import (
	"testing"
	"github.com/stretchr/testify/assert"
)


var cases = []struct{
	num1		string
	num2		string
	expected	string
}{
	{
		num1:		"199",
		num2:		"9487",
		expected: 	"9686",
	},
	{
		num1:		"0",
		num2:		"9487",
		expected: 	"9487",
	},
	{
		num1:		"9487",
		num2:		"0",
		expected: 	"9487",
	},
	{
		num1:		"999",
		num2:		"1",
		expected: 	"1000",
	},
}


func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(addStrings(c.num1, c.num2), c.expected)
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		for _, c := range cases {
			addStrings(c.num1, c.num2)
		}
	}
}