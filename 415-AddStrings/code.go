package leetcode


func max(a int, b int) int {
	if a > b {
		return a
	}

	return b
}


func addStrings(num1 string, num2 string) string {
	var (
		maxlen 	= max(len(num1), len(num2))
		res 	= ""
		carry 	= 0
		len1	= len(num1)
		len2	= len(num2)
	)

	for i := 0; i < maxlen; i ++ {
		if len1 - 1 - i >= 0 {
			carry += int(num1[len1 - 1 - i] - '0')
		}

		if len2 - 1 - i >= 0 {
			carry += int(num2[len2 - 1 - i] - '0')
		}

		res = string(carry%10 + '0') + res
		carry = carry/10
	}

	if carry == 1 {
		res = "1" + res
	}

	return res
}