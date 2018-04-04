package main

import "fmt"

type Vertex int

type Graph struct {
	v   Vertex
	adj map[Vertex]map[Vertex]bool
}

func NewGraph(v Vertex) *Graph {
	return &Graph{
		v:   v,
		adj: make(map[Vertex]map[Vertex]bool),
	}
}

func (g *Graph) AddEdge(v, w Vertex) {
	if v >= g.v || w >= g.v {
		panic("Tried to define edge with out-of-bounds vertex")
	}
	if _, exists := g.adj[v]; !exists {
		g.adj[v] = make(map[Vertex]bool)
	}
	g.adj[v][w] = true
}

func (g *Graph) walkGraph(v Vertex, seen map[Vertex]bool) bool {
	fmt.Printf("node %v\n", v)
	if _, exists := seen[v]; exists {
		return true
	}
	if _, exists := g.adj[v]; !exists {
		return false
	}
	seen[v] = true
	for w := range g.adj[v] {
		if g.walkGraph(w, seen) == true {
			return true
		}
	}
	return false
}

func (g *Graph) IsCyclic() bool {
	var v Vertex
	seen := make(map[Vertex]bool)
	for v = 0; v < g.v; v++ {
		if _, exists := seen[v]; exists {
			continue
		}
		if _, exists := g.adj[v][v]; exists {
			return true
		}
		if g.walkGraph(v, seen) {
			return true
		}
	}
	return false
}

/*
	Example:
	Input:
	2
	2 2
	0 1 0 0
	4 3
	0 1 1 2 2 3

	Output:
	1
	0
*/

func main() {
}
