package main

import (
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 14, 32, 51, 51, 51, 243, 419, 750, 910}
	fmt.Println(binarySearch(arr, 51))
}

func isOk(a []int, index, key int) bool {
	if a[index] >= key {
		return true
	}
	return false
}

func binarySearch(a []int, key int) int {
	// ng := len(a)
	// ok := -1
	ng := -1
	ok := len(a)
	for abs(ok-ng) > 1 {
		mid := (ok + ng) / 2
		if isOk(a, mid, key) {
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

func reverse(s []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
}
