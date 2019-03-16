package leetcode

import (
	"fmt"
	"math/big"
	"strings"
)

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

func levelOrder(root *TreeNode) {
	levels := make([][]string, 0)

	if root == nil {
		return
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var (
			size  = len(queue)
			local = make([]string, 0)
		)

		for i := 0; i < size; i++ {
			node := queue[i]
			local = append(local, fmt.Sprintf("(%v)", node.Val))

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		levels = append(levels, local)
		queue = queue[size:]
	}

	for _, level := range levels {
		println(strings.Join(level, ", "))
	}

	return
}

func stesp(n int) *big.Int {
	dp := []*big.Int{
		big.NewInt(1),
		big.NewInt(1),
		big.NewInt(2),
		big.NewInt(4),
		big.NewInt(8),
		big.NewInt(16),
		big.NewInt(32),
	}

	if n <= 6 {
		return dp[n]
	}

	for i := 7; i <= n; i++ {
		temp := new(big.Int)
		temp = temp.Lsh(dp[6], 1)
		temp = temp.Sub(temp, dp[0])

		dp = append(dp[1:], temp)
	}

	return dp[len(dp)-1]
}

func testt(n int) *big.Int {
	// A = 3**(n+6)
	// M = A**6 - A**5 - A**4 - A**3 - A**2 - A - 1
	// return pow(A, n+6, M) % A
	return nil
}
