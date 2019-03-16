package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	res := []int{nums[0], max(nums[0], nums[1])}

	for i := 2; i < len(nums); i++ {
		res = append(res, max(res[i-2]+nums[i], res[i-1]))
	}

	return res[len(res)-1]
}
