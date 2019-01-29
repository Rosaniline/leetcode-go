package leetcode

func min(a, b, c int) int {
	inner := func(x, y int) int {
		if x < y {
			return x
		}

		return y
	}

	return inner(inner(a, b), c)
}

func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	var (
		m       = len(matrix)
		n       = len(matrix[0])
		dp      = make([][]int, 0)
		maxSize = 0
	)

	for i := 0; i < m; i++ {
		dp = append(dp, make([]int, n))
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if (i == 0 || j == 0) && matrix[i][j] == '1' {
				dp[i][j] = 1
			}

			if i > 0 && j > 0 {
				if matrix[i][j] == '0' {
					dp[i][j] = 0

				} else {
					temp := min(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1

					if temp > dp[i][j] {
						dp[i][j] = temp
					}
				}
			}

			if dp[i][j] > maxSize {
				maxSize = dp[i][j]
			}
		}
	}

	return maxSize * maxSize
}
