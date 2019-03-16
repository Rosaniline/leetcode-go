package leetcode

func initTwoDimensionalArr(size int) [][]bool {
	temp := make([][]bool, size)

	for i := range temp {
		temp[i] = make([]bool, size)
	}

	return temp
}

func countSubstrings(s string) int {
	var (
		size  = len(s)
		temp  = initTwoDimensionalArr(size)
		count = 0
	)

	for i := 0; i < size; i++ {
		for j := 0; j < size-i; j++ {
			x := j
			y := j + i

			if i <= 1 && s[x] == s[y] {
				temp[x][y] = true
				count++

			}

			if i > 1 && s[x] == s[y] && temp[x+1][y-1] {
				temp[x][y] = true
				count++
			}
		}
	}

	return count
}
