package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	num1     *ListNode
	num2     *ListNode
	expected *ListNode
}{
	{
		num1:     nil,
		num2:     nil,
		expected: nil,
	},
	{
		num1:     nil,
		num2:     &ListNode{Val: 5},
		expected: &ListNode{Val: 5},
	},
	{
		num1:     &ListNode{Val: 5},
		num2:     nil,
		expected: &ListNode{Val: 5},
	},
	{
		num1:     &ListNode{Val: 0, Next: nil},
		num2:     &ListNode{Val: 0, Next: nil},
		expected: &ListNode{Val: 0, Next: nil},
	},
	{
		num1:     &ListNode{Val: 7, Next: nil},
		num2:     &ListNode{Val: 3, Next: nil},
		expected: &ListNode{Val: 0, Next: &ListNode{Val: 1}},
	},
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(c.expected, addTwoNumbers(c.num1, c.num2))
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			addTwoNumbers(c.num1, c.num2)
		}
	}
}
