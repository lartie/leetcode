package problems

// https://leetcode.com/problems/unique-number-of-occurrences

func uniqueOccurrences(arr []int) bool {
	s := make(map[int]int, len(arr))

	for k := range arr {
		key := arr[k]

		if _, ok := s[key]; ok {
			s[key]++
		} else {
			s[key] = 0
		}
	}

	set := make(map[int]bool, len(s))
	for k := range s {
		if _, ok := set[s[k]]; ok {
			return false
		}

		set[s[k]] = true
	}
	return true
}
