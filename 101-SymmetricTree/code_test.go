package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	num1     *TreeNode
	expected bool
}{
	{
		num1:     nil,
		expected: true,
	},
	{
		num1:     &TreeNode{Val: 1},
		expected: true,
	},
	{
		num1:     &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2}},
		expected: false,
	},
	{
		num1:     &TreeNode{Val: 1, Left: &TreeNode{Val: 0}},
		expected: false,
	},
	{
		num1:     &TreeNode{Val: 1, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}}, Right: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}}},
		expected: false,
	},
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(c.expected, isSymmetric(c.num1))
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			isSymmetric(c.num1)
		}
	}
}
