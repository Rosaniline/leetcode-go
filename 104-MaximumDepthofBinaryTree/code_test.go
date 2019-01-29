package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	num1     *TreeNode
	expected int
}{
	{
		num1:     nil,
		expected: 0,
	},
	{
		num1:     &TreeNode{Val: 1},
		expected: 1,
	},
	{
		num1:     &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2}},
		expected: 2,
	},
}

func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(c.expected, maxDepth(c.num1))
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, c := range cases {
			maxDepth(c.num1)
		}
	}
}
