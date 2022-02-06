package problems

// https://leetcode.com/problems/binary-search

const notFound = -1

func search(nums []int, target int) int {
	return search2(nums, 0, target)
}

func search2(nums []int, offset int, target int) int {
	l := len(nums)
	if l == 1 {
		if nums[0] == target {
			return offset
		}
		return notFound
	}

	mid := calcMid(l)

	v := nums[mid]
	if v < target {
		return search2(nums[mid:], mid+offset, target)
	}
	if v > target {
		return search2(nums[:mid], offset, target)
	}


	return offset + mid
}

func calcMid(l int) int {
	mid := notFound
	if l > 0 {
		mid = int(l / 2)
	}
	return mid
}
