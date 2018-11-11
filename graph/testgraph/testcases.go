// Copyright ©2018 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testgraph

import (
	"math"

	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
)

var testCases = []struct {
	// name is the name of the test.
	name string

	// nodes is the set of nodes that should be used
	// to construct the graph.
	nodes []graph.Node

	// edges is the set of edges that should be used
	// to construct the graph.
	edges []graph.WeightedEdge

	// nonexist is a set of nodes that should not be
	// found within the graph.
	nonexist []graph.Node

	// self is the weight value associated with
	// a self edge for simple graphs that do not
	// store individual self edges.
	self float64

	// absent is the weight value associated
	// with absent edges.
	absent float64
}{
	{
		name:     "empty",
		nonexist: []graph.Node{simple.Node(-1), simple.Node(0), simple.Node(1)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:     "one - negative",
		nodes:    []graph.Node{simple.Node(-1)},
		nonexist: []graph.Node{simple.Node(0), simple.Node(1)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:     "one - zero",
		nodes:    []graph.Node{simple.Node(0)},
		nonexist: []graph.Node{simple.Node(-1), simple.Node(1)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:     "one - positive",
		nodes:    []graph.Node{simple.Node(1)},
		nonexist: []graph.Node{simple.Node(-1), simple.Node(0)},
		self:     0,
		absent:   math.Inf(1),
	},

	{
		name:     "two - positive",
		nodes:    []graph.Node{simple.Node(1), simple.Node(2)},
		edges:    []graph.WeightedEdge{simple.WeightedEdge{F: simple.Node(1), T: simple.Node(2), W: 0.5}},
		nonexist: []graph.Node{simple.Node(-1), simple.Node(0)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:     "two - negative",
		nodes:    []graph.Node{simple.Node(-1), simple.Node(-2)},
		edges:    []graph.WeightedEdge{simple.WeightedEdge{F: simple.Node(-1), T: simple.Node(-2), W: 0.5}},
		nonexist: []graph.Node{simple.Node(0), simple.Node(-2)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:     "two - zero spanning",
		nodes:    []graph.Node{simple.Node(-1), simple.Node(1)},
		edges:    []graph.WeightedEdge{simple.WeightedEdge{F: simple.Node(-1), T: simple.Node(1), W: 0.5}},
		nonexist: []graph.Node{simple.Node(0), simple.Node(2)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:     "two - zero contiguous",
		nodes:    []graph.Node{simple.Node(0), simple.Node(1)},
		edges:    []graph.WeightedEdge{simple.WeightedEdge{F: simple.Node(0), T: simple.Node(1), W: 0.5}},
		nonexist: []graph.Node{simple.Node(-1), simple.Node(2)},
		self:     0,
		absent:   math.Inf(1),
	},

	{
		name:     "three - positive",
		nodes:    []graph.Node{simple.Node(1), simple.Node(2), simple.Node(3)},
		edges:    []graph.WeightedEdge{simple.WeightedEdge{F: simple.Node(1), T: simple.Node(2), W: 0.5}},
		nonexist: []graph.Node{simple.Node(-1), simple.Node(0)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:     "three - negative",
		nodes:    []graph.Node{simple.Node(-1), simple.Node(-2), simple.Node(-3)},
		edges:    []graph.WeightedEdge{simple.WeightedEdge{F: simple.Node(-1), T: simple.Node(-2), W: 0.5}},
		nonexist: []graph.Node{simple.Node(0), simple.Node(1)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:     "three - zero spanning",
		nodes:    []graph.Node{simple.Node(-1), simple.Node(0), simple.Node(1)},
		edges:    []graph.WeightedEdge{simple.WeightedEdge{F: simple.Node(-1), T: simple.Node(1), W: 0.5}},
		nonexist: []graph.Node{simple.Node(-2), simple.Node(2)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:     "three - zero contiguous",
		nodes:    []graph.Node{simple.Node(0), simple.Node(1), simple.Node(2)},
		edges:    []graph.WeightedEdge{simple.WeightedEdge{F: simple.Node(0), T: simple.Node(1), W: 0.5}},
		nonexist: []graph.Node{simple.Node(-1), simple.Node(3)},
		self:     0,
		absent:   math.Inf(1),
	},

	{
		name: "4-clique - single(non-prepared)",
		edges: func() []graph.WeightedEdge {
			n := 4
			edges := make([]graph.WeightedEdge, 0, (n*n-n)/2)
			for i := 0; i < 4; i++ {
				for j := i + 1; j < 4; j++ {
					edges = append(edges, simple.WeightedEdge{F: simple.Node(i), T: simple.Node(j), W: 0.5})
				}
			}
			return edges
		}(),
		nonexist: []graph.Node{simple.Node(-1), simple.Node(4)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name: "4-clique - single(prepared)",
		nodes: func() []graph.Node {
			n := 4
			nodes := make([]graph.Node, n)
			for i := range nodes {
				nodes[i] = simple.Node(i)
			}
			return nodes
		}(),
		edges: func() []graph.WeightedEdge {
			n := 4
			edges := make([]graph.WeightedEdge, 0, (n*n-n)/2)
			for i := 0; i < n; i++ {
				for j := i + 1; j < n; j++ {
					edges = append(edges, simple.WeightedEdge{F: simple.Node(i), T: simple.Node(j), W: 0.5})
				}
			}
			return edges
		}(),
		nonexist: []graph.Node{simple.Node(-1), simple.Node(4)},
		self:     0,
		absent:   math.Inf(1),
	},

	{
		name: "4-clique - double(non-prepared)",
		edges: func() []graph.WeightedEdge {
			n := 4
			edges := make([]graph.WeightedEdge, 0, n*n-n)
			for i := 0; i < n; i++ {
				for j := i + 1; j < n; j++ {
					edges = append(edges, simple.WeightedEdge{F: simple.Node(i), T: simple.Node(j), W: 0.5})
					edges = append(edges, simple.WeightedEdge{F: simple.Node(j), T: simple.Node(i), W: 0.5})
				}
			}
			return edges
		}(),
		nonexist: []graph.Node{simple.Node(-1), simple.Node(4)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name: "4-clique - double(prepared)",
		nodes: func() []graph.Node {
			n := 4
			nodes := make([]graph.Node, n)
			for i := range nodes {
				nodes[i] = simple.Node(i)
			}
			return nodes
		}(),
		edges: func() []graph.WeightedEdge {
			n := 4
			edges := make([]graph.WeightedEdge, 0, n*n-n)
			for i := 0; i < n; i++ {
				for j := i + 1; j < n; j++ {
					edges = append(edges, simple.WeightedEdge{F: simple.Node(i), T: simple.Node(j), W: 0.5})
					edges = append(edges, simple.WeightedEdge{F: simple.Node(j), T: simple.Node(i), W: 0.5})
				}
			}
			return edges
		}(),
		nonexist: []graph.Node{simple.Node(-1), simple.Node(4)},
		self:     0,
		absent:   math.Inf(1),
	},
}
