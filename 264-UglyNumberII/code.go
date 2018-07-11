package _264


func min (a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func nthUglyNumber(n int) int {
	var (
		i = 0
		j = 0
		k = 0
		res = []int{1}
	)

	for len(res) < n {
		res = append(res, min(min(res[i]*2, res[j]*3), res[k]*5))
		if res[len(res) - 1] == res[i]*2 {
			i ++
		}

		if res[len(res) - 1] == res[j]*3 {
			j ++
		}

		if res[len(res) - 1] == res[k]*5 {
			k ++
		}
	}

	return res[len(res) - 1]
}
