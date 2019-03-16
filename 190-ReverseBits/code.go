package leetcode

func reverseBits(num uint32) uint32 {
	res := num & 1

	for i := 0; i < 31; i++ {
		num >>= 1
		res = res<<1 + num&1
	}

	return res
}
