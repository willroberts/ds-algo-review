package bfs

import (
	"container/list"

	"github.com/willroberts/ds-algo-review/datastructures/graph"
)

// BreadthFirstSearch searches a Graph structure's edges for paths.
// Returns 'true' when a path exists between the given nodes.
func BreadthFirstSearch(g *graph.Graph, src *graph.GraphNode, dst *graph.GraphNode) bool {
	next := &list.List{}
	visited := make(map[int]struct{}, 0)
	next.PushBack(src)
	for next.Len() > 0 {
		node := next.Remove(next.Back()).(*graph.GraphNode)
		if node == dst {
			return true
		}
		if _, ok := visited[node.ID()]; ok {
			continue
		}
		visited[node.ID()] = struct{}{}
		for e := node.GetAdjacent().Front(); e != nil; e = e.Next() {
			next.PushBack(e.Value.(*graph.GraphNode))
		}
	}
	return false
}
