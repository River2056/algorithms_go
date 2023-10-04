package countthedigit

func NbDig(n int, d int) int {
    ref := make([]int, 0)
    res := 0

    for i := 0; i <= n; i++ {
        ref = append(ref, i * i)
    }

    for _, val := range ref {
        num := val
        for num > 0 {
            r := num % 10
            if r == d {
                res++
            }
            num /= 10
        }
    }

    if d == 0 {
        res++
    }

	return res
}
