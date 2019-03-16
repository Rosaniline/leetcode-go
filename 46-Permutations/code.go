package leetcode

func permute(nums []int) [][]int {
	res := make([][]int, 0)

	var permuteHelper func([]int, int)

	copySlice := func(src []int) []int {
		temp := make([]int, len(src))
		copy(temp, src)

		return temp
	}

	permuteHelper = func(nums []int, start int) {
		if start == len(nums)-1 {
			res = append(res, copySlice(nums))
			return
		}

		for i := start; i < len(nums); i++ {
			nums[start], nums[i] = nums[i], nums[start]
			permuteHelper(nums, start+1)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}

	permuteHelper(nums, 0)

	return res
}
