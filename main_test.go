package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleVertexPointingToNothingIsNotCycle(t *testing.T) {
	target := NewGraph(1)
	assert.False(t, target.IsCyclic())
}

func TestSingleVertexPointingToItselfIsCycle(t *testing.T) {
	target := NewGraph(1)
	target.AddEdge(0, 0)
	assert.True(t, target.IsCyclic())
}

func TestInputSet1(t *testing.T) {
	target := NewGraph(2)
	target.AddEdge(0, 1)
	target.AddEdge(0, 0)
	assert.True(t, target.IsCyclic())
}

func TestInputSet2(t *testing.T) {
	target := NewGraph(4)
	target.AddEdge(0, 1)
	target.AddEdge(1, 2)
	target.AddEdge(2, 3)
	assert.False(t, target.IsCyclic())
}

func TestSimpleCycle(t *testing.T) {
	target := NewGraph(2)
	target.AddEdge(0, 1)
	target.AddEdge(1, 0)
	assert.True(t, target.IsCyclic())
}

func TestCycleWithDistantVertex(t *testing.T) {
	target := NewGraph(4)
	target.AddEdge(0, 1)
	target.AddEdge(1, 2)
	target.AddEdge(2, 3)
	target.AddEdge(3, 0)
	assert.True(t, target.IsCyclic())
}

func TestNoCycleForDisjointedGraph(t *testing.T) {
	target := NewGraph(4)
	target.AddEdge(0, 1)
	target.AddEdge(2, 3)
	assert.False(t, target.IsCyclic())
}

func TestCycleForDisjointedGraph(t *testing.T) {
	target := NewGraph(4)
	target.AddEdge(0, 1)
	target.AddEdge(2, 3)
	target.AddEdge(3, 2)
	assert.True(t, target.IsCyclic())
}
