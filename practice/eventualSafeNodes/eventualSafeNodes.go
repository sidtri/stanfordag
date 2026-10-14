package main

import (
	"fmt"
	"sort"
)

func buildLL(graph [][]int) AdjList {
	adjList := make(AdjList)

	for n, v := range graph {
		adjList.Add(n, v)
	}

	return adjList
}

func eventualSafeNodes(graph [][]int) []int {
	g := buildLL(graph)
	l := g.Len()

	output := make(map[int]struct{})
	for i := range l {
		o := dfs(g, i)
		for _, ii := range o {
			output[ii] = struct{}{}
		}
	}
	keys := make([]int, 0, len(output))
	for k := range output {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys
}

func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func dfs(g AdjList, n int) []int {
	s := Stack{}
	s.Push(n)

	nodes := []int{}
	ignoreNodes := []int{}
	for {
		n := s.Pop()
		if n == -1 {
			break
		}
		if contains(ignoreNodes, n) {
			// already cycle detected earlier. No need to consider
			continue
		}

		if contains(nodes, n) {
			// cycle detected
			ignoreNodes = append(ignoreNodes, nodes...)
			nodes = []int{}
			continue
		}

		nodes = append(nodes, n)

		l := g[n]
		for _, aa := range l {
			s.Push(aa)
		}
	}

	return nodes
}

type Stack []int

func (s *Stack) Push(node int) {
	*s = append(*s, node)
}

func (s *Stack) Pop() int {
	lenS := len(*s)
	if lenS == 0 {
		return -1
	}
	node := (*s)[lenS-1]
	*s = (*s)[0:(lenS - 1)]
	return node
}

type AdjList map[int][]int

func (a *AdjList) Add(node int, val []int) {
	(*a)[node] = val
}

func (a AdjList) Len() int {
	return len(a)
}

func main() {
	input := [][]int{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}}
	fmt.Println("eventual safe nodes", eventualSafeNodes(input))
}
