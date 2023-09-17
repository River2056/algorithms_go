package sumdigitpow

func getPow(target, pow uint64) uint64 {
	result := target
	var i uint64 = 1
	for ; i < pow; i++ {
		result *= target
	}
	return result
}

func SumDigitPow(a, b uint64) []uint64 {
	res := make([]uint64, 0)

	for i := a; i <= b; i++ {
		var total uint64 = 0
		num := i
		calculatePow := i
		var pow uint64 = 1
		for calculatePow > 10 {
			calculatePow /= 10
			pow++
		}

		for num > 0 {
			remainder := num % 10
			total += getPow(remainder, pow)
			pow--
			num /= 10
		}

		if total == i {
			res = append(res, i)
		}
	}

	return res
}
