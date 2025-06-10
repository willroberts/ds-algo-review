package bfs

import (
	"testing"

	"github.com/willroberts/ds-algo-review/datastructures/graph"
)

func TestBreadthFirstSearch(t *testing.T) {
	// Graph:
	// n1 -> n2 -> n4 -> n6
	//    -> n3 -> n5
	g := graph.NewGraph()
	n1 := graph.NewGraphNode(1)
	g.AddNode(n1)
	n2 := graph.NewGraphNode(2)
	g.AddNode(n2)
	n3 := graph.NewGraphNode(3)
	g.AddNode(n3)
	n4 := graph.NewGraphNode(4)
	g.AddNode(n4)
	n5 := graph.NewGraphNode(5)
	g.AddNode(n5)
	n6 := graph.NewGraphNode(6)
	g.AddNode(n6)
	g.AddEdge(1, 2)
	g.AddEdge(2, 4)
	g.AddEdge(4, 6)
	g.AddEdge(1, 3)
	g.AddEdge(3, 5)

	// Assert a path exists from n1 to n6.
	if !BreadthFirstSearch(g, n1, n6) {
		t.Log("a path does not exist from n1 to n6")
		t.Fail()
	}

	// Assert no path exists from n3 to n6.
	if BreadthFirstSearch(g, n3, n6) {
		t.Log("a path exists from n3 to n6")
		t.Fail()
	}
}
