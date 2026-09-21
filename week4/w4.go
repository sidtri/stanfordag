package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.ReadFile("mincut.txt")

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")

	adjList := make([]map[string][]string, 0)

	for _, l := range lines {
		d := strings.Fields(l)
		l := map[string][]string{
			d[0]: d[1:],
		}
		adjList = append(adjList, l)
	}

	mc := minCut(adjList)

	fmt.Println(mc)
}

func minCut(adjList []map[string][]string) (mc int) {
	edges := buildEdges(adjList)
	n := 200

	mc = len(edges)
	for i := 0; i < n*n*3; i++ {
		c := oneTrial(edges, n)
		if c < mc {
			mc = c
		}
	}
	return mc
}

func oneTrial(edges [][2]int, n int) int {
	parent := make([]int, n+1)
	for i := 1; i <= n; i++ {
		parent[i] = i
	}

	components := n
	for components > 2 {
		e := edges[rand.Intn(len(edges))]
		if union(parent, e[0], e[1]) {
			components--
		}
	}

	cuts := 0
	for _, e := range edges {
		if find(parent, e[0]) != find(parent, e[1]) {
			cuts++
		}
	}
	return cuts
}

func find(parent []int, x int) int {
	for parent[x] != x {
		x = parent[x]
	}
	return x
}

func union(parent []int, a, b int) bool {
	ra, rb := find(parent, a), find(parent, b)
	if ra == rb {
		return false
	}
	parent[ra] = rb
	return true
}

func buildEdges(adjList []map[string][]string) [][2]int {
	seen := make(map[[2]int]bool) // remember edges we've already added
	var edges [][2]int
	for _, val := range adjList {
		for key, val := range val {
			vtx, _ := strconv.Atoi(key)
			for _, edge := range val {
				e, _ := strconv.Atoi(edge)
				pair := [2]int{vtx, e}
				if seen[pair] {
					continue
				}
				seen[pair] = true
				seen[[2]int{e, vtx}] = true
				edges = append(edges, pair)
			}
		}
	}
	return edges
}
