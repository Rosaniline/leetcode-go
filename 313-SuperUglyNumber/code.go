package _13_SuperUglyNumber

import "golang.org/x/tools/container/intsets"


func min (a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func nthSuperUglyNumber(n int, primes []int) int {
	var (
		param = make([]int, len(primes))
		res = []int{1}
	)

	for len(res) < n {
		minVal := intsets.MaxInt

		for i := range param {
			minVal = min(minVal, res[param[i]]*primes[i])
		}

		res = append(res, minVal)

		for i := range param {
			if minVal == res[param[i]]*primes[i] {
				param[i] ++
			}
		}
	}

	return res[len(res) - 1]
}