package problems

var romanToIntMap = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	sarr := []rune(s)
	l := len(sarr) - 1

	v := 0
	nextChar := ' '

	for i := 0; i <= l; i++ {
		c := sarr[i]
		if i == l {
			nextChar = ' '
		} else {
			nextChar = sarr[i+1]
		}

		switch c {
		case 'M':
			v = v + 1000
		case 'D':
			v = v + 500
		case 'C':
			if nextChar == 'M' {
				v = v + 900
				i++
				continue
			}
			if nextChar == 'D' {
				v = v + 400
				i++
				continue
			}
			v = v + 100
		case 'L':
			v = v + 50
		case 'X':
			if nextChar == 'C' {
				v = v + 90
				i++
				continue
			}
			if nextChar == 'L' {
				v = v + 40
				i++
				continue
			}
			v = v + 10
		case 'V':
			v = v + 5
		case 'I':
			if nextChar == 'X' {
				v = v + 9
				i++
				continue
			}
			if nextChar == 'V' {
				v = v + 4
				i++
				continue
			}
			v = v + 1
		}

	}

	return v
}
