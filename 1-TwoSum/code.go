package leetcode

func twoSum(nums []int, target int) []int {
	diffIndex := make(map[int]int)

	for i, n := range nums {
		index, ok := diffIndex[n]

		if ok {
			return []int{index, i}
		} else {
			diffIndex[target-n] = i
		}
	}

	return nil
}
