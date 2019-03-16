package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	var (
		res  = nums[0]
		vmax = nums[0]
		vmin = nums[0]
	)

	for _, n := range nums[1:] {
		if n < 0 {
			vmin, vmax = vmax, vmin
		}

		vmax = max(n, vmax*n)
		vmin = min(n, vmin*n)

		res = max(res, vmax)
	}

	return res
}
