package leetcode

// Invert a binary tree.

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	res := append(inorderTraversal(root.Left), root.Val)
	res = append(res, inorderTraversal(root.Right)...)

	return res
}
