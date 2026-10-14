package main

import "fmt"

func eventualSafeNodes(graph [][]int) []int {
	visited := make(map[int]bool)
	cycle := make(map[int]bool)
	output := []int{}
	for i := range 7 {
		isc := dfs(graph, 0, visited, cycle)
		if !isc {
			output = append(output, i)
		}
	}
	return output
}

func dfs(graph [][]int, index int, visited map[int]bool, cycle map[int]bool) bool {
	if visited[index] {
		cycle[index] = true
		return true
	}

	visited[index] = true
	for _, i := range graph[index] {
		cycle := dfs(graph, i, visited, cycle)
		if cycle {
			return true
		}
	}

	return false
}

func main() {
	input := [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}
	fmt.Println("eventual safe nodes", eventualSafeNodes(input))
}
