package sort

import (
	"sort"
)

func reverse(s []int) {
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
}

type RuneSlice []rune

func (p RuneSlice) Len() int           { return len(p) }
func (p RuneSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p RuneSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
