package main

import "fmt"

func main() {
	// SegmentTree
	seg := InitSegmentTree([]int{1, 9, 5, 3, 7, 2, 4, 6}, int(1e9), IntMergerMin)
	// [1, 4) = [9, 5, 3] = 3
	fmt.Println(seg.GetRange(1, 4))
	seg.UpdateAt(5, 6)
	// [4, 7) = [7, 6, 4] = 4
	fmt.Println(seg.GetRange(4, 7))

	// BIT
	bit := NewBit(8)
	bit.add(1, 1)
	bit.add(3, 3)
	bit.add(7, 7)
	// [1, 3] = [1, 0, 3] = 4
	fmt.Println(bit.sum(3))
	// [1, 8] = [1, 0, 3, 0, 0, 0, 7, 0] = 11
	fmt.Println(bit.sum(8))
}

// 1-indexed
type BIT struct {
	data []int
}

func NewBit(n int) *BIT {
	return &BIT{
		data: make([]int, n+1),
	}
}

func (tree *BIT) add(a, v int) {
	for i := a; i <= len(tree.data)-1; i += i & -i {
		tree.data[i] += v
	}
}

func (tree *BIT) sum(to int) int {
	res := 0
	for i := to; i > 0; i -= i & -i {
		res += tree.data[i]
	}
	return res
}

func IntMergerMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

type Merger func(a, b int) int

// 0-indexed
type SegmentTree struct {
	offset int
	data   []int
	inf    int
	merge  Merger
}

func InitSegmentTree(a []int, inf int, merge Merger) *SegmentTree {
	n := len(a)
	size := 1
	for size < n {
		size *= 2
	}
	data := make([]int, size*2)

	for j, i := 0, size-1; i < size+n-1; i++ {
		data[i] = a[j]
		j++
	}
	for i := size + n - 1; i < size*2; i++ {
		data[i] = inf
	}
	for i := size - 2; i >= 0; i-- {
		data[i] = merge(data[i*2+1], data[i*2+2])
	}

	return &SegmentTree{
		offset: size,
		data:   data,
		inf:    inf,
		merge:  merge,
	}
}

func (tree *SegmentTree) GetRange(from, to int) int {
	return tree.getRange(from, to, 0, 0, tree.offset)
}

func (tree *SegmentTree) getRange(from, to, index, left, right int) int {
	if to <= left || right <= from {
		return tree.inf
	}
	if from <= left && right <= to {
		return tree.data[index]
	}

	med := (left + right) / 2
	lvalue := tree.getRange(from, to, index*2+1, left, med)
	rvalue := tree.getRange(from, to, index*2+2, med, right)

	return tree.merge(lvalue, rvalue)
}

func (tree *SegmentTree) UpdateAt(index, value int) {
	idx := tree.offset - 1 + index
	tree.data[idx] = value
	for idx >= 1 {
		parent := (idx - 1) / 2
		left := parent*2 + 1
		right := parent*2 + 2

		tree.data[parent] = tree.merge(tree.data[left], tree.data[right])
		idx = parent
	}
}
