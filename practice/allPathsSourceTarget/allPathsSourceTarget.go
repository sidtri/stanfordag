package main

import "fmt"

func allPathsSourceTarget(graph [][]int) [][]int {
	target := len(graph) - 1
	var paths [][]int
	path := []int{0}

	var dfs func(node int)
	dfs = func(node int) {
		if node == target {
			paths = append(paths, append([]int{}, path...))
			return
		}
		for _, neighbor := range graph[node] {
			path = append(path, neighbor)
			dfs(neighbor)
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return paths
}

func main() {
	fmt.Println("output:", allPathsSourceTarget([][]int{{1, 2}, {3}, {3}, {}}))
}
