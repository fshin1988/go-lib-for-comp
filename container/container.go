package main

import (
	"container/heap"
	"fmt"
)

type Stack struct {
	arr []int
}

func (s *Stack) push(v int) {
	s.arr = append(s.arr, v)
}

func (s *Stack) pop() int {
	res := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return res
}

func (s *Stack) isEmpty() bool {
	return len(s.arr) == 0
}

type Queue struct {
	arr []int
}

func (q *Queue) push(v int) {
	q.arr = append(q.arr, v)
}

func (q *Queue) pop() int {
	res := q.arr[0]
	q.arr = q.arr[1:]
	return res
}

func (q *Queue) isEmpty() bool {
	return len(q.arr) == 0
}

type UnionFind struct {
	parent []int
}

func NewUnionFind(n int) *UnionFind {
	p := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
	}
	return &UnionFind{p}
}

func (a *UnionFind) root(x int) int {
	if a.parent[x] == x {
		return x
	} else {
		a.parent[x] = a.root(a.parent[x])
		return a.parent[x]
	}
}

func (a *UnionFind) isSame(x, y int) bool {
	return a.root(x) == a.root(y)
}

func (a *UnionFind) unite(x, y int) {
	x = a.root(x)
	y = a.root(y)
	if x == y {
		return
	} else {
		a.parent[x] = y
	}
}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	fmt.Println()
}
