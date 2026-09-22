package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// union-find over nodes 1..n
type UnionFind struct {
	parent []int
	rank   []int
	count  int // number of components
}

func NewUnionFind(n int) *UnionFind {
	uf := &UnionFind{
		parent: make([]int, n+1),
		rank:   make([]int, n+1),
		count:  n,
	}
	for i := 1; i <= n; i++ {
		uf.parent[i] = i
	}
	return uf
}

func (uf *UnionFind) Find(x int) int {
	for uf.parent[x] != x {
		uf.parent[x] = uf.parent[uf.parent[x]] // path halving
		x = uf.parent[x]
	}
	return x
}

// Union returns true if the two nodes were in different components.
func (uf *UnionFind) Union(a, b int) bool {
	ra, rb := uf.Find(a), uf.Find(b)
	if ra == rb {
		return false
	}
	if uf.rank[ra] < uf.rank[rb] {
		ra, rb = rb, ra
	}
	uf.parent[rb] = ra
	if uf.rank[ra] == uf.rank[rb] {
		uf.rank[ra]++
	}
	uf.count--
	return true
}

// clustering runs Kruskal's algorithm, stopping when k clusters remain,
// and returns the maximum spacing (cost of the next edge that would merge two clusters).
func clustering(edges *Edges, n int, k int) int {
	uf := NewUnionFind(n)

	for edges.Len() > 0 {
		e := heap.Pop(edges).(Edge)

		if uf.Union(e.source, e.destination) {
			if uf.count == k {
				// the next edge that merges two clusters is the max spacing
				for edges.Len() > 0 {
					next := heap.Pop(edges).(Edge)
					if uf.Find(next.source) != uf.Find(next.destination) {
						return next.cost
					}
				}
			}
		}
	}

	return 0
}

type Edge struct {
	source      int
	destination int
	cost        int
}

type Edges []Edge

// implement  edges heap
func (h Edges) Len() int {
	return len(h)
}

func (h Edges) Less(i, j int) bool {
	return h[i].cost < h[j].cost
}

func (h Edges) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *Edges) Push(x any) {
	*h = append(*h, x.(Edge))
}

func (h *Edges) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}


func main() {
	dataBytes, err := os.ReadFile("clustering1.txt")
	if err != nil {
		return
	}

	data := strings.Split(strings.TrimSpace(string(dataBytes)), "\n")
	if err != nil {
		return
	}

	totalNodes, _ := strconv.Atoi(strings.TrimSpace(data[0]))

	edges := Edges{}
	for _, a := range data[1:] {
		dd := strings.Fields(a)
		s, _ := strconv.Atoi(strings.TrimSpace(dd[0]))
		d, _ := strconv.Atoi(strings.TrimSpace(dd[1]))
		c, _ := strconv.Atoi(strings.TrimSpace(dd[2]))
		edges = append(edges, Edge{s, d, c})
	}
	heap.Init(&edges)

	const k = 4
	fmt.Println(clustering(&edges, totalNodes, k))
}
