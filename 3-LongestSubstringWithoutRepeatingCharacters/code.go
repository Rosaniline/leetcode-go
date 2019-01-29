package leetcode

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func lengthOfLongestSubstring(s string) int {
	length := -1
	dict := make(map[rune]int)

	for i, r := range s {
		preIndex, ok := dict[r]

		if ok {
			length = max(length, i-preIndex)
		}

		dict[r] = i
	}

	if length == -1 {
		return len(s)
	}

	return length
}
