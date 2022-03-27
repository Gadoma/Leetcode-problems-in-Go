package goleetcode0190

func uint32Pow(n, m uint32) uint32 {
	if m == 0 {
		return 1
	}

	var i, result uint32

	result = n
	for i = 2; i <= m; i++ {
		result *= n
	}

	return result
}

func reverseBits(num uint32) uint32 {
	var j, result uint32
	j = 31

	for i := 0; i < 32; i++ {
		result += uint32Pow(2, j) * (num % 2)
		num = num / 2
		j--
	}

	return result
}
