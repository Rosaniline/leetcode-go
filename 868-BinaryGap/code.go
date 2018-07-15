package leetcode



func binaryGap(N int) int {
	var (
		localDist 	= 0
		maxDist 	= 0
		prev		= false
	)

	for N > 0 {
		if N & 1 == 1 {
			if prev && localDist > maxDist {
				maxDist = localDist
			}

			prev = true
			localDist = 0
		}

		localDist ++
		N >>= 1
	}

	return maxDist
}