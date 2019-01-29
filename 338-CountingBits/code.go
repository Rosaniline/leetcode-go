package leetcode

// Given a non negative integer number num.
// For every numbers i in the range 0 ≤ i ≤ num calculate the number of 1's in their binary representation and return them as an array.
//
// Example:
// For num = 5 you should return [0,1,1,2,1,2].


func countBits(num int) []int {
	res := make([]int, num + 1)

	for i := 1; i < len(res); i ++ {
		res[i] = res[i >> 1] + i&1
	}

	return res
}