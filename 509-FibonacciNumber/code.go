package leetcode

import "math/big"

func fib(n int) *big.Int {
	if n < 2 {
		return big.NewInt(int64(n))
	}

	p := big.NewInt(0)
	q := big.NewInt(1)

	for i := 2; i <= n; i++ {
		q.Add(p, q)
		p.Sub(q, p)
	}

	return q
}
