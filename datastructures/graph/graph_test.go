package graph

import "testing"

func TestGraph(t *testing.T) {
	g := NewGraph()
	n1 := NewGraphNode(1)
	g.AddNode(n1)
	n2 := NewGraphNode(2)
	g.AddNode(n2)
	g.AddEdge(1, 2)

	// Assert edge exists between n1 and n2.
	if n1.GetAdjacent().Front().Value != n2 {
		t.Fail()
	}
}
