package problems

// https://leetcode.com/problems/detect-capital

func detectCapitalUse(word string) bool {

	isCapitals := true
	isMainWord := true
	isLower := true

	if unicode.IsUpper(rune(word[0])) {
		isLower = false
	} else {
		isCapitals = false
	}

	for k := range word {
		if k == 0 {
			continue
		}

		char := rune(word[k])
		if unicode.IsUpper(char) {
			if isMainWord {
				isMainWord = false
			}
			if isLower {
				isLower = false
			}

			continue
		}

		if unicode.IsLower(char) {
			if isCapitals {
				isCapitals = false
			}
			continue
		}
	}

	return isCapitals || isMainWord || isLower
}
