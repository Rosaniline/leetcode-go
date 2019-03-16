package leetcode

import (
	"fmt"
	"testing"
)

var cases = []struct {
	num1     *TreeNode
	expected [][]int
}{
	{
		num1:     nil,
		expected: [][]int{},
	},
	{
		num1:     &TreeNode{Val: 1},
		expected: [][]int{{1}},
	},
	{
		num1: &TreeNode{Val: 1, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 2}},
		expected: [][]int{
			{1},
			{0, 2},
		},
	},
	{
		num1: &TreeNode{Val: 3, Left: &TreeNode{Val: 9, Left: &TreeNode{Val: 10}}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
		expected: [][]int{
			{3},
			{9, 20},
			{10, 15, 7},
		},
	},
}

func TestOK(t *testing.T) {
	// ast := assert.New(t)

	for _, c := range cases {
		// ast.Equal(c.expected, levelOrder(c.num1))
		levelOrder(c.num1)
	}

	for i := 0; i < 20; i++ {
		fmt.Println(i, stesp(i))
	}

	fmt.Println(610, stesp(610))
}

// func Benchmark(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		for _, c := range cases {
// 			levelOrder(c.num1)
// 		}
// 	}
// }

func Benchmark(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stesp(100)
	}
}
