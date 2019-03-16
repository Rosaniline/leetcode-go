package leetcode

func lengthOfLastWord(s string) int {
	length := 0

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			length++
		}

		if s[i] == ' ' && length > 0 {
			break
		}
	}

	return length
}
