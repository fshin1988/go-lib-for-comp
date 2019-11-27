package sort

import (
	"sort"
)

func reverse(s []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
}

func binarySearch(a []int, key int) int {
	ng := -1
	ok := len(a)
	for abs(ok-ng) > 1 {
		mid := (ok + ng) / 2
		if a[mid] >= key {
			ok = mid
		} else {
			ng = mid
		}
	}
	return ok
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

type Val struct {
	id, num int
}

type ByNum []Val

func (a ByNum) Len() int           { return len(a) }
func (a ByNum) Less(i, j int) bool { return a[i].num < a[j].num }
func (a ByNum) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
