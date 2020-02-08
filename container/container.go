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

type Node struct {
	x, y int
}

type Queue struct {
	arr []Node
}

func (q *Queue) push(v Node) {
	q.arr = append(q.arr, v)
}

func (q *Queue) pop() Node {
	res := q.arr[0]
	q.arr = q.arr[1:]
	return res
}

func (q *Queue) isEmpty() bool {
	return len(q.arr) == 0
}

type UnionFind struct {
	parent []int
	size   []int
}

func NewUnionFind(n int) *UnionFind {
	p := make([]int, n)
	s := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = i
		s[i] = 1
	}
	return &UnionFind{p, s}
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
		a.size[y] += a.size[x]
	}
}

func (a *UnionFind) getSize(x int) int {
	return a.size[a.root(x)]
}

type Element struct {
	id, cost int
}

type PQ []Element

func (pq PQ) Len() int           { return len(pq) }
func (pq PQ) Less(i, j int) bool { return pq[i].cost < pq[j].cost }
func (pq PQ) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PQ) Push(x interface{}) {
	*pq = append(*pq, x.(Element))
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func main() {
	pq := &PQ{}
	heap.Init(pq)
	heap.Push(pq, Element{0, 3})
	heap.Push(pq, Element{1, 5})
	heap.Push(pq, Element{2, 1})
	for pq.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(pq).(Element))
	}
	fmt.Println()
}
