package main

import (
	"fmt"
)


// eventualSafeNodes returns, in ascending order, every node from which
// all paths lead to a terminal node (i.e. no cycle is reachable).
//
// It uses Kahn's algorithm on the reversed graph: a node is safe iff all
// of its outgoing neighbours are safe. Terminal nodes (out-degree 0) are
// safe by definition, and we peel safe nodes off layer by layer.
func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)

	// Reverse adjacency list and out-degree count.
	rev := make([][]int, n)
	outDeg := make([]int, n)
	for u, neighbours := range graph {
		outDeg[u] = len(neighbours)
		for _, v := range neighbours {
			rev[v] = append(rev[v], u)
		}
	}

	// Seed the queue with all terminal nodes.
	queue := make([]int, 0, n)
	for u := 0; u < n; u++ {
		if outDeg[u] == 0 {
			queue = append(queue, u)
		}
	}

	safe := make([]bool, n)
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]
		safe[u] = true

		for _, p := range rev[u] {
			outDeg[p]--
			if outDeg[p] == 0 {
				queue = append(queue, p)
			}
		}
	}

	result := make([]int, 0, n)
	for u := 0; u < n; u++ {
		if safe[u] {
			result = append(result, u)
		}
	}
	return result
}

func main() {
	input := [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}
	fmt.Println("eventual safe nodes", eventualSafeNodes(input))
}
