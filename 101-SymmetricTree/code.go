package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func iteratice(root *TreeNode) bool {
	if root == nil {
		return true
	}

	var (
		left  = new(TreeNode)
		right = new(TreeNode)
		queue = []*TreeNode{root.Left, root.Right}
	)

	for len(queue) > 0 {
		left, right, queue = queue[0], queue[1], queue[2:]

		if left == nil && right == nil {
			continue
		}

		if left != nil && right != nil {
			if left.Val != right.Val {
				return false
			}

			queue = append(queue, left.Left, right.Right, left.Right, right.Left)

			continue
		}

		return false
	}

	return true
}

func recursive(root *TreeNode) bool {
	var isMirror func(left *TreeNode, right *TreeNode) bool

	isMirror = func(left *TreeNode, right *TreeNode) bool {
		if left == nil || right == nil {
			return left == right
		}

		return left.Val == right.Val && isMirror(left.Left, right.Right) && isMirror(left.Right, right.Left)
	}

	if root == nil {
		return true
	}

	return isMirror(root.Left, root.Right)
}

func isSymmetric(root *TreeNode) bool {
	return recursive(root)
}
