package problems

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	xs := intToSlice(x)
	res := true
	l := len(xs) - 1
	j := l
	for i := 0; i <= l; i++ {
		if i == j {
			break
		}
		if xs[i] == xs[j] {
			j = j - 1
			continue
		} else {
			res = false
			break
		}
	}

	return res
}

func intToSlice(x int) []int {
	res := make([]int, 0)
	for {
		res = append(res, x%10)
		x = x/10
		if x < 1 {
			break
		}
	}

	return res
}
