package leetcode


func longestPalindrome(s string) int {
	dict := make(map[rune]int)

	for _, r := range []rune(s) {
		dict[r] ++
	}

	var (
		odd = false
		res = 0
	)

	for _, v := range dict {
		res += (v/2)*2

		if v % 2 == 1 {
			odd = true
		}
	}

	if odd {
		res ++
	}

	return res
}