package main

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

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

func (g *Graph) IsCyclic() bool {
	seen := make(map[Vertex]bool)
	for v, edgeMap := range g.adj {
		if _, exists := edgeMap[v]; exists {
			return true
		}
		seen[v] = true
		for w := range edgeMap {
			fmt.Printf("%v %v\n", v, w)
			spew.Dump(seen)
			if _, exists := seen[w]; exists {
				return true
			}
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
