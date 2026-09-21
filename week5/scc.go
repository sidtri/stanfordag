package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Stack []Phase

func (s *Stack) Pop() Phase {
	l := len(*s)
	el := (*s)[l-1]
	*s = (*s)[0:(l - 1)]

	return el
}

func (s *Stack) Push(el Phase) {
	*s = append(*s, el)
}

type AdjList map[int][]int

func (g AdjList) Add(key int, value int) {
	g[key] = append(g[key], value)
}

type Visit map[int]bool

var visit Visit = make(map[int]bool)

type Phase struct {
	node     int
	finished bool
}

func dfs(gph AdjList, node int) Stack {
	var stack Stack
	stack = []Phase{}

	var outStack Stack
	outStack = []Phase{}

	stack = append(stack, Phase{node, false})

	for len(stack) > 0 {
		phase := stack.Pop()
		node := phase.node

		if phase.finished {
			outStack.Push(Phase{phase.node, false})
			continue
		}

		if visit[node] {
			continue
		}

		visit[node] = true
		stack.Push(Phase{node, true})

		for _, edge := range gph[node] {
			if !visit[edge] {
				stack.Push(Phase{edge, false})
			}
		}

	}

	return outStack
}

func scc(g AdjList, revGraph AdjList) []int {
	var finishOrder Stack

	for node := range g {
		if !visit[node] {
			part := dfs(g, node)
			finishOrder = append(finishOrder, part...)
		}
	}

	for revNode := range revGraph {
		if !visit[revNode] {
			part := dfs(g, revNode)
			finishOrder = append(finishOrder, part...)
		}
	}

	visit = make(map[int]bool)
	var sizes []int

	for fo := len(finishOrder) - 1; fo >= 0; fo-- {
		node := finishOrder[fo].node
		if !visit[node] {
			revPartArray := dfs(revGraph, node)
			sizes = append(sizes, len(revPartArray))
		}
	}

	return sizes
}

func main() {
	data, err := os.ReadFile("scc.txt")
	if err != nil {
		fmt.Println("failed to load file")
		return
	}

	var graph AdjList
	graph = make(AdjList)

	var revGraph AdjList
	revGraph = make(AdjList)

	d := strings.Split(strings.TrimSpace(string(data)), "\n")

	for _, a := range d {
		dd := strings.Split(a, " ")
		v, _ := strconv.Atoi(dd[0])
		e, _ := strconv.Atoi(dd[1])

		graph.Add(v, e)
		revGraph.Add(e, v)
	}

	for key := range graph {
		visit[key] = false
	}

	output := scc(graph, revGraph)

	sort.Sort(sort.Reverse(sort.IntSlice(output)))
	fmt.Println(output[:5])
}
