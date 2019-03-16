package leetcode

import (
	"math"
)

func If(condition bool, a, b interface{}) interface{} {
	if condition {
		return a
	}

	return b
}

func clumsy(N int) int {
	var (
		sum    = 0
		prefix = 1
	)

	op := func(a, b, c, d int) int {
		x := int(math.Floor(float64(a * b / c)))

		return prefix*x + d
	}

	for N >= 4 {
		sum += op(N, N-1, N-2, N-3)

		N -= 4
		prefix = -1
	}

	if N == 3 {
		sum += 6 * prefix
	} else if N == 2 {
		sum += 2 * prefix
	} else if N == 1 {
		sum += 1 * prefix
	}

	return sum
}
