package leetcode

import (
	"sort"
	"fmt"
	"strings"
)


type Number []int


func (num Number) Len() int {
	return len(num)
}

func (num Number) Less(i, j int) bool {
	si := fmt.Sprintf("%d%d", num[i], num[j])
	sj := fmt.Sprintf("%d%d", num[j], num[i])

	return strings.Compare(si, sj) < 0
}

func (num Number) Swap(i, j int) {
	num[i], num[j] = num[j], num[i]
}


func largestNumber(nums []int) string {
	sort.Sort(sort.Reverse(Number(nums)))

	res := ""

	for _, n := range nums {
		if res == "0" && n == 0 {
			continue
		}

		res = fmt.Sprintf("%v%v", res, n)
	}

	return res
}