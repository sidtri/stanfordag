package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type EdgeCost struct {
	source      int
	destination int
	cost        int
}

type EdgeCostHeap []EdgeCost

func (h EdgeCostHeap) Len() int {
	return len(h)
}

func (h EdgeCostHeap) Less(i, j int) bool {
	return h[i].cost < h[j].cost
}

func (h EdgeCostHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *EdgeCostHeap) Push(x any) {
	*h = append(*h, x.(EdgeCost))
}

func (h *EdgeCostHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func mst(adjList *map[int][]EdgeCost, n int) int {
	costs := 0

	visited := map[int]bool{}
	visited[1] = true
	edgeCosts := EdgeCostHeap{}
	edgeCosts = append(edgeCosts, (*adjList)[1]...)
	heap.Init(&edgeCosts)

	for len(visited) <= n-1 {
		minEdge := heap.Pop(&edgeCosts).(EdgeCost)
		if visited[minEdge.destination] {
			continue
		}

		costs += minEdge.cost
		for _, edge := range (*adjList)[minEdge.destination] {
			heap.Push(&edgeCosts, edge)
		}
		visited[minEdge.destination] = true
	}

	return costs
}

func main() {
	data, err := os.ReadFile("edges.txt")
	if err != nil {
		return
	}

	d := strings.Split(strings.TrimSpace(string(data)), "\n")

	firstLine := strings.Fields(strings.TrimSpace(d[0]))
	n, _ := strconv.Atoi(firstLine[0])
	// nedges, _ := strconv.Atoi(firstLine[1])

	adjList := map[int][]EdgeCost{}
	for _, a := range d[1:] {
		dd := strings.Fields(a)
		s, _ := strconv.Atoi(strings.TrimSpace(dd[0]))
		d, _ := strconv.Atoi(strings.TrimSpace(dd[1]))
		c, _ := strconv.Atoi(strings.TrimSpace(dd[2]))
		adjList[s] = append(adjList[s], EdgeCost{s, d, c})
		adjList[d] = append(adjList[d], EdgeCost{d, s, c})
	}

	fmt.Println("cost is", mst(&adjList, n))
}
