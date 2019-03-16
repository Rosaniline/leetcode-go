package leetcode

func productExceptSelf(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}

	res := make([]int, len(nums))
	seed := nums[0]

	for i := 1; i < len(nums); i++ {
		res[i] = seed
		seed *= nums[i]
	}

	seed = nums[len(nums)-1]

	for i := len(nums) - 2; i > 0; i-- {
		res[i] *= seed
		seed *= nums[i]
	}

	res[0] = seed

	return res
}
