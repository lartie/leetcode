package problems

import "sort"

// https://leetcode.com/problems/average-salary-excluding-the-minimum-and-maximum-salary

func average(salary []int) float64 {
	sum := 0
	l := len(salary)
	lastIndex := l-1

	sort.Ints(salary)

	for k := range salary {
		if k == 0 || k == lastIndex {
			continue
		}

		sum += salary[k]
	}
	return float64(sum) / float64(l-2)
}
