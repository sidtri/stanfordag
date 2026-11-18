package main

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {
	nodes := len(graph)
	var paths [][]int
	type stackEntry struct {
		val  int
		path []int
	}
	stack := []stackEntry{{val: 0, path: []int{0}}}

	for len(stack) > 0 {
		entry := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if entry.val == (nodes - 1) {
			paths = append(paths, entry.path)
			continue
		}

		neighbors := graph[entry.val]
		if len(neighbors) == 0 {
			continue
		}

		for _, neighbor := range neighbors {
			newPath := append([]int{}, entry.path...)
			newPath = append(newPath, neighbor)
			stack = append(stack, stackEntry{val: neighbor, path: newPath})
		}
	}
	return paths
}

func main() {
	fmt.Println("output:", allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
}
