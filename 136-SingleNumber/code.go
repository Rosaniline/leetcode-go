package leetcode

func singleNumber(nums []int) int {
	seed := 0

	for _, n := range nums {
		seed = seed ^ n
	}

	return seed
}
