package leetcode

// You are climbing a stair case. It takes n steps to reach to the top.
//
// Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?
//
// Note: Given n will be a positive integer.


func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	var (
		prev	= 2
		pprev	= 1
	)

	for i := 3; i <= n; i ++ {
		prev = prev + pprev
		pprev = prev - pprev
	}

	return prev
}