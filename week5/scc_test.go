package main

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

// bruteForceSCC computes true SCC sizes via reachability (O(V*(V+E))).
func bruteForceSCC(g AdjList, revGraph AdjList) []int {
	all := make(map[int]bool)
	for n := range g {
		all[n] = true
	}
	for n := range revGraph {
		all[n] = true
	}
	var nodes []int
	for n := range all {
		nodes = append(nodes, n)
	}

	reach := func(start int, adj AdjList) map[int]bool {
		seen := map[int]bool{start: true}
		stack := []int{start}
		for len(stack) > 0 {
			n := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			for _, next := range adj[n] {
				if !seen[next] {
					seen[next] = true
					stack = append(stack, next)
				}
			}
		}
		return seen
	}

	var sizes []int
	assigned := make(map[int]bool)
	for _, n := range nodes {
		if assigned[n] {
			continue
		}
		fwd := reach(n, g)
		back := reach(n, revGraph)
		var comp []int
		for m := range all {
			if fwd[m] && back[m] {
				comp = append(comp, m)
				assigned[m] = true
			}
		}
		sizes = append(sizes, len(comp))
	}
	return sizes
}

func TestSCCAgainstBruteForce(t *testing.T) {
	for trial := 0; trial < 200; trial++ {
		rng := rand.New(rand.NewSource(int64(trial)))
		n := rng.Intn(12) + 1

		g := make(AdjList)
		rev := make(AdjList)
		// ensure vertices exist as keys (some may be pure sinks)
		for v := 1; v <= n; v++ {
			g[v] = nil
			rev[v] = nil
		}
		for v := 1; v <= n; v++ {
			for w := 1; w <= n; w++ {
				if rng.Intn(4) == 0 {
					g.Add(v, w)
					rev.Add(w, v)
				}
			}
		}

		visit = make(map[int]bool)
		got := scc(g, rev)

		want := bruteForceSCC(g, rev)

		sort.Ints(got)
		sort.Ints(want)
		if !reflect.DeepEqual(got, want) {
			t.Fatalf("trial %d: got %v want %v\ngraph=%v\nrev=%v",
				trial, got, want, g, rev)
		}
	}
}
