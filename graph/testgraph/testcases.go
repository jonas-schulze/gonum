// Copyright ©2018 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package testgraph

import (
	"math"

	"gonum.org/v1/gonum/graph"
)

// node is a graph.Node implementation that is not exported
// so that other packages will not be aware of it's implementation.
type node int64

func (n node) ID() int64 { return int64(n) }

// line is an extended graph.Edge implementation that is not exported
// so that other packages will not be aware of its implementation. It
// covers all the edge types exported by graph.
type line struct {
	F, T graph.Node
	UID  int64
	W    float64
}

func (e line) From() graph.Node { return e.F }
func (e line) To() graph.Node   { return e.T }
func (e line) ID() int64        { return e.UID }
func (e line) Weight() float64  { return e.W }

var testCases = []struct {
	// name is the name of the test.
	name string

	// nodes is the set of nodes that should be used
	// to construct the graph.
	nodes []graph.Node

	// edges is the set of edges that should be used
	// to construct the graph.
	edges []graph.WeightedLine

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
		nonexist: []graph.Node{node(-1), node(0), node(1)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:     "one - negative",
		nodes:    []graph.Node{node(-1)},
		nonexist: []graph.Node{node(0), node(1)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:     "one - zero",
		nodes:    []graph.Node{node(0)},
		nonexist: []graph.Node{node(-1), node(1)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:     "one - positive",
		nodes:    []graph.Node{node(1)},
		nonexist: []graph.Node{node(-1), node(0)},
		self:     0,
		absent:   math.Inf(1),
	},

	{
		name:     "two - positive",
		nodes:    []graph.Node{node(1), node(2)},
		edges:    []graph.WeightedLine{line{F: node(1), T: node(2), UID: 0, W: 0.5}},
		nonexist: []graph.Node{node(-1), node(0)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:     "two - negative",
		nodes:    []graph.Node{node(-1), node(-2)},
		edges:    []graph.WeightedLine{line{F: node(-1), T: node(-2), UID: 0, W: 0.5}},
		nonexist: []graph.Node{node(0), node(-2)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:     "two - zero spanning",
		nodes:    []graph.Node{node(-1), node(1)},
		edges:    []graph.WeightedLine{line{F: node(-1), T: node(1), UID: 0, W: 0.5}},
		nonexist: []graph.Node{node(0), node(2)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:     "two - zero contiguous",
		nodes:    []graph.Node{node(0), node(1)},
		edges:    []graph.WeightedLine{line{F: node(0), T: node(1), UID: 0, W: 0.5}},
		nonexist: []graph.Node{node(-1), node(2)},
		self:     0,
		absent:   math.Inf(1),
	},

	{
		name:     "three - positive",
		nodes:    []graph.Node{node(1), node(2), node(3)},
		edges:    []graph.WeightedLine{line{F: node(1), T: node(2), UID: 0, W: 0.5}},
		nonexist: []graph.Node{node(-1), node(0)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:     "three - negative",
		nodes:    []graph.Node{node(-1), node(-2), node(-3)},
		edges:    []graph.WeightedLine{line{F: node(-1), T: node(-2), UID: 0, W: 0.5}},
		nonexist: []graph.Node{node(0), node(1)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:     "three - zero spanning",
		nodes:    []graph.Node{node(-1), node(0), node(1)},
		edges:    []graph.WeightedLine{line{F: node(-1), T: node(1), UID: 0, W: 0.5}},
		nonexist: []graph.Node{node(-2), node(2)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name:     "three - zero contiguous",
		nodes:    []graph.Node{node(0), node(1), node(2)},
		edges:    []graph.WeightedLine{line{F: node(0), T: node(1), UID: 0, W: 0.5}},
		nonexist: []graph.Node{node(-1), node(3)},
		self:     0,
		absent:   math.Inf(1),
	},

	{
		name: "4-clique - single(non-prepared)",
		edges: func() []graph.WeightedLine {
			n := 4
			var uid int64
			edges := make([]graph.WeightedLine, 0, (n*n-n)/2)
			for i := 0; i < 4; i++ {
				for j := i + 1; j < 4; j++ {
					edges = append(edges, line{F: node(i), T: node(j), UID: uid, W: 0.5})
					uid++
				}
			}
			return edges
		}(),
		nonexist: []graph.Node{node(-1), node(4)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name: "4-clique - single(prepared)",
		nodes: func() []graph.Node {
			n := 4
			nodes := make([]graph.Node, n)
			for i := range nodes {
				nodes[i] = node(i)
			}
			return nodes
		}(),
		edges: func() []graph.WeightedLine {
			n := 4
			var uid int64
			edges := make([]graph.WeightedLine, 0, (n*n-n)/2)
			for i := 0; i < n; i++ {
				for j := i + 1; j < n; j++ {
					edges = append(edges, line{F: node(i), T: node(j), UID: uid, W: 0.5})
					uid++
				}
			}
			return edges
		}(),
		nonexist: []graph.Node{node(-1), node(4)},
		self:     0,
		absent:   math.Inf(1),
	},

	{
		name: "4-clique - double(non-prepared)",
		edges: func() []graph.WeightedLine {
			n := 4
			var uid int64
			edges := make([]graph.WeightedLine, 0, n*n-n)
			for i := 0; i < n; i++ {
				for j := i + 1; j < n; j++ {
					edges = append(edges, line{F: node(i), T: node(j), UID: uid, W: 0.5})
					uid++
					edges = append(edges, line{F: node(j), T: node(i), UID: uid, W: 0.5})
					uid++
				}
			}
			return edges
		}(),
		nonexist: []graph.Node{node(-1), node(4)},
		self:     0,
		absent:   math.Inf(1),
	},
	{
		name: "4-clique - double(prepared)",
		nodes: func() []graph.Node {
			n := 4
			nodes := make([]graph.Node, n)
			for i := range nodes {
				nodes[i] = node(i)
			}
			return nodes
		}(),
		edges: func() []graph.WeightedLine {
			n := 4
			var uid int64
			edges := make([]graph.WeightedLine, 0, n*n-n)
			for i := 0; i < n; i++ {
				for j := i + 1; j < n; j++ {
					edges = append(edges, line{F: node(i), T: node(j), UID: uid, W: 0.5})
					uid++
					edges = append(edges, line{F: node(j), T: node(i), UID: uid, W: 0.5})
					uid++
				}
			}
			return edges
		}(),
		nonexist: []graph.Node{node(-1), node(4)},
		self:     0,
		absent:   math.Inf(1),
	},
}
