package leetcode

import (
	"sort"
)

// Given a non-empty array containing only positive integers,
// find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.
//
// Note:
// Each of the array element will not exceed 100.
// The array size will not exceed 200.


func sum(nums []int) (s int) {
	for _, n := range nums {
		s += n
	}
	return
}


func canPartition(nums []int) bool {
	s := sum(nums)

	if s & 1 == 1 {
		return false
	}

	sort.Ints(nums)
	s = s/2
	dp := make([]bool, s + 1)
	dp[0] = true

	for _, n := range nums {
		for i := len(dp) - 1; i >= n; i -- {
			dp[i] = dp[i] || dp[i - n]
		}
	}

	return dp[s]
}