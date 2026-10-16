package main

import "fmt"

type Visit map[int]bool

func findCircleNum(isConnected [][]int) int {
	totalProvinces := 0

	visited := Visit{}

	for index := range isConnected {
		if visited[index] {
			continue
		}
		dfs(isConnected, index, visited)
		totalProvinces++
	}
	return totalProvinces
}

func dfs(isConnected [][]int, node int, visited Visit) {
	visited[node] = true

	for neighbour, connected := range isConnected[node] {
		if connected == 0 {
			continue
		}

		if visited[neighbour] {
			continue
		}

		dfs(isConnected, neighbour, visited)
	}
}

func main() {
	provinces := findCircleNum([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}})
	fmt.Println("totalProvinces", provinces)
}
