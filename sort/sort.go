package sort

import (
	"sort"
)

func reverse(s []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
}

func binarySearch(nums []int, key int) int {
	left := -1
	right := len(nums)
	for right-left > 1 {
		mid := left + (right-left)/2
		if nums[mid] >= key {
			right = mid
		} else {
			left = mid
		}
	}
	return right
}

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
