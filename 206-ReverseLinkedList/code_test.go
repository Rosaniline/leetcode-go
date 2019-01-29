package leetcode


import (
	"testing"
	"github.com/stretchr/testify/assert"
)


var cases = []struct{
	num1		*ListNode
	expected	*ListNode
}{
	{
		num1:		&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}},
		expected: 	&ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}},
	},
	{
		num1:		nil,
		expected:	nil,
	},
}


func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(c.expected, reverseList(c.num1))
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		for _, c := range cases {
			reverseList(c.num1)
		}
	}
}