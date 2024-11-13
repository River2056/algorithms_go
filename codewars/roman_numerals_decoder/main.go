package romannumeralsdecoder

func Decode(roman string) int {
	mapping := map[string]int{
		"M": 1000,
		"D": 500,
		"C": 100,
		"L": 50,
		"X": 10,
		"V": 5,
		"I": 1,
	}

	arr := make([]int, 0)
	for i := 0; i < len(roman); i++ {
		n := mapping[string(roman[i])]
		if i != 0 {
			m := arr[i-1]
			if m < n {
				arr[i-1] = -m
			}
		}

		arr = append(arr, n)
	}

	sum := 0
	for _, v := range arr {
		sum += v
	}

	return sum
}
