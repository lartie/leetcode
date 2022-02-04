package problems

// https://leetcode.com/problems/consecutive-characters

func maxPower(s string) int {
	l := len(s)
	max := 0
	if l == 0 {
		return max
	}

	max = 1
	if l == 1 {
		return max
	}

	prevChar := ' '
	tmpMax := 1
	for k, v := range []rune(s) {
		if k == 0 {
			prevChar = v
			continue
		}

		if prevChar == v {
			tmpMax++
		} else {
			tmpMax = 1
		}

		if max < tmpMax {
			max = tmpMax
		}
		prevChar = v
	}

	return max
}
