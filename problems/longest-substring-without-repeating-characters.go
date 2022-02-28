package problems

// https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	sub := make(map[rune]int, 0)
	maxLen := 0
	var char rune

	l := len(s)
	lastIndex := l - 1

	for i := 0; i < l; i++ {
		char = rune(s[i])
		if index, ok := sub[char]; ok {
			sl := len(sub)
			if sl > maxLen {
				maxLen = sl
			}
			sub = make(map[rune]int, 0)
			i = index
			continue
		}
		sub[char] = i

		if i == lastIndex {
			sl := len(sub)
			if sl > maxLen {
				maxLen = sl
			}
		}
	}

	return maxLen
}
