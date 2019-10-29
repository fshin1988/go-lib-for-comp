package sort

import (
	"sort"
)

func reverse(s []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
}
