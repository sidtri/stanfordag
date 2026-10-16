package main

import (
	"fmt"
	"math/rand"
	"reflect"
)

// brute force: a node is unsafe iff some cycle is reachable from it
func bruteSafe(graph [][]int) []int {
	n := len(graph)
	reach := make([][]bool, n)
	for i := range reach {
		reach[i] = make([]bool, n)
		st := []int{i}
		for len(st) > 0 {
			u := st[len(st)-1]
			st = st[:len(st)-1]
			for _, w := range graph[u] {
				if !reach[i][w] {
					reach[i][w] = true
					st = append(st, w)
				}
			}
		}
	}
	cycle := make([]bool, n)
	for v := 0; v < n; v++ {
		if reach[v][v] {
			cycle[v] = true
		}
	}
	unsafe := make([]bool, n)
	for u := 0; u < n; u++ {
		for c := 0; c < n; c++ {
			if cycle[c] && reach[u][c] {
				unsafe[u] = true
				break
			}
		}
	}
	out := []int{}
	for i := 0; i < n; i++ {
		if !unsafe[i] {
			out = append(out, i)
		}
	}
	return out
}

func main() {
	fixed := [][][]int{
		{{1, 2}, {2, 3}, {5}, {0}, {5}, {}, {}},
		{{1, 2}, {2}, {3}, {2}},
		{{0}, {2, 3, 4}, {3, 4}, {0, 4}, {}},
	}
	for _, in := range fixed {
		got := eventualSafeNodes(in)
		want := bruteSafe(in)
		fmt.Printf("input %v => got %v want %v ok=%v\n", in, got, want, reflect.DeepEqual(got, want))
	}

	// random fuzz
	rng := rand.New(rand.NewSource(1))
	fails := 0
	for iter := 0; iter < 20000; iter++ {
		n := 1 + rng.Intn(8)
		g := make([][]int, n)
		for u := 0; u < n; u++ {
			for v := 0; v < n; v++ {
				if rng.Intn(4) == 0 { // ~25% edge prob, incl self-loops
					g[u] = append(g[u], v)
				}
			}
		}
		got := eventualSafeNodes(g)
		want := bruteSafe(g)
		if !reflect.DeepEqual(got, want) {
			fails++
			if fails <= 5 {
				fmt.Printf("FAIL input %v => got %v want %v\n", g, got, want)
			}
		}
	}
	fmt.Println("random failures:", fails, "/ 20000")
}
