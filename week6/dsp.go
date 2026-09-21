package main

import (
	"container/heap"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Edge struct {
	node   int
	weight int
}

type pqItem struct {
	node     int
	distance int
}

type PriorityQueue []pqItem

func (h PriorityQueue) Len() int           { return len(h) }
func (h PriorityQueue) Less(i, j int) bool { return h[i].distance < h[j].distance }
func (h PriorityQueue) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *PriorityQueue) Push(x any) {
	*h = append(*h, x.(pqItem))
}

func (h *PriorityQueue) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func dsp(nodes map[int][]Edge) map[int]int {
	distances := map[int]int{}
	pq := PriorityQueue{}
	heap.Init(&pq)
	var firstNode int
	for node := range nodes {
		if node == 1 {
			firstNode = node
			distances[node] = 0
			continue
		}
		distances[node] = math.MaxInt
	}
	heap.Push(&pq, pqItem{firstNode, 0})

	for pq.Len() > 0 {
		currentNode := heap.Pop(&pq).(pqItem)
		if currentNode.distance != distances[currentNode.node] {
			// current node path is decreased so we continue next step
			continue
		}
		for _, edge := range nodes[currentNode.node] {
			currentDistance := distances[currentNode.node] + edge.weight
			if currentDistance < distances[edge.node] {
				distances[edge.node] = currentDistance
				heap.Push(&pq, pqItem{edge.node, currentDistance})
			}
		}
	}

	return distances
}

func main() {
	bytes, err := os.ReadFile("dijkstraData.txt")
	if err != nil {
		fmt.Println("Error   at bytes   parsing")
		panic(err)
	}

	raw := strings.Split(strings.TrimSpace(string(bytes)), "\n")

	nodes := make(map[int][]Edge)

	for _, d := range raw {
		dd := strings.Split(strings.TrimSpace(d), "\t")

		var edges []Edge

		for _, ed := range dd[1:] {
			edd := strings.Split(ed, ",")
			node, err := strconv.Atoi(strings.TrimSpace(edd[0]))
			if err != nil {
				fmt.Println("Error   at d", d)
				panic(err)
			}

			edge, err := strconv.Atoi(strings.TrimSpace(edd[1]))
			if err != nil {
				fmt.Println("Error   at d", d)
				panic(err)
			}

			edges = append(edges, Edge{node, edge})
		}

		mainNode, err := strconv.Atoi(dd[0])
		if err != nil {
			fmt.Println("Error   at d", d)
			panic(err)
		}

		nodes[mainNode] = edges
	}

	testNodes := []int{7, 37, 59, 82, 99, 115, 133, 165, 188, 197}
	answer := dsp(nodes)
	for _, n := range testNodes {
		fmt.Printf("%d,", answer[n])
	}
}
