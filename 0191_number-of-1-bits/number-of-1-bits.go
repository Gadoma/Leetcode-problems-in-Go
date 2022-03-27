package goleetcode0191

func hammingWeight(num uint32) int {
	var cnt uint32
	for i := 0; i < 32; i++ {
		cnt += 1 * (num % 2)
		num = num / 2
	}

	return int(cnt)
}
