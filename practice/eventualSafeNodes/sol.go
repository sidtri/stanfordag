package main

import (
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

	ignoreNodes := map[int]bool{}
	for i := range l {
		dfs(g, i, ignoreNodes)
	}
	output := []int{}
	for i := range l {
		if ignoreNodes[i] {
			continue
		}
		output = append(output, i)
	}
	sort.Ints(output)
	return output
}

func contains(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func dfs(g AdjList, n int, ignoreNodes map[int]bool) []int {
	s := Stack{}
	s.Push(PNode{val: n})

	for {
		n := s.Pop()
		if n.isEmpty {
			break
		}
		if ignoreNodes[n.val] {
			// already cycle detected earlier. No need to consider
			for _, nval := range n.backtrace {
				ignoreNodes[nval] = true
			}
			continue
		}

		if contains(n.backtrace, n.val) {
			// cycle detected
			for _, ii := range n.backtrace {
				ignoreNodes[ii] = true
			}
			continue
		}

		l := g[n.val]
		for _, aa := range l {
			bt := append(n.backtrace, n.val)
			pn := PNode{val: aa, backtrace: bt}
			s.Push(pn)
		}
	}

	return []int{}
}

type PNode struct {
	val       int
	backtrace []int
	isEmpty   bool
}

type Stack []PNode

func (s *Stack) Push(n PNode) {
	*s = append(*s, n)
}

func (s *Stack) Pop() PNode {
	lenS := len(*s)
	if lenS == 0 {
		return PNode{isEmpty: true}
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

