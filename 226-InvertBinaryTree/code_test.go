package leetcode


import (
	"testing"
	"github.com/stretchr/testify/assert"
)


var cases = []struct{
	num1		*TreeNode
	expected	*TreeNode
}{
	{
		num1:		&TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2}},
		expected: 	&TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 0}},
	},
	{
		num1:		nil,
		expected:	nil,
	},
}


func TestOK(t *testing.T) {
	ast := assert.New(t)

	for _, c := range cases {
		ast.Equal(c.expected, invertTree(c.num1))
	}
}

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i ++ {
		for _, c := range cases {
			invertTree(c.num1)
		}
	}
}