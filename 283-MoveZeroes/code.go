package leetcode

func moveZeroes(nums []int) {
	index := 0

	for _, n := range nums {
		if n != 0 {
			nums[index] = n

			index++
		}
	}

	for i := index; i < len(nums); i++ {
		nums[i] = 0
	}
}
