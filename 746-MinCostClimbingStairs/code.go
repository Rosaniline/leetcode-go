package leetcode


// On a staircase, the i-th step has some non-negative cost cost[i] assigned (0 indexed).
// Once you pay the cost, you can either climb one or two steps.
// You need to find minimum cost to reach the top of the floor,
// and you can either start from the step with index 0, or the step with index 1.


func minCostClimbingStairs(cost []int) int {
	if len(cost) == 0 {
		return 0
	}

	if len(cost) == 1 {
		return cost[0]
	}

	dp := make([]int, len(cost))
	dp[0] = cost[0]
	dp[1] = cost[1]

	for i := 2; i < len(cost); i ++ {
		dp[i] = min(dp[i - 1], dp[i - 2]) + cost[i]
	}

	return min(dp[len(dp) - 1], dp[len(dp) - 2])
}


func min (a, b int) int {
	if a < b {
		return a
	}

	return b
}