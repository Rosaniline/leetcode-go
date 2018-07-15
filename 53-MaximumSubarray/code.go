package leetcode


// Given an integer array nums, find the contiguous subarray
// (containing at least one number) which has the largest sum and return its sum.


func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var (
		globalSum 	= nums[0]
		localSum	= nums[0]
	)

	for i := 1; i < len(nums); i ++ {
		if localSum < 0 {
			localSum = nums[i]
		} else {
			localSum += nums[i]
		}

		if localSum > globalSum {
			globalSum = localSum
		}
	}

	return globalSum
}