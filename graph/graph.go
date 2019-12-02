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
	id           int
	edgesTo      []int
	edgesCost    []int
	done         bool
	cost, fromId int
}

type Graph struct {
	nodes []*Node
}

func NewGraph() *Graph {
	return &Graph{}
}

func (g *Graph) addNode(id int, edgesTo, edgesCost []int) {
	cost := -1
	if id == 0 {
		cost = 0
	}
	g.nodes = append(g.nodes, &Node{id, edgesTo, edgesCost, false, cost, 0})
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
				g.nodes[to].fromId = doneNode.id
				g.nodes[to].cost = cost
				heap.Push(h, g.nodes[to])
			}
		}
	}
}

func main() {
	g := NewGraph()
	// Start Node
	g.addNode(0, []int{1, 2, 3}, []int{5, 4, 2})
	g.addNode(1, []int{0, 3, 5}, []int{5, 2, 6})
	g.addNode(2, []int{0, 1, 3, 4}, []int{4, 2, 3, 2})
	g.addNode(3, []int{0, 2, 4}, []int{2, 3, 6})
	g.addNode(4, []int{2, 3, 5}, []int{2, 6, 4})
	g.addNode(5, []int{1, 4}, []int{6, 4})
	g.search()
	for _, n := range g.nodes {
		fmt.Println("id:", n.id, "cost:", n.cost, "fromId:", n.fromId)
	}
}
