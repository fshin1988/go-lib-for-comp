package main

import (
	"container/heap"
	"fmt"
)

type NodeHeap []*Node

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].cost < h[j].cost }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}

func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type Node struct {
	edgesTo   []int
	edgesCost []int
	done      bool
	cost      int
}

type Graph struct {
	nodes []*Node
}

func NewGraph() *Graph {
	return &Graph{}
}

func (g *Graph) addNode(edgesTo, edgesCost []int, isStart bool) {
	cost := -1
	if isStart {
		cost = 0
	}
	g.nodes = append(g.nodes, &Node{edgesTo, edgesCost, false, cost})
}

func (g *Graph) search() {
	h := &NodeHeap{}
	heap.Init(h)
	heap.Push(h, g.nodes[0])
	for h.Len() > 0 {
		doneNode := heap.Pop(h).(*Node)
		if doneNode.done {
			continue
		}
		doneNode.done = true
		for i := 0; i < len(doneNode.edgesTo); i++ {
			to := doneNode.edgesTo[i]
			cost := doneNode.cost + doneNode.edgesCost[i]
			if g.nodes[to].cost < 0 || cost < g.nodes[to].cost {
				g.nodes[to].cost = cost
				heap.Push(h, g.nodes[to])
			}
		}
	}
}

func main() {
	g := NewGraph()
	// Start Node
	g.addNode([]int{1, 2, 3}, []int{5, 4, 2}, true)
	g.addNode([]int{0, 3, 5}, []int{5, 2, 6}, false)
	g.addNode([]int{0, 1, 3, 4}, []int{4, 2, 3, 2}, false)
	g.addNode([]int{0, 2, 4}, []int{2, 3, 6}, false)
	g.addNode([]int{2, 3, 5}, []int{2, 6, 4}, false)
	g.addNode([]int{1, 4}, []int{6, 4}, false)
	g.search()
	for i, n := range g.nodes {
		fmt.Println(i, n.cost)
	}
}
