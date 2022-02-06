package problems

import "strings"

// https://leetcode.com/problems/number-of-segments-in-a-string

func countSegments(s string) int {
	s = strings.Trim(s, "\n\r\t ")
	if s == "" {
		return 0
	}

	sr := strings.Split(s, " ")

	data := make([]string, 0)
	for _, v := range sr {
		v = strings.Trim(v, "\n\r\t ")
		if v == "" {
			continue
		}
		data = append(data, v)
	}

	return len(data)
}
