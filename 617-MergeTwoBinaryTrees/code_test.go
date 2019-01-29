package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var cases = []struct {
	num1     *TreeNode
	num2     *TreeNode
	expected *TreeNode
}{
	{
		num1:     nil,
		num2:     nil,
		expected: nil,
	},
	{
		num1:     &TreeNode{Val: 8},
		num2:     nil,
		expected: &TreeNode{Val: 8},
	},
	{
		num1:     nil,
		num2:     &TreeNode{Val: 8},
		expected: &TreeNode{Val: 8},
	},
	{
		num1:     &TreeNode{Val: 10},
		num2:     &TreeNode{Val: 8},
		expected: &TreeNode{Val: 18},
	},
	{
		num1:     &TreeNode{Val: 10, Left: &TreeNode{Val: 7}},
		num2:     &TreeNode{Val: 8},
		expected: &TreeNode{Val: 18, Left: &TreeNode{Val: 7}},
	},
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(c.expected, mergeTrees(c.num1, c.num2))
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			mergeTrees(c.num1, c.num2)
		}
	}
}
