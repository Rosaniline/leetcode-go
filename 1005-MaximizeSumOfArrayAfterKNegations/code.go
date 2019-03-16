package leetcode

import (
	"sort"
)

func sum(x []int) int {
	sum := 0
	for _, num := range x {
		sum += num
	}

	return sum
}

func min(x []int) int {
	min := x[0]

	for _, v := range x {
		if min > v {
			min = v
		}
	}

	return min
}

func largestSumAfterKNegations(A []int, K int) int {
	sort.Ints(A)

	i := 0

	for i < len(A) && i < K && A[i] < 0 {
		A[i] = -A[i]
		i += 1
	}

	return sum(A) - (K-i)%2*min(A)*2
}
