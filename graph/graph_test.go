package graph

import (
	"testing"
)

type testGraph struct {
	input    [][]int
	numNodes int
	numEdges int
}

var tests = []testGraph{
	testGraph{
		[][]int{{-1}},
		1,
		0,
	},
	testGraph{
		[][]int{{1}},
		1,
		1,
	},
	testGraph{
		[][]int{
			{-1, 2, 5, -1, 9},
			{6, -1, 12, -1, -1},
			{9, -1, -1, 12, 4},
			{65, 2, 4, -1, 12},
			{-1, 1234, -1, 4, -1},
		},
		5,
		14,
	},
}

func TestCreateGraph(t *testing.T) {
	for _, testGraph := range tests {
		g := CreateGraph(testGraph.input)
		if len(g.Nodes) != testGraph.numNodes {
			t.Error(
				"For", testGraph.input,
				"expected", testGraph.numNodes,
				"got", len(g.Nodes),
			)
		}
		if len(g.Edges) != testGraph.numEdges {
			t.Error(
				"For", testGraph.input,
				"expected", testGraph.numEdges,
				"got", len(g.Edges),
			)
		}
	}
}
