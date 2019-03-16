package leetcode

func numJewelsInStones(J string, S string) int {
	temp := make(map[rune]bool)

	for _, j := range J {
		temp[j] = true
	}

	count := 0

	for _, s := range S {
		if _, ok := temp[s]; ok {
			count += 1
		}
	}

	return count
}
