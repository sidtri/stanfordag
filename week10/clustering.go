package main

import (
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Cluster []Edge

func clustering(edges *Edges) {
	lenClusters := 0
	clusters := Clusters{}

	for edges.Len() > 0 {
		lowEdge := heap.Pop(&edges).(Edge)

		clusters[lowEdge.Name()] = append(clusters[lowEdge.Name()], lowEdge.Name())
	}
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

func (e Edge) Name() string {
	return string(e.source) + "-" + string(e.destination)
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

	// totalNodes, _ := strconv.Atoi(strings.TrimSpace(data[0]))
	// adjList := map[int]Edges{}
	edges := Edges{}
	for _, a := range data[1:] {
		dd := strings.Fields(a)
		s, _ := strconv.Atoi(strings.TrimSpace(dd[0]))
		d, _ := strconv.Atoi(strings.TrimSpace(dd[1]))
		c, _ := strconv.Atoi(strings.TrimSpace(dd[2]))
		edges = append(edges, Edge{s, d, c})
		heap.Init(&edges)
		// adjList[s] = append(adjList[s], Edge{s, d, c})
	}

	fmt.Println(clustering(&edges))
}
