package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var cases = []struct {
	num1     *TreeNode
	expected []int
}{
	{
		num1:     &TreeNode{Val: 1, Right: &TreeNode{Val: 3, Left: &TreeNode{Val: 2}}},
		expected: []int{1, 2, 3},
	},
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(c.expected, inorderTraversal(c.num1))
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			inorderTraversal(c.num1)
		}
	}
}
