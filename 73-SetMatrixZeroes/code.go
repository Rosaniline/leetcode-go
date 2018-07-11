package leetcode


func setZeroes(matrix [][]int)  {
	if len(matrix) == 0 {
		return
	}

	var (
		m = len(matrix)
		n = len(matrix[0])
	)

	for i := 0; i < m; i ++ {
		for j := 0; j < n; j ++ {

			if matrix[i][j] == 0 {

				// using the leading cell of each row, col
				// to remember the should be set to zeros
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 0; i < m; i ++ {
		if matrix[i][0] == 0 {
			for j := 0; j < n; j ++ {
				matrix[i][j] = 0
			}
		}
	}

	for j := 0; j < n; j ++ {
		if matrix[0][j] == 0 {
			for i := 0; i < m; i ++ {
				matrix[i][j] = 0
			}
		}
	}
}