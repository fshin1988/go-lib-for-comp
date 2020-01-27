package graph

import (
	"container/heap"
)

type Edge struct {
	to, cost int
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

func dijkstra(G [][]Edge, D []int) {
	pq := &PQ{}
	heap.Init(pq)
	D[0] = 0
	done := make([]bool, len(D))
	heap.Push(pq, Element{0, 0})
	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Element).id
		if done[cur] {
			continue
		}
		done[cur] = true
		for i := 0; i < len(G[cur]); i++ {
			next, cost := G[cur][i].to, G[cur][i].cost
			if done[next] {
				continue
			}
			if D[cur]+cost < D[next] {
				D[next] = D[cur] + cost
				heap.Push(pq, Element{next, D[next]})
			}
		}
	}
}
