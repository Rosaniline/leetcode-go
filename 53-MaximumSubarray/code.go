package leetcode

// Given an integer array nums, find the contiguous subarray
// (containing at least one number) which has the largest sum and return its sum.

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var (
		globalSum = nums[0]
		localSum  = nums[0]
	)

	for _, n := range nums[1:] {
		localSum = max(localSum+n, n)
		globalSum = max(globalSum, localSum)
	}

	return globalSum
}
